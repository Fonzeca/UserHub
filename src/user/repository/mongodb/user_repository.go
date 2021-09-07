package mongodb

import (
	"context"

	"github.com/Fonzeka/Jame/src/domain"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoUserRepository struct {
	col *qmgo.Collection
}

func NewMongoUserRepository(db *qmgo.Database) domain.UserRepository {
	collection := db.Collection("users")

	//Creamos el indice para que el userName no sea duplicado
	collection.CreateOneIndex(context.TODO(), options.IndexModel{Key: []string{"userName"}, Unique: true})

	return &MongoUserRepository{col: collection}
}

func (r *MongoUserRepository) GetAll(ctx context.Context) (res []domain.User, resErr error) {
	r.col.Find(ctx, bson.M{}).All(&res)
	return res, nil
}

func (r *MongoUserRepository) Insert(ctx context.Context, user *domain.User) (res domain.User, resErr error) {
	result, err := r.col.InsertOne(ctx, user)
	if err != nil {
		return *user, err
	}

	user.Id = result.InsertedID.(primitive.ObjectID)

	return *user, nil
}

func (r *MongoUserRepository) GetByUserName(ctx context.Context, userName string) (domain.User, error) {
	user := domain.User{}

	err := r.col.Find(ctx, bson.M{"userName": userName}).One(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}
