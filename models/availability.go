package models

type Availability struct {
	AvailabilityId   int    `json:"availability_id"`
	AvailabilityDate string `json:"availability_date" binding:"required"`
	StaffId          int    `json:"staff_id" binding:"required"`
}

type AvailabilityDetailsResponse struct {
	StaffEmail       string `json:"staff_email"`
	StaffName        string `json:"staff_name"`
	AvailabilityDate string `json:"availability_date"`
}

//func (MgPortfolios *Availability) BeforeCreate(_ *gorm.DB) error {
//	MgPortfolios.AvailabilityId = uuid.New().String()
//	return nil
//}
