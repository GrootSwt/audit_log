package model

import (
	"time"
)
//	logModel结构体
type LogModel struct {
	Id                      int64
	BusinessName            string
	ClassName               string
	CreateDate              time.Time
	CurrentDepartment       string
	CurrentDepartmentName   string
	CurrentOrganization     string
	CurrentOrganizationName string
	EndDate                 time.Time
	ExpendTime              int64
	Ip                      string
	MethodName              string
	ModuleName              string
	OperatorAccount         string
	OperatorName            string
	OperatorNo              string
	ResponseData            string
	RequestData             string
	RequestMethod           string
	RequestUri              string
	ServiceName             string
	StartDate               time.Time
	Success                 bool
}
//	设置日志存储表名
func (model LogModel) TableName() string {
	return "audit_logging"
}
