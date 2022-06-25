package model

type IModel interface {
	GetId() uint64

	SetTenant()

	Validate() error
}
