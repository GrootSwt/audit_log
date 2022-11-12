package dao

import (
	"audit_log/conn"
	"audit_log/convertor"
	"audit_log/model"
	"fmt"
	"time"
)

// 新增日志
func InsertLog(logModel model.LogModel) {
	db := conn.GetDB()
	defer db.Close()

	tx := db.Begin()

	err := tx.Create(&logModel).Error
	if err != nil {
		fmt.Println("插入数据错误：", err)
		tx.Rollback()
	} else {
		tx.Commit()
	}
}

// 根据id查询日志
func GetLogById(id int64) model.LogModel {
	db := conn.GetDB()
	defer db.Close()
	logModel := model.LogModel{
		Id: id,
	}
	db.First(&logModel)
	return logModel
}

// 分页条件查询
func GetByPaging(startRow int, pageSize int, currentDepartment string, currentOrganization string,
	startDate time.Time, endDate time.Time) (int, []model.LogModel) {
	db := conn.GetDB()
	defer db.Close()
	var totalCount int
	var logModelList []model.LogModel
	if len(currentDepartment) != 0 {
		db = db.Where("current_department = ?", currentDepartment)
	}
	if len(currentOrganization) != 0 {
		db = db.Where("current_organization = ?", currentOrganization)
	}
	if !startDate.IsZero() && !endDate.IsZero() {
		db = db.Where("create_date >= ? and create_date <= ?", startDate, endDate)
	}
	db.Table("audit_logging").Count(&totalCount)
	db.Offset(startRow).Limit(pageSize).Find(&logModelList)
	return totalCount, logModelList
}

// 数据转移
func TransferLog() {
	db := conn.GetDB()
	defer db.Close()
	tx := db.Begin()
	//	查询数据
	logModelList := make([]model.LogModel, 0)
	tx.Find(&logModelList)
	for _, logModel := range logModelList {
		logHistory := convertor.LogModelToLogModelHistory(logModel)
		tx.Create(&logHistory)
		tx.Delete(&logModel)
	}
	tx.Commit()
}
