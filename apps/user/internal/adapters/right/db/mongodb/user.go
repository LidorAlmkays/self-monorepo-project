package mongodb

import (
	"errors"

	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const userCollection = "user"

func (mApi *mongoApi) AddUser(user entities.User) error {

	coll := mApi.connection.Database(mApi.ctx.Value("database").(string)).Collection(userCollection)
	_, err := coll.InsertOne(mApi.ctx, user)
	if err != nil {
		panic(err)
	}

	return nil
}

func (mApi *mongoApi) GetUserByEmail(email string) (*entities.User, error) {
	coll := mApi.connection.Database(mApi.ctx.Value("database").(string)).Collection(userCollection)
	var user entities.User
	filter := bson.M{"email": email}
	err := coll.FindOne(mApi.ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			mApi.l.Error(errors.New("No document found with that email:" + email))
		}
		return nil, err
	}
	return &user, nil

}

func (mApi *mongoApi) GetUserByUsernameAndPassword(username string, password string) (*entities.User, error) {
	coll := mApi.connection.Database(mApi.ctx.Value("database").(string)).Collection(userCollection)
	var user entities.User
	filter := bson.M{"username": username, "password": password}
	err := coll.FindOne(mApi.ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			mApi.l.Error(errors.New("No document found with that username:" + username + ", and password: " + password))
		}
		return nil, err
	}
	return &user, nil
}
