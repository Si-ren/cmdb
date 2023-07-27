package meeting

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type Meeting struct {
	global.GVA_MODEL
	MeetingRoomName string `json:"meetingRoomName" form:"meetingRoomName" gorm:"comment:会议室名称"` // 会议室名称
	MeetingTitle    string `json:"meetingTitle" form:"meetingTitle" gorm:"comment:会议名称"`        //会议标题
	MeetingDetail   string `json:"meetingDetail" form:"meetingDetail" gorm:"comment:会议细节"`      //会议细节
	Attendees       string `json:"attendee" form:"attendee" gorm:"comment:参会人"`                 //参会人
	StartTime       uint64 `json:"startTime" form:"startTime" gorm:"comment:会议室开始时间"`           //会议开始时间
	EndTime         uint64 `json:"endTime" form:"endTime" gorm:"comment:会议室结束时间"`               //会议结束时间

}
