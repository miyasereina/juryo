package replying

import (
	"io/ioutil"
	"math/rand"
	"sync"

	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"gopkg.in/yaml.v2"

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

	path := "./modules/replying/reply.yaml"
	bytes, err := ioutil.ReadFile("./modules/replying/reply.yaml")
	err = yaml.Unmarshal(bytes, &tem)
	if err != nil {

		logger.WithError(err).Errorf("unable to read config file in %s", path)
	}
}

func (a *ar) PostInit() {
}

func (a *ar) Serve(b *bot.Bot) {
	b.OnGroupMessage(func(c *client.QQClient, msg *message.GroupMessage) {
		out := autoreply(msg.ToString())
		if out == "" {
			return
		}
		m := message.NewSendingMessage().Append(message.NewText(out))
		c.SendGroupMessage(msg.GroupCode, m)
	})

	b.OnPrivateMessage(func(c *client.QQClient, msg *message.PrivateMessage) {
		out := autoreply(msg.ToString())
		if out == "" {
			return
		}
		m := message.NewSendingMessage().Append(message.NewText(out))
		c.SendPrivateMessage(msg.Sender.Uin, m)
	})
}

func (a *ar) Start(bot *bot.Bot) {
}

func (a *ar) Stop(bot *bot.Bot, wg *sync.WaitGroup) {
	defer wg.Done()
}

var instance *ar
var logger = utils.GetModuleLogger("logiase.autoreply")

var tem map[string]string

func autoreply(in string) string {
	j := 0
	data := []byte(in)
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
		return roll(reply[1:])
	default:
		return ""

	}

}

func roll(reply []string) string {
	if len(reply) <= 1 {
		return "这rand尼玛呢rand"
	}
	reply = reply[1:]
	return reply[rand.Intn(len(reply)+1)]
}
