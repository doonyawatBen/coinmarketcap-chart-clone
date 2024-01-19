package infrastructure

import (
	"context"
)

var Log *Logger
var ConfigGlobal config
var CtxGlobal context.Context
var DataBaseGlobal *DataBase
var ServiceCloses []func()

var LogServiceManage = &LogRepository{}
var LogErrorServiceManage = &LogErrorRepository{}

func Init() {
	Log = NewLogger()

	EnvLoadStruct()

	CtxGlobal = context.Background()

	ConfigGlobal.DataBase.Collection.Quota = "quota"
	ConfigGlobal.DataBase.Collection.LogError = "logs_error"
	ConfigGlobal.DataBase.Collection.LogHTTP = "logs_http"
	DataBaseGlobal = NewDataBase()
}
