package convertor

import (
	"audit_log/model"
)

func LogModelToLogModelHistory(logModel model.LogModel) model.LogHistoryModel {
	var logHistory model.LogHistoryModel
	logHistory.Id = logModel.Id
	logHistory.BusinessName = logModel.BusinessName
	logHistory.ClassName = logModel.ClassName
	logHistory.CreateDate = logModel.CreateDate
	logHistory.CurrentDepartment = logModel.CurrentDepartment
	logHistory.CurrentDepartmentName = logModel.CurrentDepartmentName
	logHistory.CurrentOrganization = logModel.CurrentOrganization
	logHistory.CurrentOrganizationName = logModel.CurrentOrganizationName
	logHistory.EndDate = logModel.EndDate
	logHistory.ExpendTime = logModel.ExpendTime
	logHistory.MethodName = logModel.MethodName
	logHistory.ModuleName = logModel.ModuleName
	logHistory.OperatorAccount = logModel.OperatorAccount
	logHistory.OperatorName = logModel.OperatorName
	logHistory.OperatorNo = logModel.OperatorNo
	logHistory.ResponseData = logModel.ResponseData
	logHistory.RequestData = logModel.RequestData
	logHistory.RequestMethod = logModel.RequestMethod
	logHistory.RequestUri = logModel.RequestUri
	logHistory.ServiceName = logModel.ServiceName
	logHistory.StartDate = logModel.StartDate
	logHistory.Success = logModel.Success
	return logHistory
}
