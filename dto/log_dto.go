package dto

//	logDto结构体
type LogDto struct {
	Id                      int64  `json:"id"`
	BusinessName            string `json:"businessName"`
	ClassName               string `json:"className"`
	CreateDate              int64  `json:"createDate"`
	CurrentDepartment       string `json:"currentDepartment"`
	CurrentDepartmentName   string `json:"currentDepartmentName"`
	CurrentOrganization     string `json:"currentOrganization"`
	CurrentOrganizationName string `json:"currentOrganizationName"`
	EndDate                 int64  `json:"endDate"`
	ExpendTime              int64  `json:"expendTime"`
	Ip                      string `json:"ip"`
	MethodName              string `json:"methodName"`
	ModuleName              string `json:"moduleName"`
	OperatorAccount         string `json:"operatorAccount"`
	OperatorName            string `json:"operatorName"`
	OperatorNo              string `json:"operatorNo"`
	ResponseData            string `json:"reponseData"`
	RequestData             string `json:"requestData"`
	RequestMethod           string `json:"requestMethod"`
	RequestUri              string `json:"requestUri"`
	ServiceName             string `json:"serviceName"`
	StartDate               int64  `json:"startDate"`
	Success                 bool   `json:"success"`
}
