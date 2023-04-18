package users

import "github.com/golang-jwt/jwt/v4"

type User struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Username string `gorm:"varchar(300)" json:"username"`
	Password string `gorm:"varchar(300)" json:"password"`
}

type MyClaims struct {
	Username    string `json:"username"`
	NamaLengkap string `json:"nama_lengkap"`
	jwt.StandardClaims
}
