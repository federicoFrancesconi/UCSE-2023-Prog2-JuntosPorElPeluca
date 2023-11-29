package dto

type NuevaParada struct {
	IdEnvio      string `json:"id_envio"`
	Ciudad       string `json:"ciudad"`
	KmRecorridos int    `json:"km_recorridos"`
}

func (nuevaParada NuevaParada) GetParada() Parada {
	return Parada{
		Ciudad:       nuevaParada.Ciudad,
		KmRecorridos: nuevaParada.KmRecorridos,
	}
}