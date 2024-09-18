package models

type UserModel struct {
	UserName string `bson:"user_name"`
	Password string `bson:"password"`
	Age      int    `bson:"age"`
	Email    string `bson:"email"`
}
