package meeting

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type Meeting struct {
	global.GVA_MODEL
	MeetingRoomName string `json:"meetingRoomName" form:"meetingRoomName" gorm:"comment:会议室名称"` // 会议室名称
	StartTime       uint64 `json:"startTime" form:"startTime" gorm:"comment:会议室开始时间"`           //会议开始时间
	EndTime         uint64 `json:"endTime" form:"endTime" gorm:"comment:会议室结束时间"`               //会议结束时间
}
