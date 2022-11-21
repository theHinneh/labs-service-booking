package models

import (
	"time"
)

type Shop struct {
	ShopId        int       `json:"shop_id"`
	ShopName      string    `json:"shop_name"`
	Country       string    `json:"country"`
	City          string    `json:"city"`
	Address1      string    `json:"address_1"`
	Address2      string    `json:"address_2"`
	Zip           string    `json:"zip"`
	Email         string    `json:"email"`
	ContactId     int       `json:"contact-id"`
	WorkStartTime time.Time `json:"work_start_time"`
	WorkEndTime   time.Time `json:"work_end_time"`
}

//func (MgPortfolios *Shop) BeforeCreate(_ *gorm.DB) error {
//	MgPortfolios.ShopId = uuid.New().String()
//	return nil
//}
