package models

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	Id       int    `json:"id"`
	Make     string `json:"make"`
	CarModel string `json:"car_model"`
	Engine   string `json:"engine"`
	Year     string `json:"year"`
	EngineHP string `json:"engine_hp"`
}
