package dto

type BeneficioTemporal struct {
	BeneficiosAnuales []BeneficioAnual `json:"anios"`
	BeneficiosMensuales []BeneficioMensual `json:"meses"`
}