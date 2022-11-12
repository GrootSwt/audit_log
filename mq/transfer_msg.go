package mq

import (
	"audit_log/common"
	"audit_log/dao"

	"github.com/robfig/cron/v3"
)

// 定时完成log表数据转移到log_history
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
