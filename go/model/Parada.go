package model

type Parada struct {
	IdEnvio      string `bson:"id_envio"`
	Ciudad       string `bson:"ciudad"`
	KmRecorridos int    `bson:"km_recorridos"`
}
