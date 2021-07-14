package service

import (
	"go_code/audit_log/convertor"
	"go_code/audit_log/dao"
	"go_code/audit_log/dto"
	"go_code/audit_log/model"
	"go_code/audit_log/util"
	"log"
	"strconv"
	"time"
)

const (
	//	时间格式化值
	PATTERN = "2006-01-02 15:04:05"
)
//	插入log
func InsertLogService(logModel model.LogModel) {
	//	添加创建时间
	logModel.CreateDate = time.Now()
	//	获取id
	worker, _ := util.NewWorker(1)
	logModel.Id = worker.GetId()
	//	插入日志
	dao.InsertLog(logModel)
}
//	根据id查询日志
func GetById(idStr string) model.LogModel {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Fatalf("id错误: %s", err)
	}
	logModel := dao.GetLogById(id)
	return logModel
}
//	分页条件查询
func GetByPaging(pageStr string, pageSizeStr string,
	currentDepartment string, currentOrganization string,
	startDateStr string, endDateStr string) (int, interface{}) {
	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	startRow := (page - 1) * pageSize
	var startDate time.Time
	var endDate time.Time
	//	如果时间字符串不为0，再转换
	if startDateStr != "" {
		startDate, _ = time.ParseInLocation(PATTERN, startDateStr, time.Local)
	}
	if endDateStr != "" {
		endDate, _ = time.ParseInLocation(PATTERN, endDateStr, time.Local)
	}
	totalCount, logModelList := dao.GetByPaging(startRow, pageSize,
		currentDepartment, currentOrganization, startDate, endDate)
	logDtoList := make([]dto.LogDto, 0)
	if len(logModelList) != 0 {
		for _, logModel := range logModelList {
			logDtoList = append(logDtoList, convertor.LogModelToLogDto(logModel))
		}
		return totalCount, logDtoList
	} else {
		data := make([]interface{}, 0)
		return totalCount, data
	}
}
