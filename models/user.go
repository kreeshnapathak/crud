package models

type User struct {
	ID         string `gorm:"primary_key;not null;unique" json:"id"`
	Name       string `json:"name" form:"name"`
	Email      string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
	IsVerified bool   `json:"is_verified" bson:"is_verified"`
}
