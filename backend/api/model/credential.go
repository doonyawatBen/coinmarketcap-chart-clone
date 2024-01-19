package model

type Credential struct {
	AppName  string `json:"app_name"`
	AppCode3 string `json:"app_code_3"`
	IsAdmin  bool   `json:"is_admin"`
}
