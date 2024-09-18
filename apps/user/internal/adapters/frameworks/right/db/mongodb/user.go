package mongodb

import "github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/models"

const userCollection = "user"

func (mApi mongoApi) AddUser(user models.UserModel) error {
	coll := mApi.connection.Database(mApi.ctx.Value("database").(string)).Collection(userCollection)
	_, err := coll.InsertOne(mApi.ctx, user)
	if err != nil {
		panic(err)
	}
	return nil
}
