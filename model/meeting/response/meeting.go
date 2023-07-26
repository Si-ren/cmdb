package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/meeting"
)

type MeetingListResponse struct {
	MeetingList []meeting.Meeting
}
