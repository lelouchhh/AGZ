package repository

import (
	"AGZ/pkg/structures"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	RegisterUserEmail(user structures.Params) error
	ConfirmUserEmail(user structures.Params) error
	ResetUserEmailPass(user structures.Params) error
	SignIn(user structures.Params) (structures.Tokens, error)
	ActionUserAccess(token structures.Tokens) error
	ActionUserRefresh(token structures.Tokens) (structures.Tokens, error)
}
type Profile interface {
	AddPurchase(user structures.Params) error
	RemovePurchase(user structures.Params) error
	GetBasket(user structures.Params) ([]structures.Purchases, error)
	GetLinksBasket(user structures.Params) ([]structures.Links, error)
	AddLink(user structures.Params) error
	RemoveLink(user structures.Params) error
	AddCustomNote(user structures.Params) error
	AddCustomStatus(user structures.Params) error
	GetProfileInfo(token structures.Tokens) (structures.UserPersonalInfo, error)
	AddProfileInfo(user structures.UserPersonalInfo) error
}

type Repository struct {
	Authorization
	Profile
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Profile:       NewProfilePostgres(db),
	}
}
