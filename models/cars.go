package models

type Car struct {
	Id                int    `json:"id"`
	Marca             string `json:"marca"`
	Modelo            string `json:"modelo"`
	Motor             string `json:"motor"`
	Ano               string `json:"ano"`
	PotenciaEmCavalos string `json:"potencia_em_cavalos"`
}
