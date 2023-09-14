package models
import "time"

type Customers struct {
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name"`
	GetStart time.Time `json:"getStart"`
	FinishTime    time.Time `json:"finishTime"`
	Time    string `json:"time"`
	
}