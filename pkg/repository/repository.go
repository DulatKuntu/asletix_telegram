package repository

import (
	"asletix_telegram/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type User interface {
	Referal() ([]*model.RefCount, error)
	RegistrationLastMonthByDays() ([]*model.DayCount, error)
	OpenAppLastMonthByDays() ([]*model.DayCount, error)
	OpenAppLastDays(int) (int, error)
	UniqueWorkoutLastMonthByDays() ([]*model.DayCount, error)
	UniqueWorkoutLastDays(int) (int, error)
	TotalCount() (int64, error)
}

type Repository struct {
	User
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		User: NewUserMongo(db),
	}
}
