package global

import (
	"github.com/camtrik/gin-blog/pkg/logger"
	"github.com/camtrik/gin-blog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
