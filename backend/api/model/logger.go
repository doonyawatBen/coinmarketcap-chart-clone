package model

type LogHTTP struct {
	Method    string `bson:"method" json:"method"`
	Status    int    `bson:"status" json:"status"`
	StatusMsg string `bson:"status_msg" json:"status_msg"`
	Path      string `bson:"path" json:"path"`
	IP        string `bson:"ip" json:"ip"`
	AppName   string `bson:"app_name" json:"app_name"`
	TokenID   string `bson:"token_id" json:"token_id"`
	Duration  int64  `bson:"duration" json:"duration"`
	Time      string `bson:"time" json:"time"`
	CreatedAt string `bson:"created_at" json:"created_at"`
}

type LogError struct {
	Time      string `bson:"time" json:"time"`
	Message   string `bson:"message" json:"message"`
	Error     string `bson:"error" json:"error"`
	Path      string `bson:"path" json:"path"`
	IP        string `bson:"ip" json:"ip"`
	AppName   string `bson:"app_name" json:"app_name"`
	TokenID   string `bson:"token_id" json:"token_id"`
	CreatedAt string `bson:"created_at" json:"created_at"`
}

type LogFiber struct {
	Path    string `bson:"path" json:"path"`
	IP      string `bson:"ip" json:"ip"`
	AppName string `bson:"app_name" json:"app_name"`
	TokenID string `bson:"token_id" json:"token_id"`
}
