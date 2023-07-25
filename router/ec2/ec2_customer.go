package ec2

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CustomerRouter struct{}

func (e *CustomerRouter) InitCustomerRouter(Router *gin.RouterGroup) {
	customerRouter := Router.Group("ec2").Use(middleware.OperationRecord())
	customerRouterWithoutRecord := Router.Group("ec2")
	ec2CustomerApi := v1.ApiGroupApp.Ec2ApiGroup.CustomerApi
	{
		customerRouter.POST("customer", ec2CustomerApi.CreateEc2Host) // 测试prometheus指标

	}
	{
		customerRouterWithoutRecord.GET("customer", ec2CustomerApi.GetEc2Customer)         // 获取单一客户信息
		customerRouterWithoutRecord.GET("customerList", ec2CustomerApi.GetEc2CustomerList) // 获取客户列表
	}
}
