package models

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	UserID    uuid.UUID `gorm:"primaryKey;<-:create" json:"user_id"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	ProfileID uuid.UUID `json:"profile_id"`
	Profile   *Profile  `gorm:"foreignKey:profile_id" json:"profile"`
	DeviceID  string    `json:"device_id"`
}

type Profile struct {
	ID          uuid.UUID `gorm:"primaryKey;<-:create" json:"profile_id"`
	Level       uint      `json:"level"`
	DisplayName string    `json:"display_name"`
	AvatarURL   string    `json:"avatar_url"`
	Location    string    `json:"location"`
	CreatedDate time.Time `gorm:"<-:create" json:"created_date"`
	TimeZone    string    `json:"timezone"`
	State       string    `json:"state"`
	Email       string    `json:"email"`
}

type Authentication struct {
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	DeviceID string `json:"device_id"`
}

// type FriendList struct {
// 	UserIDFirst  string `json:"first_uuid"`
// 	UserIDSecond string `json:"second_uuid"`
// 	State int `json:"state"`
// }
