package meeting

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CustomerRouter struct{}

func (e *CustomerRouter) InitCustomerRouter(Router *gin.RouterGroup) {
	// customerRouter := Router.Group("meeting").Use(middleware.OperationRecord())
	customerRouter := Router.Group("meeting")
	customerRouterWithoutRecord := Router.Group("meeting")
	meetingCustomerApi := v1.ApiGroupApp.MeetingApiGroup.CustomerApi
	{
		customerRouter.POST("createmeeting", meetingCustomerApi.Createmeeting) //占用会议室

	}
	{
		customerRouterWithoutRecord.POST("meetingslist", meetingCustomerApi.GetMeetingsList) // 获取当前会议室的占用情况
	}
}
