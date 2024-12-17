package mongodb

import (
	"errors"
	"reflect"

	"github.com/LidorAlmkays/self-monorepo-project/apps/authDB/internal/entities"
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

// UpdateUserByEmail implements db.DbPort.
func (mApi *mongoApi) UpdateUserByEmail(userEmail string, user *entities.User) error {
	coll := mApi.connection.Database(mApi.ctx.Value("database").(string)).Collection(userCollection)
	filter := bson.M{"email": userEmail}
	update, err := buildUpdateUserDocument(user)
	if err != nil {
		err = errors.New("Failed to update build user document update query, error: " + err.Error())
		mApi.l.Error(err)
		return err
	}
	// Perform the update
	result, err := coll.UpdateOne(mApi.ctx, filter, update)
	if err != nil {
		err = errors.New("Failed to update user document by email, error: " + err.Error())
		mApi.l.Error(err)
		return err
	}

	mApi.l.Info("successfully changed user password")

	if result.MatchedCount > 1 {
		err = errors.New(" when updating user password, found " + string(result.MatchedCount) + " users with the same email")
		mApi.l.Error(err)
		return err
	}

	return nil
}

// buildUpdateUserDocument dynamically creates the update document based on struct tags
func buildUpdateUserDocument(updatedUser *entities.User) (bson.M, error) {
	updateFields := bson.M{}

	// Get the actual value of the struct (dereference if necessary)
	v := reflect.ValueOf(updatedUser)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	typeOfUser := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := typeOfUser.Field(i)
		bsonTag := fieldType.Tag.Get("bson")

		// Skip fields with no BSON tag or explicitly set to "-"
		if bsonTag == "" || bsonTag == "-" {
			continue
		}

		// Add the field to the update document
		updateFields[bsonTag] = field.Interface()
	}

	return bson.M{"$set": updateFields}, nil
}
