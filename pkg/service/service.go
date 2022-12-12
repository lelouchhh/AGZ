package service

import (
	"AGZ/pkg/repository"
	"AGZ/pkg/structures"
)

type Authorization interface {
	ActionUserAccess(token structures.Tokens) (err error)
	ActionUserRefresh(token structures.Tokens) (structures.Tokens, error)

	RegisterUserEmail(user structures.Params) error
	ConfirmUserEmail(user structures.Params) error

	ResetUserEmailPass(user structures.Params) error

	SignIn(user structures.Params) (structures.Tokens, error)
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
	AddProfileInfo(profile structures.UserPersonalInfo) error
}
type Service struct {
	Authorization
	Profile
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Profile:       NewProfileService(repos.Profile),
	}
}
