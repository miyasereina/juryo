package replying

import (
	"fmt"
	"juryo/models"
	"math/rand"
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

func groupReply(msg *message.GroupMessage, c *client.QQClient) {
	morning(c)
	reply := []string{}
	ver := ""
	if re(msg.ToString()) != nil {
		fmt.Println(re(msg.ToString()))
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
	case "/up[Image:":
		{

			k, n := upImage(msg)
			m := message.NewSendingMessage().Append(message.NewText("识别到" + strconv.Itoa(k) + "张图片," + strconv.Itoa(n) + "张上传成功"))
			c.SendGroupMessage(msg.GroupCode, m)
		}
	case "/get":
		{
			fmt.Println("get")
			newE := models.Get(msg.GroupCode)
			url := "http://juryo.asakurayui.top/images/" + strconv.Itoa(int(msg.GroupCode)) + "/" + newE.Id
			m := message.NewSendingMessage().
				Append(message.NewText(url)).
				Append(message.NewGroupImage2(newE.Id, url, newE.Md5, newE.Fid, newE.Size, newE.Width, newE.Height, newE.ImageType))
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

func registerReply(b *bot.Bot) {
	b.OnGroupMessage(func(qqClient *client.QQClient, groupMessage *message.GroupMessage) {
		groupReply(groupMessage, qqClient)
	})

	b.OnPrivateMessage(func(qqClient *client.QQClient, privateMessage *message.PrivateMessage) {
		privateReply(privateMessage, qqClient)
	})

}
