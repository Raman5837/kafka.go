package model

import "time"

/*
Base Model Implementation

Contains:- is_deleted, created_at, modified_at
*/
type AbstractModel struct {
	CreatedAt  time.Time `gorm:"autoCreateTime;"`
	ModifiedAt time.Time `gorm:"autoUpdateTime;"`
	IsDeleted  bool      `gorm:"not null; default:false"`
}
