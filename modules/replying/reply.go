package replying

import (
	"bufio"
	"fmt"
	"juryo/models"

	"math/rand"
	"net/http"
	"os"
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

func groupReply(msg *message.GroupMessage, c *client.QQClient) {
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
			c.SendGroupMessage(msg.GroupCode, m)
		}
	case "/up[Image:":
		{
			n := 0
			k := 0
			for _, elem := range msg.Elements {
				switch e := elem.(type) {
				case *message.GroupImageElement:
					{
						k++
						models.Upload(e, msg.GroupCode)
						download(e.Url, msg.GroupCode, e.ImageId)
						n++
					}
				default:
					continue
				}

			}
			m := message.NewSendingMessage().Append(message.NewText("识别到" + strconv.Itoa(k) + "张图片," + strconv.Itoa(n) + "张上传成功"))
			c.SendGroupMessage(msg.GroupCode, m)
		}
	case "/get":
		{
			newE := models.Get(msg.GroupCode)
			m := message.NewSendingMessage().Append(message.NewGroupImage(newE.Id, newE.Md5, newE.Fid, newE.Size, newE.Width, newE.Height, newE.ImageType))
			c.SendGroupMessage(msg.GroupCode, m)
		}
	case "/help":
		{
			m := message.NewSendingMessage().Append(message.NewText("https://github.com/miyasereina/juryo"))
			c.SendGroupMessage(msg.GroupCode, m)
		}

	}
}

func download(url string, path int64, fileName string) {
	imgPath := "images/" + strconv.Itoa(int(path))
	_, err := os.Stat(imgPath)
	if err != nil {
		if os.IsExist(err) {
			panic(err)
		}

		if os.IsNotExist(err) {
			err = os.Mkdir(imgPath, 755)
			if err != nil {
				panic(err)
			}
		}
	}
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("A error occurred!")
		return
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)

	file, err := os.Create(imgPath + "/" + fileName)
	if err != nil {
		panic(err)
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)
	bytes := make([]byte, 32*1024)
	for {
		len, err := reader.Read(bytes)

		if len < 0 || err != nil {
			return
		}
		// 注意这里byte数组后的[0:len]，不然可能会导致写入多余的数据
		_, _ = writer.Write(bytes[0:len])
		fmt.Printf("%d ", len)
	}

}

func roll(reply []string) string {
	//var n uint8
	//for _, value := range reply {
	//	n = n ^ value=="".(uint8)
	//}
	if len(reply) <= 1 /*|| n == 0*/ {
		return "这rand尼玛呢rand"
	}
	return reply[rand.Intn(len(reply))]
}

//func upH(reply []string) string {
//	msg:= *message.Image
//	imag()
//
//	return "上传成功"
//}

func registerReply(b *bot.Bot) {
	b.OnGroupMessage(func(qqClient *client.QQClient, groupMessage *message.GroupMessage) {
		groupReply(groupMessage, qqClient)
	})

	b.OnPrivateMessage(func(qqClient *client.QQClient, privateMessage *message.PrivateMessage) {
		privateReply(privateMessage, qqClient)
	})

}
