package meeting

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	CustomerApi
}

var (
	meetingService = service.ServiceGroupApp.MeetingServiceGroup.CustomerService
)
