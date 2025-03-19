package service

import (
	"errors"
	"your-project/model"
	"your-project/repository"
)

func GetNursingHomes() ([]model.NursingHome, error) {
	return repository.GetNursingHomes()
}

func GetNursingHomeByID(id string) (*model.NursingHome, error) {
	return repository.GetNursingHomeByID(id)
}

func CreateNursingHome(home *model.NursingHome) error {
	// 添加基础验证
	if home.Name == "" || home.Address == "" {
		return errors.New("名称和地址不能为空")
	}
	return repository.CreateNursingHome(home)
}

func UpdateNursingHome(id string, home *model.NursingHome) error {
	if home.Name == "" || home.Address == "" {
		return errors.New("名称和地址不能为空")
	}
	return repository.UpdateNursingHome(id, home)
}

func DeleteNursingHome(id string) error {
	return repository.DeleteNursingHome(id)
} 