package models

type Contacts struct {
	ContactId int    `json:"contact_id"`
	Contact   string `json:"contact"`
}

//func (MgPortfolios *Contacts) BeforeCreate(_ *gorm.DB) error {
//	MgPortfolios.ContactId = uuid.New().String()
//	return nil
//}
