package meeting

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/meeting"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
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
func (e *CustomerApi) Createmeeting(c *gin.Context) {
	var meeting meeting.Meeting
	err := c.ShouldBindJSON(&meeting)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(meeting, utils.MeetingVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = meetingService.CreateMeeting(meeting)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)

	global.RequestTotal.With(prometheus.Labels{
		"method":   "POST",
		"function": "Createmeeting",
		"status":   "200",
	}).Inc()

}

func (e *CustomerApi) GetMeetingsList(c *gin.Context) {
	start := time.Now()
	var meeting meeting.Meeting
	err := c.ShouldBindJSON(&meeting)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	//获取meetingList
	meetingList, err := meetingService.GetMeetingsList(meeting)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	global.GVA_LOG.Debug("Get meetings list: ", zap.Any("meetingList", meetingList))

	response.OkWithDetailed(meetingList, "查询成功", c)
	// 记录响应时间指标
	global.ResponseTime.With(prometheus.Labels{
		"method":   c.Request.Method,
		"function": "GetMeetingsList",
		"status":   "200",
	}).Observe(time.Since(start).Seconds())

}
