package model

type Parada struct {
	Ciudad       string `bson:"ciudad"`
	KmRecorridos int    `bson:"kmRecorridos"`
}
