package models

type Url struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	OriginalUrl string `gorm:"type:varchar(300)" json:"originial_url"`
	ShortenUrl  string `gorm:"type:varchar(300)" json:"shorten_url"`
	Description string `gorm:"type:text" json:"description"`
	PublishDate string `gorm:"type:date" json:"publish_date"`
}
