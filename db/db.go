package db

import "go.mongodb.org/mongo-driver/bson/permitive"

const DBNAME = "specialnight"

func ToObjectID(id string) permitive.ObjectID {
	oid, err := permitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	return oid
}
