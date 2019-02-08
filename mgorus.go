package mgorus

import (
	"context"
	"fmt"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/sirupsen/logrus"
)

type hooker struct {
	c *mongo.Collection
}

type M bson.M

func NewHooker(url, db, collection string) (*hooker, error) {
	client, err := mongo.NewClient(url)
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	return &hooker{c: client.Database(db).Collection(collection)}, nil
}

func NewHookerFromCollection(collection *mongo.Collection) *hooker {
	return &hooker{c: collection}
}

func NewHookerWithAuth(mgoUrl, db, collection, user, pass string) (*hooker, error) {
	panic("unimplemented")
}

func NewHookerWithAuthDb(mgoUrl, authDb, db, collection, user, pass string) (*hooker, error) {
	panic("unimplemented")
}

func (h *hooker) Fire(entry *logrus.Entry) error {
	data := make(logrus.Fields)
	data["Level"] = entry.Level.String()
	data["Time"] = entry.Time
	data["Message"] = entry.Message

	for k, v := range entry.Data {
		if errData, isError := v.(error); logrus.ErrorKey == k && v != nil && isError {
			data[k] = errData.Error()
		} else {
			data[k] = v
		}
	}

	_, err := h.c.InsertOne(context.Background(), M(data))

	if err != nil {
		return fmt.Errorf("failed to send log entry to mongodb: %v", err)
	}

	return nil
}

func (h *hooker) Levels() []logrus.Level {
	return logrus.AllLevels
}
