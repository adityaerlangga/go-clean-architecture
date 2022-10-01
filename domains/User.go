package domains

type User struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	FirstName string `gorm:"type:varchar(255)" json:"first_name"`
	LastName  string `gorm:"type:varchar(255)" json:"last_name"`
	Email     string `gorm:"type:varchar(255)" json:"email"`
	Password  string `gorm:"type:varchar(255)" json:"password"`
}

type Register struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ChangePassword struct {
	ID          uint   `json:"id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
