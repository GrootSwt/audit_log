package mq

import (
	"github.com/robfig/cron/v3"
	"go_code/audit_log/common"
	"go_code/audit_log/dao"
)

//	定时完成log表数据转移到log_history
func TransferMsg() {
	crontab := cron.New()
	task := func() {
		dao.TransferLog()
	}
	// 	添加定时任务
	_, _ = crontab.AddFunc(common.ConfigInformation.Spec, task)
	//  启动定时器
	crontab.Start()
	//	阻塞
	select {}
}
