package models

import "gopkg.in/mgo.v2/bson"

// Representation of the object in mongodb (BSON)
// During import pay attention how you begin your variable name: lowercase(private) or UPPERCASE(public)

type Numobject struct {
	ID     bson.ObjectId `bson:"_id" json:"id"`
	Text   string        `bson:"text" json:"text"`
	Number int           `bson:"number" json:"number"`
	Found  bool          `bson:"found" json:"found"`
	Type   string        `bson:"type" json:"type"`
}
