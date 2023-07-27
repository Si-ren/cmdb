package meeting

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/meeting"
	"github.com/flipped-aurora/gin-vue-admin/server/model/meeting/response"
	"go.uber.org/zap"
)

type CustomerService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateExaCustomer
//@description: 创建客户
//@param: e model.ExaCustomer
//@return: err error

func (mcs *CustomerService) CreateMeeting(meeting meeting.Meeting) (err error) {
	err = global.GVA_DB.Create(&meeting).Error
	if err != nil {
		global.GVA_LOG.Error("插入失败", zap.Error(err))
	}
	return err
}

func (mcs *CustomerService) GetMeetingsList(m meeting.Meeting) (ml response.MeetingListResponse, err error) {
	// ml.MeetingList = make([]meeting.Meeting, 10)
	global.GVA_DB.Where("meeting_room_name = ?", m.MeetingRoomName).Find(&ml.MeetingList)
	if len(ml.MeetingList) == 0 {
		err = errors.New("meeting not found")
	}
	return ml, err
}
