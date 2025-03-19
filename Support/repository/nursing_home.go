package repository

import (
	"gorm.io/gorm"
	"your-project/model"
)

var db *gorm.DB

func SetDB(database *gorm.DB) {
	db = database
}

func GetNursingHomes() ([]model.NursingHome, error) {
	var homes []model.NursingHome
	result := db.Find(&homes)
	return homes, result.Error
}

func GetNursingHomeByID(id string) (*model.NursingHome, error) {
	var home model.NursingHome
	result := db.First(&home, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &home, nil
}

func CreateNursingHome(home *model.NursingHome) error {
	return db.Create(home).Error
}

func UpdateNursingHome(id string, home *model.NursingHome) error {
	return db.Model(&model.NursingHome{}).Where("id = ?", id).Updates(home).Error
}

func DeleteNursingHome(id string) error {
	return db.Delete(&model.NursingHome{}, id).Error
} 