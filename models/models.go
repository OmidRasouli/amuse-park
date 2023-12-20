package models

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	UserID   uuid.UUID `gorm:"primaryKey;<-:create" json:"uuid"`
	Username string    `json:"username"`
	Role     string    `json:"role"`
	Profile  *Profile  `json:"profile"`
}

type Profile struct {
	Level       uint          `json:"level"`
	DisplayName string        `json:"display_name"`
	AvatarURL   *string       `json:"avatar_url"`
	Location    *string       `json:"location"`
	CreatedDate time.Time     `gorm:"<-:create" json:"created_date"`
	TimeZone    time.Location `json:"timezone"`
	State       string        `json:"state"`
}

type Authentication struct {
	Email    string  `gorm:"unique" json:"email"`
	Password *string `json:"password"`
	DeviceID string  `json:"device_id"`
}

// type FriendList struct {
// 	UserIDFirst  string `json:"first_uuid"`
// 	UserIDSecond string `json:"second_uuid"`
// 	State int `json:"state"`
// }
