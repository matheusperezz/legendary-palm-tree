package models

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	Make        string `json:"make"`
	CarModel    string `json:"car_model"`
	Engine      string `json:"engine"`
	Year        string `json:"year"`
	EnginePower string `json:"engine_power"`
}

/*	  	ID        uint           `gorm:"primaryKey"`
 *		CreatedAt time.Time
 *		UpdatedAt time.Time
 *		DeletedAt gorm.DeletedAt `gorm:"index"`
 */
