package ec2

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	CustomerApi
}

var (
	customerService              = service.ServiceGroupApp.Ec2ServiceGroup.CustomerService
	fileUploadAndDownloadService = service.ServiceGroupApp.Ec2ServiceGroup.FileUploadAndDownloadService
)
