package logger

import (
	envconfig "centr_rosta/internal/config"

	"github.com/nool01/velog/pkg/velog"
	"github.com/nool01/velog/pkg/velog/velog_config"
)

var Log velog.DefaultLogger

func InitLogger() {
	config := velog_config.Config{
		Format:  "${level} ${l} ${name} ${l} ${content} ${l} ${timestamp}",
		Literal: " | ",
		Debug:   envconfig.Env.Debug,
	}

	Log = velog.Start(&config)
}
