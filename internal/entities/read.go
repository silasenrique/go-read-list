package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Read struct {
	Id               primitive.ObjectID `json:"id" bson:"_id"`
	Name             string             `json:"name" bson:"name"`
	Category         Category           `json:"category" bson:"category"`
	Status           ReadingState       `json:"status" bson:"status"`
	LastChapterRead  int                `json:"lastChapterRead" bson:"lastChapterRead"`
	WeeklyReleaseDay WeekDay            `json:"releaseDay" bson:"releaseDay"`
	ReadIn           []string           `json:"readIn" bson:"readIn"`
}
