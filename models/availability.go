package models

import (
	"time"
)

type Availability struct {
	AvailabilityId int       `json:"availability_id"`
	Date           time.Time `json:"availability_date"`
	StaffId        int       `json:"staff_id"`
}

//func (MgPortfolios *Availability) BeforeCreate(_ *gorm.DB) error {
//	MgPortfolios.AvailabilityId = uuid.New().String()
//	return nil
//}
