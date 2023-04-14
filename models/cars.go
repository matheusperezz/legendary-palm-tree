package models

type Car struct {
	Id       int    `json:"id"`
	Make     string `json:"make"`
	Model    string `json:"model"`
	Engine   string `json:"engine"`
	Year     string `json:"year"`
	EngineHP string `json:"engine_hp"`
}
