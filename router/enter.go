package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/ec2"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Ec2     ec2.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
