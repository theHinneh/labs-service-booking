package models

//type myTime time.Time
//
//func (mt *myTime) UnmarshalJSON(bs []byte) error {
//	var timestamp int64
//	err := json.Unmarshal(bs, &timestamp)
//	if err != nil {
//		return err
//	}
//
//	*mt = myTime(time.Unix(timestamp/1000, timestamp%1000*1e6))
//	return nil
//}
//
//func (mt myTime) MarshalJSON() ([]byte, error) {
//	timestamp := time.Time(mt).UnixNano() / 1e6
//	log.Println(time.Time(mt).UnixNano())
//	return json.Marshal(timestamp)
//}

type Appointment struct {
	AppointmentId   int    `json:"appointment_id"`
	ClientName      string `json:"client_name" binding:"required"`
	AppointmentDate string `json:"appointment_date" binding:"required"`
	ClientEmail     string `json:"client_email" binding:"required"`
	Note            string `json:"note" binding:"required"`
	Completed       bool   `json:"completed" binding:"required"`
	ServiceId       int    `json:"service_id" binding:"required"`
	ShopId          int    `json:"shop_id" binding:"required"`
	StaffId         int    `json:"staff_id" binding:"required"`
	ContactId       int    `json:"contact_id" binding:"required"`
}

type AppointmentStaffData struct {
	Shop    string `json:"shop"`
	Staff   string `json:"staff"`
	Service string `json:"service"`
	Contact string `json:"contact"`
}

//func (MgPortfolios *Appointment) BeforeCreate(_ *gorm.DB) error {
//	MgPortfolios.AppointmentId = uuid.New().String()
//	return nil
//}
