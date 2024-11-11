package entities

import "time"

type User struct {
	Name     string    `bson:"name"`
	UserName string    `bson:"username"`
	Password string    `bson:"password"`
	BirthDay time.Time `bson:"birthday"`
	Email    string    `bson:"email"`
	Salt     string    `bson:"salt"`
	Role     string    `bson:"role"`
}
