package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"example.com/sarang-apis/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate mockgen -destination=../mocks/repository/mockUserRepository.go -package=repository example.com/sarang-apis/repository UserRepository
type UserRepositoryDB struct {
	UserCollection *mongo.Collection
}

type UserRepository interface {
	Insert(user models.User) (bool, error)
	GetAll() ([]models.User, error)
	Delete(id primitive.ObjectID) (bool, error)
	Update(id primitive.ObjectID, update bson.M) (bool, error)
}

func (t UserRepositoryDB) Insert(user models.User) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() //last place to work

	user.Id = primitive.NewObjectID()

	result, err := t.UserCollection.InsertOne(ctx, user) //process of adding data to mongo
	if result.InsertedID == nil || err != nil {

		return false, errors.New("failed add")
	}
	return true, nil
}

func (t UserRepositoryDB) GetAll() ([]models.User, error) {
	var user models.User
	var users []models.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.UserCollection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatalln(err)
	}

	for result.Next(ctx) {
		if err := result.Decode(&user); err != nil {
			log.Fatalln(err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil

}

func (t UserRepositoryDB) Delete(id primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.UserCollection.DeleteOne(ctx, bson.M{"id": id})

	if err != nil || result.DeletedCount <= 0 {
		return false, err
	}
	return true, nil
}
func (t UserRepositoryDB) Update(id primitive.ObjectID, update bson.M) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.UserCollection.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": update})
	if err != nil || result.ModifiedCount <= 0 {
		return false, err
	}
	return true, nil
}

func NewUserRepositoryDb(dbClient *mongo.Collection) UserRepositoryDB {
	return UserRepositoryDB{UserCollection: dbClient}
} //open a connection from main to db client
