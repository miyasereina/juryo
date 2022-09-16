package replying

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"juryo/models"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	"sync"

	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"juryo/bot"
	"juryo/utils"
)

func ReplyInit() {
	instance = &ar{}
	bot.RegisterModule(instance)
}

type ar struct {
}

func (a *ar) MiraiGoModule() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "logiase.reply",
		Instance: instance,
	}
}
func (a *ar) Init() {
}

func (a *ar) PostInit() {
}

func (a *ar) Serve(b *bot.Bot) {
	registerReply(b)
}

func (a *ar) Start(bot *bot.Bot) {
}

func (a *ar) Stop(bot *bot.Bot, wg *sync.WaitGroup) {
	defer wg.Done()
}

var instance *ar
var logger = utils.GetModuleLogger("logiase.autoreply")

func privateReply(msg *message.PrivateMessage, c *client.QQClient) {
	if []byte(msg.ToString())[0] != '/' {
		return
	}
	j := 0
	data := []byte(msg.ToString())

	reply := []string{}
	//双指针遍历还原为数组
	for i := 0; i <= len(data)-1; i++ {
		if data[i] == ' ' {
			str := string(data[j:i])
			reply = append(reply, str)
			j = i + 1
			continue
		}
		if i == len(data)-1 {
			str := string(data[j : i+1])
			reply = append(reply, str)
		}
	}
	switch reply[0] {
	case "/rand":
		{
			m := message.NewSendingMessage().Append(message.NewText(roll(reply[1:])))
			c.SendPrivateMessage(msg.Sender.Uin, m)
		}
	}
}

func re(str string) []string {
	rand := regexp.MustCompile("/up\\[Reply:.*?\\]")
	st := rand.FindAllString(str, -1)
	return st
}

var once sync.Once

type imageUploadResponse struct {
	UploadKey  []byte
	UploadIp   []uint32
	UploadPort []uint32
	ResourceId string
	Message    string
	FileId     int64
	Width      int32
	Height     int32
	ResultCode int32
	IsExists   bool
}

func groupReply(msg *message.GroupMessage, c *client.QQClient) {
	morning(c)
	reply := []string{}
	ver := ""
	if re(msg.ToString()) != nil {
		ver = re(msg.ToString())[0][0:9]
	} else {
		if []byte(msg.ToString())[0] != '/' {
			return
		}
		j := 0
		data := []byte(msg.ToString())

		//双指针遍历还原为数组
		for i := 0; i <= len(data)-1; i++ {
			if data[i] == ' ' {
				str := string(data[j:i])
				reply = append(reply, str)
				j = i + 1
				continue
			}
			if i == len(data)-1 {
				str := string(data[j : i+1])
				reply = append(reply, str)
				ver = reply[0]
			}
		}
	}
	switch ver {
	case "/ping":
		m := message.NewSendingMessage().Append(message.NewText("/pong"))
		c.SendGroupMessage(msg.GroupCode, m)
	case "/rand":
		{
			m := message.NewSendingMessage().Append(message.NewText("才不想告诉你应该选" + roll(reply[1:])))
			c.SendGroupMessage(msg.GroupCode, m)
		}
	case "/setu":
		{
			var r18 string
			n := len(reply)

			var img Img
			var r io.ReadSeeker
			if n <= 1 {
				r18 = "0"
				img, r = getsetu(r18, nil)
			} else {
				_, err := strconv.Atoi(reply[n-1])
				if err != nil {
					r18 = "0"
				} else {
					r18 = reply[n-1]
				}
				img, r = getsetu(r18, reply[1:n-1])
			}
			imgItem, err := c.UploadGroupImage(msg.GroupCode, r)
			if err != nil {
				panic(err)
			}
			m := message.NewSendingMessage().
				Append(message.NewText("[pid]:" + strconv.Itoa(img.Pid) + "\n")).
				Append(message.NewText("[老师]:" + img.Author)).
				Append(message.NewText(img.Urls.Original)).
				Append(imgItem)
			c.SendGroupMessage(msg.GroupCode, m)

		}
	case "/up[Image:":
		{

			k, n := upImage(msg)
			m := message.NewSendingMessage().Append(message.NewText("识别到" + strconv.Itoa(k) + "张图片," + strconv.Itoa(n) + "张上传成功"))
			c.SendGroupMessage(msg.GroupCode, m)
		}
	case "/get":
		{
			newE := models.Get(msg.GroupCode)
			url := "http://juryo.asakurayui.top/images/" + strconv.Itoa(int(msg.GroupCode)) + "/" + newE.Id
			m := message.NewSendingMessage().
				Append(message.NewText(url)).
				Append(NewGroupImage(newE.Id, url, newE.Md5, newE.Fid, newE.Size, newE.Width, newE.Height, newE.ImageType))
			c.SendGroupMessage(msg.GroupCode, m)
		}
	case "/help":
		{
			m := message.NewSendingMessage().Append(message.NewText("https://github.com/miyasereina/juryo"))
			c.SendGroupMessage(msg.GroupCode, m)
		}

	case "/今天":
		if msg.GroupCode != 735214237 {
			m := message.NewSendingMessage().Append(message.NewText("滚"))
			c.SendGroupMessage(msg.GroupCode, m)
		} else {
			m := message.NewSendingMessage().Append(message.NewText(month()))
			c.SendGroupMessage(msg.GroupCode, m)
		}

	case "/add":
		add(reply)
		m := message.NewSendingMessage().Append(message.NewText("已添加"))
		c.SendGroupMessage(msg.GroupCode, m)
	case "/clock":
		m := message.NewSendingMessage().Append(message.NewText(Clock(msg.Sender.Nickname)))
		c.SendGroupMessage(msg.GroupCode, m)

	}

}

func roll(reply []string) string {
	if len(reply) <= 1 {
		return "这rand尼玛呢rand"
	}
	return reply[rand.Intn(len(reply))]
}

type postform struct {
	R18   string   `json:"r18"`
	Proxy string   `json:"proxy"`
	Tag   []string `json:"tag"`
	size  string   `json:"size"`
}

func proxy(rawurl string) io.ReadSeeker {
	uri, err := url.Parse("http://127.0.0.1:7890")

	if err != nil {
		log.Fatal("parse url error: ", err)
	}
	client := http.Client{
		Transport: &http.Transport{
			// 设置代理
			Proxy: http.ProxyURL(uri),
		},
	}
	//client:=http.Client{}
	req := http.Request{}
	req.Method = "GET"
	req.URL, _ = url.Parse(rawurl)
	resp, err := client.Do(&req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	defer client.CloseIdleConnections()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%v", "done")
	return bytes.NewReader(data)

}
func getsetu(r18 string, tags []string) (Img, io.ReadSeeker) {
	var form postform
	fmt.Println(r18)
	form.R18 = r18
	form.Proxy = "i.pixiv.re"
	if len(tags) != 0 {
		form.Tag = tags
	}
	form.size = "small"
	bytesData, err := json.Marshal(form)
	req, _ := http.NewRequest("POST", "https://api.lolicon.app/setu/v2", bytes.NewReader(bytesData))
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	byts, err := io.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}
	var data Data
	err = json.Unmarshal(byts, &data)
	if err != nil {
		panic(err)
	}
	img := data.Data[0]
	fmt.Printf("%v", "done")
	return img, proxy(img.Urls.Original)
}

type Data struct {
	Data []Img `json:"data"`
}

type Img struct {
	Pid        int      `json:"pid"`         //作品 pid
	P          int      `json:"p"`           //作品所在页
	Uid        int      `json:"uid"`         //作者 uid
	Title      string   `json:"title"`       //作品标题
	Author     string   `json:"author"`      //作者名（入库时，并过滤掉 @ 及其后内容）
	R18        bool     `json:"r_18"`        //是否 R18（在库中的分类，不等同于作品本身的 R18 标识）
	Width      int      `json:"width"`       //原图宽度 px
	Height     int      `json:"height"`      //原图高度 px
	Tags       []string `json:"tags"`        //作品标签，包含标签的中文翻译（有的话）
	Ext        string   `json:"ext"`         //图片扩展名
	UploadDate int      `json:"upload_date"` //作品上传日期；时间戳，单位为毫秒
	Urls       object   `json:"urls"`        //包含了所有指定size的图片地址
}
type object struct {
	Original string
}

func registerReply(b *bot.Bot) {
	b.OnGroupMessage(func(qqClient *client.QQClient, groupMessage *message.GroupMessage) {
		groupReply(groupMessage, qqClient)
	})

	b.OnPrivateMessage(func(qqClient *client.QQClient, privateMessage *message.PrivateMessage) {
		privateReply(privateMessage, qqClient)
	})

}

func NewGroupImage(id, url string, md5 []byte, fid int64, size, width, height, imageType int32) *message.GroupImageElement {
	return &message.GroupImageElement{
		ImageId:   id,
		FileId:    fid,
		Md5:       md5,
		Size:      size,
		ImageType: imageType,
		Width:     width,
		Height:    height,
		Url:       url,
	}
}
