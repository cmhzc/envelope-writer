package entity

type Envelope struct {
	EnvelopeID int64 `gorm:"column:envelope_id" json:"envelope_id"`
	UserID     int64 `gorm:"column:user_id" json:"uid"`
	Opened     bool  `gorm:"column:opened" json:"opened"`
	Value      int64 `gorm:"column:value" json:"value"`
	SnatchTime int64 `gorm:"column:snatch_time" json:"snatch_time"`
}
