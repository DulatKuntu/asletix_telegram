package repository

import (
	"asletix_telegram/model"
	"context"
	"errors"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type UserMongo struct {
	db *mongo.Database
}

func NewUserMongo(db *mongo.Database) *UserMongo {
	return &UserMongo{db: db}
}

func (r *UserMongo) TotalCount() (int64, error) {
	col := r.db.Collection(collectionUser)

	return col.CountDocuments(context.TODO(), bson.M{})
}

func (r *UserMongo) RegistrationLastMonthByDays() ([]*model.DayCount, error) {
	//init the loc
	loc, _ := time.LoadLocation("Asia/Almaty")

	//set timezone,
	now := time.Now().In(loc)

	var data []*model.DayCount
	log.Print("RegistrationLastMonthByDays")

	templateData := []*model.DayCount{}
	for i := 30; i >= 0; i-- {
		templateData = append(templateData,
			&model.DayCount{
				Count: 0,
				Time:  now.Add(-time.Duration(time.Hour * time.Duration(i) * 24)),
			},
		)
	}

	col := r.db.Collection(collectionUser)
	cursor, err := col.Aggregate(
		context.TODO(),
		[]bson.M{
			{"$match": bson.M{"userTimeData.registration": bson.M{"$gt": now.Add(-time.Duration(time.Hour * 30 * 24))}}},
			{"$group": bson.M{
				"_id": bson.M{
					"day":   bson.M{"$dayOfMonth": bson.M{"date": "$userTimeData.registration", "timezone": "Asia/Almaty"}},
					"month": bson.M{"$month": bson.M{"date": "$userTimeData.registration", "timezone": "Asia/Almaty"}},
				},
				"count": bson.M{"$sum": 1},
				"time":  bson.M{"$last": "$userTimeData.registration"},
			},
			},
			{"$sort": bson.M{"time": 1}},
		},
	)
	log.Print("RegistrationLastMonthByDays2")
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}

	for _, s := range data {
		log.Print(s.Time.String() + " Count: " + strconv.Itoa(s.Count))
	}

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(templateData); j++ {
			if data[i].Time.In(loc).Day() == templateData[j].Time.In(loc).Day() && data[i].Time.In(loc).Month() == templateData[j].Time.In(loc).Month() {
				templateData[j].Count = data[i].Count
				continue
			}
		}
	}

	log.Print("wrong template data")

	for _, s := range templateData {
		log.Print(s.Time.String() + " Count: " + strconv.Itoa(s.Count))
	}

	return templateData, nil
}

func (r *UserMongo) UniqueWorkoutLastDays(dayCount int) (int, error) {
	col := r.db.Collection(collectionFeeds)
	//init the loc
	loc, _ := time.LoadLocation("Asia/Almaty")

	//set timezone,
	now := time.Now().In(loc)
	timeGet := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	cursor, err := col.Aggregate(context.TODO(), []bson.M{
		{
			"$match": bson.M{
				"workoutTimeData.endTime": bson.M{
					"$gte": timeGet.Add(-time.Duration(time.Hour * time.Duration(dayCount) * 24)),
				},
			},
		},
		{
			"$group": bson.M{
				"_id": bson.M{
					"userId": "$userId",
				},
			},
		},
	})

	if err != nil {
		return 0, err
	}
	var inter []interface{}
	if err = cursor.All(context.TODO(), &inter); err != nil {
		return 0, err
	}

	return len(inter), nil
}

func (r *UserMongo) UniqueWorkoutLastMonthByDays() ([]*model.DayCount, error) {
	//init the loc
	loc, _ := time.LoadLocation("Asia/Almaty")

	//set timezone,
	now := time.Now().In(loc)

	var data []*model.DayCount

	templateData := []*model.DayCount{}
	for i := 30; i >= 0; i-- {
		templateData = append(templateData,
			&model.DayCount{
				Count: 0,
				Time:  now.Add(-time.Duration(time.Hour * time.Duration(i) * 24)),
			},
		)
	}

	col := r.db.Collection(collectionFeeds)
	cursor, err := col.Aggregate(
		context.TODO(),
		[]bson.M{
			{"$match": bson.M{"workoutTimeData.endTime": bson.M{"$gt": now.Add(-time.Duration(time.Hour * 30 * 24))}}},
			{"$group": bson.M{
				"_id": bson.M{
					"day":   bson.M{"$dayOfMonth": bson.M{"date": "$workoutTimeData.endTime", "timezone": "Asia/Almaty"}},
					"month": bson.M{"$month": bson.M{"date": "$workoutTimeData.endTime", "timezone": "Asia/Almaty"}},
				},
				"count": bson.M{"$sum": 1},
				"time":  bson.M{"$last": "$workoutTimeData.endTime"},
			},
			},
			{"$sort": bson.M{"time": 1}},
		},
	)

	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}

	for _, s := range data {
		log.Print(s.Time.String() + " Count: " + strconv.Itoa(s.Count))
	}

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(templateData); j++ {
			if data[i].Time.In(loc).Day() == templateData[j].Time.In(loc).Day() && data[i].Time.In(loc).Month() == templateData[j].Time.In(loc).Month() {
				templateData[j].Count = data[i].Count
				continue
			}
		}
	}

	log.Print("wrong template data")

	for _, s := range templateData {
		log.Print(s.Time.String() + " Count: " + strconv.Itoa(s.Count))
	}

	return templateData, nil
}

func (r *UserMongo) OpenAppLastDays(dayCount int) (int, error) {
	col := r.db.Collection(collectionUser)
	//init the loc
	loc, _ := time.LoadLocation("Asia/Almaty")

	//set timezone,
	now := time.Now().In(loc)
	timeGet := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	count, _ := col.CountDocuments(context.TODO(), bson.M{
		"userTimeData.lastIML": bson.M{
			"$gte": timeGet.Add(-time.Duration(time.Hour * time.Duration(dayCount) * 24)),
		},
	})
	return int(count), nil
}

func (r *UserMongo) OpenAppLastMonthByDays() ([]*model.DayCount, error) {
	//init the loc
	loc, _ := time.LoadLocation("Asia/Almaty")

	//set timezone,
	now := time.Now().In(loc)

	var data []*model.DayCount

	templateData := []*model.DayCount{}
	for i := 30; i >= 0; i-- {
		templateData = append(templateData,
			&model.DayCount{
				Count: 0,
				Time:  now.Add(-time.Duration(time.Hour * time.Duration(i) * 24)),
			},
		)
	}

	col := r.db.Collection(collectionUser)
	cursor, err := col.Aggregate(
		context.TODO(),
		[]bson.M{
			{"$match": bson.M{"userTimeData.lastIML": bson.M{"$gt": now.Add(-time.Duration(time.Hour * 30 * 24))}}},
			{"$group": bson.M{
				"_id": bson.M{
					"day":   bson.M{"$dayOfMonth": bson.M{"date": "$userTimeData.lastIML", "timezone": "Asia/Almaty"}},
					"month": bson.M{"$month": bson.M{"date": "$userTimeData.lastIML", "timezone": "Asia/Almaty"}},
				},
				"count": bson.M{"$sum": 1},
				"time":  bson.M{"$last": "$userTimeData.lastIML"},
			},
			},
			{"$sort": bson.M{"time": 1}},
		},
	)

	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}

	for _, s := range data {
		log.Print(s.Time.String() + " Count: " + strconv.Itoa(s.Count))
	}

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(templateData); j++ {
			if data[i].Time.In(loc).Day() == templateData[j].Time.In(loc).Day() && data[i].Time.In(loc).Month() == templateData[j].Time.In(loc).Month() {
				templateData[j].Count = data[i].Count
				continue
			}
		}
	}

	log.Print("wrong template data")

	for _, s := range templateData {
		log.Print(s.Time.String() + " Count: " + strconv.Itoa(s.Count))
	}

	return templateData, nil
}

func (r *UserMongo) Referal() ([]*model.RefCount, error) {
	col := r.db.Collection(collectionRefers)
	var allData []*model.RefCount

	cursor, err := col.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &allData); err != nil {
		return nil, errors.New("not found")
	}

	return allData, nil
}
