package convertor

import (
	"audit_log/dto"
	"audit_log/model"
	"time"
)

// logDto转LogModel
func LogDtoTogModel(dto dto.LogDto) model.LogModel {
	var logModel model.LogModel
	logModel.Id = dto.Id
	logModel.BusinessName = dto.BusinessName
	logModel.ClassName = dto.ClassName
	if dto.CreateDate != 0 {
		logModel.CreateDate = time.Unix(dto.CreateDate/1000, 0)
	}
	logModel.CurrentDepartment = dto.CurrentDepartment
	logModel.CurrentDepartmentName = dto.CurrentDepartmentName
	logModel.CurrentOrganization = dto.CurrentOrganization
	logModel.CurrentOrganizationName = dto.CurrentOrganizationName
	if dto.EndDate != 0 {
		logModel.EndDate = time.Unix(dto.EndDate/1000, 0)
	}
	logModel.ExpendTime = dto.ExpendTime
	logModel.Ip = dto.Ip
	logModel.MethodName = dto.MethodName
	logModel.ModuleName = dto.ModuleName
	logModel.OperatorAccount = dto.OperatorAccount
	logModel.OperatorName = dto.OperatorName
	logModel.OperatorNo = dto.OperatorNo
	logModel.ResponseData = dto.ResponseData
	logModel.RequestData = dto.RequestData
	logModel.RequestMethod = dto.RequestMethod
	logModel.RequestUri = dto.RequestUri
	logModel.ServiceName = dto.ServiceName
	if dto.StartDate != 0 {
		logModel.StartDate = time.Unix(dto.StartDate/1000, 0)
	}
	logModel.Success = dto.Success
	return logModel
}

// logModel转logDto
func LogModelToLogDto(model model.LogModel) dto.LogDto {
	var logDto dto.LogDto
	logDto.Id = model.Id
	logDto.BusinessName = model.BusinessName
	logDto.ClassName = model.ClassName
	logDto.CreateDate = model.CreateDate.Unix()
	logDto.CurrentDepartment = model.CurrentDepartment
	logDto.CurrentDepartmentName = model.CurrentDepartmentName
	logDto.CurrentOrganization = model.CurrentOrganization
	logDto.CurrentOrganizationName = model.CurrentOrganizationName
	logDto.EndDate = model.EndDate.Unix()
	logDto.ExpendTime = model.ExpendTime
	logDto.Ip = model.Ip
	logDto.MethodName = model.MethodName
	logDto.ModuleName = model.ModuleName
	logDto.OperatorAccount = model.OperatorAccount
	logDto.OperatorName = model.OperatorName
	logDto.OperatorNo = model.OperatorNo
	logDto.ResponseData = model.ResponseData
	logDto.RequestData = model.RequestData
	logDto.RequestMethod = model.RequestMethod
	logDto.RequestUri = model.RequestUri
	logDto.ServiceName = model.ServiceName
	logDto.StartDate = model.StartDate.Unix()
	logDto.Success = model.Success
	return logDto
}
