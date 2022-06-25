package repository

import (
	"api-merca/src/contexto"
	"api-merca/src/database"
	"api-merca/src/model"

	"gorm.io/gorm"
)

type Basic struct {
}

func (Basic) Insert(model model.IModel) (id uint64, err error) {
	if err = model.Validate(); err != nil {
		return
	}
	if err = getBancoDados().Create(model).Error; err == nil {
		id = model.GetId()
	}
	return
}

func (Basic) Update(model model.IModel) (err error) {
	if err = model.Validate(); err == nil {
		err = getBancoDados().Save(model).Error
	}
	return
}

func (Basic) Save(model model.IModel) (id uint64, err error) {
	if err = model.Validate(); err != nil {
		return
	}
	if err = getBancoDados().Save(model).Error; err == nil {
		id = model.GetId()
	}
	return
}

func (Basic) SaveAll(models interface{}) (err error) {
	err = getBancoDados().Save(models).Error
	return
}

func (Basic) FindById(receiver model.IModel, id interface{}) (err error) {
	err = where("", nil).Statement.First(receiver, id).Error
	return
}

func (Basic) FindFirst(receiver model.IModel, query string, args ...interface{}) (err error) {
	err = where(query, args...).Statement.Limit(1).Find(receiver).Error
	return
}

func (Basic) FindAll(models interface{}, query string, args ...interface{}) (err error) {
	err = where(query, args...).Statement.Find(models).Error
	return
}

func (Basic) Delete(model model.IModel, query string, args ...interface{}) (err error) {
	err = where(query, args...).Statement.Delete(&model).Error
	return
}

func where(query string, args ...interface{}) *gorm.DB {
	tenantId := contexto.ContextoAutenticacao.GetTenantId()
	if tenantId == 0 {
		return getBancoDados().Where(query, args...)
	}

	if len(query) == 0 && len(args) == 0 {
		query = "tenant_id = ?"
	} else {
		query = query + " and tenant_id = ?"
	}
	args = append(args, tenantId)
	return getBancoDados().Where(query, args...)
}

func getBancoDados() (bancoDados *gorm.DB) {
	return database.Connection.WithContext()
}
