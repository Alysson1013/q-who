package models

import ( 
	"time"
	"gorm.io/gorm"
)

type User struct {
	id 			uuid.UUID  `gorm:"unique;not null;type:uuid;default:uuid_generate_v4()"`
	name 		string	   `gorm:"not null"`
	lastName 	string	   `gorm:"not null"`
	email 		string     `gorm:"unique;not null;type:varchar(100);default:null"`
	password    string     `gorm:"not null"`

	avatar      string
	role 		int
	createdAt 	time.Time  `gorm:"default:current_timestamp"`
}