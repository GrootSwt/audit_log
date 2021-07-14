package main

import (
	"github.com/gin-gonic/gin"
	"go_code/audit_log/common"
	"go_code/audit_log/controller"
	"go_code/audit_log/mq"
)
//	项目初始化时获取配置文件信息
func init() {
	common.LoadConfigInformation()
}

func main() {
	//	从RabbitMQ中获取消息
	go mq.GetMsg()
	//	定期转移Log数据
	go mq.TransferMsg()
	g := gin.Default()
	controller.Load(g)
	_ = g.Run(common.ConfigInformation.Port)

}
