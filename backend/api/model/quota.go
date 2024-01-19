package model

import "time"

type Quota struct {
	AppName        string    `json:"appName" bson:"appName"`
	Quota          int       `json:"quota" bson:"quota"`
	QuotaUsed      int       `json:"quota_used" bson:"quota_used"`
	QuotaRemaining int       `json:"quota_remaining" bson:"quota_remaining"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" bson:"updated_at"`
}

type QuotaDB struct {
	AppName        string    `json:"appName" bson:"appName"`
	Quota          int       `json:"quota" bson:"quota"`
	QuotaUsed      int       `json:"quota_used" bson:"quota_used"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" bson:"updated_at"`
}

type UpdateQuotaRequest struct {
	Quota int `json:"quota" bson:"quota" validate:"required"`
}

type UpdateQuotaRes struct {
	Status  string `json:"status"`
	Quota   int    `json:"quota" bson:"quota"`
	AppName string `json:"appName"`
}

type ResetQuotaRes struct {
	Status    string `json:"status"`
	QuotaUsed int    `json:"quota_used"`
	AppName   string `json:"appName,omitempty"`
}

type ResetAllQuotaRes struct {
	Status string `json:"status"`
}
