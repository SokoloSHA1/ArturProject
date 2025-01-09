package arturproject

import "time"

type User struct {
	Id         string    `json:"id" db "Id" binding:"required"`
	DeviceId   string    `json:"deviceId" db "DeviceId"`
	Locale     string    `json:"locale" db "Locale"`
	AppVersion string    `json:"appVersion" db "AppVersion"`
	CreatedAt  time.Time `json:"createdAt" db "CreatedAt"`
	LastSeen   time.Time `json:"lastSeen" db "LastSeen"`
}
