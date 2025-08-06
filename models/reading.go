package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Reading struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Ip          string             `json:"ip" bson:"ip"`
	Temperature string             `json:"temperature" bson:"temperature"`
	CpuUsage    string             `json:"cpuusage" bson:"cpuusage"`
	MemoryUsage string             `json:"memoryusage" bson:"memoryusage"`
	CreatedAt   time.Time          `json:"createdat" bson:"createdat"`
}
