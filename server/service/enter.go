package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/auto"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	AutoServiceGroup    auto.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
