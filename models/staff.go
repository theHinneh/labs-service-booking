package models

type Staff struct {
	StaffId   int    `json:"staff_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	ShopId    int    `json:"shop_id"`
	ContactId int    `json:"contact_id"`
}

//func (MgPortfolios *Staff) BeforeCreate(_ *gorm.DB) error {
//	MgPortfolios.StaffId = uuid.New().String()
//	return nil
//}
