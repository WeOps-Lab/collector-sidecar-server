package model

import "collector-sidecar-server/pkg/time"

type BaseModel struct {
	Id            int64          `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	EnabledStatus *int8          `gorm:"column:enabled_status;type:int;default:1" json:"-"`
	Created       time.JsonTime  `gorm:"column:created;type:timestamp;default:CURRENT_TIMESTAMP" json:"created"`
	Modified      *time.JsonTime `gorm:"column:modified;type:timestamp;default:CURRENT_TIMESTAMP" json:"modified"`
}
