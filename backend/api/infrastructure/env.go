package infrastructure

import (
	"github.com/kelseyhightower/envconfig"
)

func EnvLoadStruct() {
	err := envconfig.Process("", &ConfigGlobal)
	if err != nil {
		Log.Fatalln(nil, "init", "", "Can't load env system into struct", err)
	}
}
