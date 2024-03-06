package model

import "gorm.io/gorm"

type SidecarTokenModel struct {
	gorm.Model
	Token string
}
