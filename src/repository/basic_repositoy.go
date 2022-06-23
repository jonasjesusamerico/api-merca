package repository

import (
	"api-merca/src/contexto"
	"api-merca/src/database"
	"api-merca/src/model"

	"gorm.io/gorm"
)

type Basic struct {
}

func (Basic) Insert(model model.IModel) (uint64, error) {
	if err := model.Validate(); err != nil {
		return 0, err
	}
	if err := database.Connection.With().Create(model).Error; err != nil {
		return 0, err
	}
	return model.GetId(), nil
}

func (Basic) Update(model model.IModel) error {
	if err := model.Validate(); err != nil {
		return err
	}
	return database.Connection.With().Save(model).Error
}

func (Basic) Save(model model.IModel) (uint64, error) {
	if err := model.Validate(); err != nil {
		return 0, err
	}
	if err := database.Connection.With().Save(model).Error; err != nil {
		return 0, err
	}
	return model.GetId(), nil
}

func (Basic) SaveAll(models interface{}) error {

	if err := database.Connection.With().Save(models).Error; err != nil {
		return err
	}
	return nil
}

func (Basic) FindById(receiver model.IModel, id interface{}) error {
	return where("", nil).Statement.First(receiver, id).Error
}

func (Basic) FindFirst(receiver model.IModel, query string, args ...interface{}) error {
	return where(query, args...).Statement.Limit(1).Find(receiver).Error
}

func (Basic) FindAll(models interface{}, query string, args ...interface{}) (err error) {
	return where(query, args...).Statement.Find(models).Error
}

func (Basic) Delete(model model.IModel, query string, args ...interface{}) error {
	return where(query, args...).Statement.Delete(&model).Error
}

func where(query string, args ...interface{}) gorm.DB {
	tenantId := contexto.ContextoAutenticacao.GetTenantId()
	if tenantId == 0 {
		return *database.Connection.With().Where(query, args...)
	}

	if len(query) == 0 && len(args) == 0 {
		query = "tenant_id = ?"
	} else {
		query = query + " and tenant_id = ?"
	}
	args = append(args, tenantId)
	return *database.Connection.With().Where(query, args...)
}
