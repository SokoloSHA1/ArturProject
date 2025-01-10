package arturproject

type User struct {
	Id         string `json:"id" db:"Id" binding:"required"`
	DeviceId   string `json:"deviceId" db:"DevicedID"`
	Locale     string `json:"locale" db:"Locale"`
	AppVersion string `json:"appVersion" db:"UpVersion"`
	CreatedAt  string `json:"createdAt" db:"CreatedAt"`
	LastSeen   string `json:"lastSeen" db:"LastSeen"`
}
