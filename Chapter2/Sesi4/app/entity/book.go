package entity

import "time"

type Book struct {
	ID 		int    `gorm:"primaryKey" json:"id"`
	Title 	string `grom:"type:varchar(100)" json:"title"`
	Author 	string `grom:"type:varchar(100)" json:"author"`
	CreatedAt  time.Time 
	UpdatedAt  time.Time
}