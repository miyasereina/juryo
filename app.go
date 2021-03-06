package main

import (
	"juryo/dao"
	"juryo/models"
	"juryo/modules/logging"
	"juryo/modules/replying"
	"os"
	"os/signal"

	"juryo/bot"
	"juryo/config"
	_ "juryo/modules/logging"
	"juryo/utils"
)

func init() {
	utils.WriteLogToFS(utils.LogInfoLevel, utils.LogWithStack)
	config.Init()
	dao.MysqlInit()
	models.TableInit()
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
	bot.SaveToken()
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
