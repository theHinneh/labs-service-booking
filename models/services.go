package models

type Service struct {
	ServiceId   int     `json:"service_id"`
	StartPrice  float32 `json:"start_price"`
	EndPrice    float32 `json:"end_price"`
	StaffId     int     `json:"staff_id"`
	ServiceName string  `json:"service_name"`
	ShopId      int     `json:"shop_id"`
}

type ServiceDetails struct {
	ServiceId       int     `json:"service_id"`
	StartPrice      float32 `json:"start_price"`
	EndPrice        float32 `json:"end_price"`
	ServiceName     string  `json:"service_name"`
	AppointmentDate string  `json:"appointment_date"`
	ShopName        string  `json:"shop_name"`
	Country         string  `json:"country"`
	StaffName       string  `json:"staff_name"`
}

//func (MgPortfolios *Service) BeforeCreate(_ *gorm.DB) error {
//	MgPortfolios.ServiceId = uuid.New().String()
//	return nil
//}
