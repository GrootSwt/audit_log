package controller

import (
	"github.com/gin-gonic/gin"
	"go_code/audit_log/convertor"
	"go_code/audit_log/service"
	"net/http"
)

func Load(g *gin.Engine) {
	g.GET("getById/:id", getByIdHandler)
	g.GET("getByPaging/:page/:pageSize/:currentDepartment/:currentOrganization/" +
		":startDate/:endDate", getByPagingHandler)
}

//	根据id查询log
func getByIdHandler(c *gin.Context) {
	idStr := c.Param("id")
	logModel := service.GetById(idStr)
	logDto := convertor.LogModelToLogDto(logModel)
	c.JSON(http.StatusOK, gin.H{
		"status":     1,
		"statusInfo": "success",
		"data":       logDto,
	})
}

//	分页查询log
func getByPagingHandler(c *gin.Context) {
	pageStr := c.Param("page")
	pageSizeStr := c.Param("pageSize")
	currentDepartment := c.Param("currentDepartment")
	currentOrganization := c.Param("currentOrganization")
	startDateStr := c.Param("startDate")
	endDateStr := c.Param("endDate")
	totalCount, logDtoList := service.GetByPaging(pageStr, pageSizeStr,
		currentDepartment, currentOrganization, startDateStr, endDateStr)
	c.JSON(http.StatusOK, gin.H{
		"status":     1,
		"statusInfo": "success",
		"data":       logDtoList,
		"totalCount": totalCount,
	})
}
