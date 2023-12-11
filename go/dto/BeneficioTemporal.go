package dto

type BeneficioTemporal struct {
	BeneficiosAnuales []BeneficioAnual `json:"años"`
	BeneficiosMensuales []BeneficioMensual `json:"meses"`
}