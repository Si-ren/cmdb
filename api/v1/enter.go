package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/ec2"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/meeting"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	Ec2ApiGroup     ec2.ApiGroup
	MeetingApiGroup meeting.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
