package ec2

import (
	"math/rand"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

type CustomerApi struct{}

// CreateEc2Customer
// @Tags      Ec2Customer
// @Summary   创建客户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      Ec2mple.Ec2Customer            true  "客户用户名, 客户手机号码"
// @Success   200   {object}  response.Response{msg=string}  "创建客户"
// @Router    /customer/customer [post]
func (e *CustomerApi) CreateEc2Host(c *gin.Context) {
	global.RequestTotal.With(prometheus.Labels{
		"method":   "POST",
		"function": "CreateEc2Host",
		"status":   "200",
	}).Inc()

	response.OkWithMessage("创建成功", c)
}

func (e *CustomerApi) GetEc2Customer(c *gin.Context) {
	start := time.Now()

	// 模拟处理时间
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	response.OkWithMessage("获取成功", c)
	// 记录响应时间指标
	global.ResponseTime.With(prometheus.Labels{
		"method":   c.Request.Method,
		"function": "GetEc2Customer",
		"status":   "200",
	}).Observe(time.Since(start).Seconds())

}

func (e *CustomerApi) GetEc2CustomerList(c *gin.Context) {
	start := time.Now()

	// 模拟处理时间
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	response.OkWithMessage("获取成功", c)
	// 记录响应时间指标
	global.ResponseTime.With(prometheus.Labels{
		"method":   c.Request.Method,
		"function": "GetEc2CustomerList",
		"status":   "200",
	}).Observe(time.Since(start).Seconds())

}
