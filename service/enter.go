package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/ec2"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/meeting"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	Ec2ServiceGroup     ec2.ServiceGroup
	MeetingServiceGroup meeting.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
