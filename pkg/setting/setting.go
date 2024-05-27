package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	// configs from flag have the highest priority
	for _, config := range configs {
		if config != "" {
			vp.AddConfigPath(config)
		}
	}
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()

	if err != nil {
		return nil, err
	}
	fmt.Println("Using config file:", vp.ConfigFileUsed())

	s := &Setting{vp}
	s.WatchSettingChange()
	return &Setting{vp}, nil
}

func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			fmt.Printf("Config file changed: %s", in.Name)
			_ = s.ReloadAllSection()
		})
	}()
}
