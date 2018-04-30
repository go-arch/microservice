package models

import (
	"time"
	"github.com/jinzhu/gorm"
)


// User model
type User struct {
	// auto-populate columns: id, created_at, updated_at, deleted_at
	gorm.Model


	Name string `json:"name";sql:"type:VARCHAR(255)"`
	Email string `json:"email";gorm:"type:VARCHAR(255);unique_index"`
	Password string `json:"password";sql:"type:VARCHAR(255)"`
	PhoneNumber string `json:"phoneNumber";sql:"type:VARCHAR(255)"`
	Date_Of_Birth time.Time `json:"date_of_birth";sql:"type:timestamp"`
	Gender string `json:"gender";sql:"type:VARCHAR(36)"`
	Age  int `json:"age";sql:"type:INT"`
	Father_Name string `json:"father_name";sql:"type:VARCHAR(255)"`
	Mother_Name string `json:"mother_name";sql:"type:VARCHAR(255)"`
	Mark_Of_Identification string `json:"mark_of_identification";sql:"type:VARCHAR(255)"`
	Category string `json:"category";sql:"type:VARCHAR(255)"`
	Religion string `json:"religion";sql:"type:VARCHAR(255)"`
	Alternate_Contact_No string `json:"alternate_contact_no";sql:"type:VARCHAR(255)"`
	Marital_Status string `json:"marital_status";sql:"type:VARCHAR(255)"`
	Differently_Abled string `json:"differently_abled";sql:"type:VARCHAR(255)"`
	Disability_Category string `json:"disability_category";sql:"type:VARCHAR(255)"`
	Employed_Or_Self_Employed string `json:"employed_or_self_employed";sql:"type:VARCHAR(255)"`
	Bpl_Number string `json:"bpl_number";sql:"type:VARCHAR(255)"`
	Bpl_Or_Apl string `json:"bpl_or_apl";sql:"type:VARCHAR(255)"`
	Personal_Income_Annual string `json:"personal_income_annual";sql:"type:VARCHAR(255)"`
	Family_Income_Annual string `json:"family_income_annual";sql:"type:VARCHAR(255)"`
	Spouse_Income_Annual string `json:"spouse_income_annual";sql:"type:VARCHAR(255)"`
	Passport_No string `json:"passport_no";sql:"type:VARCHAR(255)"`
	Ration_Card_No string `json:"ration_card_no";sql:"type:VARCHAR(255)"`
	Aadhar_No int `json:"aadhar_no";sql:"type:INT"`
	Driving_License_No string `json:"driving_license_no";sql:"type:VARCHAR(255)"`
	Pan_Card_No string `json:"pan_card_no";sql:"type:VARCHAR(255)"`
	Role string `json:"role";sql:"type:VARCHAR(36)"`
}

// user create response model
type UserCreateResponse struct {
	Created bool `json:"created"`
	Id int        `json:"id"`
}
