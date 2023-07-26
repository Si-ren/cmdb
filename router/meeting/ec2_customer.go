package meeting

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CustomerRouter struct{}

func (e *CustomerRouter) InitCustomerRouter(Router *gin.RouterGroup) {
	customerRouter := Router.Group("meeting").Use(middleware.OperationRecord())
	customerRouterWithoutRecord := Router.Group("meeting")
	meetingCustomerApi := v1.ApiGroupApp.MeetingApiGroup.CustomerApi
	{
		customerRouter.POST("customer", meetingCustomerApi.CreateEc2Host) //

	}
	{
		customerRouterWithoutRecord.GET("customer", meetingCustomerApi.GetMeetingRoom) // 获取当前会议室的占用情况
	}
}
