package mongodb

import (
	"context"

	"github.com/Fonzeca/UserHub/server/domain"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	opt "go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRolesRepository struct {
	col *qmgo.Collection
}

func NewMongoRolesRepository(db *qmgo.Database) domain.RolesRepository {
	collection := db.Collection("roles")

	//Creamos el indice para que el userName no sea duplicado
	collection.CreateOneIndex(context.TODO(), options.IndexModel{Key: []string{"name"}, IndexOptions: &opt.IndexOptions{Unique: &[]bool{true}[0]}})

	return &MongoRolesRepository{col: collection}
}

func (r *MongoRolesRepository) GetAll(ctx context.Context) (res []domain.Role, err error) {
	r.col.Find(ctx, bson.M{}).All(&res)
	return res, nil
}

func (r *MongoRolesRepository) Insert(ctx context.Context, role *domain.Role) error {
	_, err := r.col.InsertOne(ctx, role)
	return err
}

func (r *MongoRolesRepository) Delete(ctx context.Context, name string) error {
	return r.col.Remove(ctx, bson.M{"name": name})
}
