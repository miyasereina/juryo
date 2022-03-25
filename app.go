package main

import (
	"github.com/Logiase/MiraiGo-Template/modules/logging"
	"github.com/Logiase/MiraiGo-Template/modules/replying"
	"os"
	"os/signal"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Logiase/MiraiGo-Template/config"
	"github.com/Logiase/MiraiGo-Template/utils"

	_ "github.com/Logiase/MiraiGo-Template/modules/logging"
)

func init() {
	utils.WriteLogToFS(utils.LogInfoLevel, utils.LogWithStack)
	config.Init()
	replying.ReplyInit()
	logging.Loginit()
}

func main() {
	// 快速初始化
	bot.Init()

	// 使用协议
	// 不同协议可能会有部分功能无法使用
	// 在登陆前切换协议
	bot.UseProtocol(bot.AndroidPhone)

	//登录
	bot.Login()
	// 初始化 Modules
	bot.StartService()
	// 刷新好友列表，群列表
	bot.RefreshList()
	//回复

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	bot.Stop()
}
