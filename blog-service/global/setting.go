package global

import (
	"github.com/go-programing-tour-book/blog-service/pkg/logger"
	"github.com/go-programing-tour-book/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DataBaseSetting *setting.DatabaseSettingS

	Logger *logger.Logger
)
