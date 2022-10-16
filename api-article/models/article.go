package models

import (
	"time"
)

type Article struct {
	Id           int64     `gorm:"primaryKey, not null, autoIncrement" json:"id"`
	Title        string    `gorm:"type:varchar(200)" json:"title"`
	Content      string    `gorm:"type:text" json:"content"`
	Category     string    `gorm:"type:varchar(100)" json:"category"`
	Created_date time.Time `json:"created_date" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	Updated_date time.Time `json:"updated_date"`
	Status       string    `gorm:"type:varchar(100)" json:"status"`
}
