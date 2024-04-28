package models

type User struct {
	BaseModel `gorm:"embedded"`
	Name      string `json:"name"`
	Email     string `gorm:"unique;not null" json:"email"`
	Password  string `gorm:"not null" json:"password"`
	RoleID    uint   `json:"roleId"`
	Role      *Role  `gorm:"foreignKey:RoleID;references:id" json:"role"`
}

type UserWithoutPassword struct {
	BaseModel `gorm:"embedded"`
	Name      string `json:"name"`
	Email     string `gorm:"unique;not null" json:"email"`
	RoleID    uint   `json:"roleId"`
	Role      *Role  `gorm:"foreignKey:RoleID" json:"role"`
}

type Role struct {
	BaseModel `gorm:"embedded"`
	Name      string `gorm:"not null" json:"name"`
}
