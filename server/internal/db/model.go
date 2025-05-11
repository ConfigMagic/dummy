package db

type Config struct {
	Name string `bson:"name"`
	Data string `bson:"data"`
}
