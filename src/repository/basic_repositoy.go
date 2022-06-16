package repository

import (
	"api-merca/src/database"
	"api-merca/src/model"
)

type Basic struct {
}

func (br Basic) Insert(model model.IModel) (uint64, error) {
	if err := model.Validate(); err != nil {
		return 0, err
	}
	if err := database.Connection.With("POSTGRES").Create(model).Error; err != nil {
		return 0, err
	}
	return model.GetId(), nil
}

func (br Basic) Update(model model.IModel) error {
	if err := model.Validate(); err != nil {
		return err
	}
	return database.Connection.With("POSTGRES").Save(model).Error
}

func (br Basic) Save(model model.IModel) (uint64, error) {
	if err := model.Validate(); err != nil {
		return 0, err
	}
	if err := database.Connection.With("POSTGRES").Save(model).Error; err != nil {
		return 0, err
	}
	return model.GetId(), nil
}

func (br Basic) SaveAll(models interface{}) error {

	if err := database.Connection.With("POSTGRES").Save(models).Error; err != nil {
		return err
	}
	return nil
}

func (br Basic) FindById(receiver model.IModel, id interface{}) error {
	return database.Connection.With("POSTGRES").First(receiver, id).Error
}

func (br Basic) FindFirst(receiver model.IModel, where string, args ...interface{}) error {
	return database.Connection.With("POSTGRES").Where(where, args...).Limit(1).Find(receiver).Error
}

func (br Basic) FindAll(models interface{}, where string, args ...interface{}) (err error) {
	err = database.Connection.With("POSTGRES").Where(where, args...).Find(models).Error
	return
}

func (br Basic) Delete(model model.IModel, where string, args ...interface{}) error {
	return database.Connection.With("POSTGRES").Where(where, args...).Delete(&model).Error
}
