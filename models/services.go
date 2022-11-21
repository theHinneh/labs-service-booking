package models

type Service struct {
	ServiceId  int     `json:"service_id"`
	StartPrice float32 `json:"start_price"`
	EndPrice   float32 `json:"end_price"`
	StaffId    int     `json:"staff_id"`
	ShopId     int     `json:"shop_id"`
}

//func (MgPortfolios *Service) BeforeCreate(_ *gorm.DB) error {
//	MgPortfolios.ServiceId = uuid.New().String()
//	return nil
//}
