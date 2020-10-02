package mongodb

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
InfoToDB : Descriptions of the data type and notes if any
*/
type InfoToDB struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	CreateAt time.Time          `bson:"createAt"`
}

/*
Validate : This function checks the validity of data fields in the Info structure
*/
func (ins *InfoToDB) Validate() error {
	if len(ins.Name) == 0 {
		return errors.New("Name field is invalid")
	}
	return nil
}
