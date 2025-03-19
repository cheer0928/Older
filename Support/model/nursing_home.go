package model

import "time"

type NursingHome struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Name        string    `json:"name" gorm:"size:100;not null"`
    Address     string    `json:"address" gorm:"size:200;not null"`
    Phone       string    `json:"phone" gorm:"size:20"`
    Beds        int       `json:"beds" gorm:"not null"`
    Price       float64   `json:"price" gorm:"not null"`
    Rating      float32   `json:"rating"`
    Services    string    `json:"services" gorm:"type:text"`
    Description string    `json:"description" gorm:"type:text"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
} 