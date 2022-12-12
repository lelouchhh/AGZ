package repository

import (
	"AGZ/pkg/structures"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"os/exec"
	"strings"
)

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

type AuthPostgres struct {
	db *sqlx.DB
}

func (r *AuthPostgres) ResetUserEmailPass(user structures.Params) error {
	query := fmt.Sprintf("call auth.reset_s	et_reg_email('%s')", user.EmailReg.Email)
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}
	{
		path := viper.GetString("tools.emailConfirm.path")
		out, err := exec.Command(path, "-email", user.EmailReg.Email).
			Output()
		if err != nil {
			return err
		}
		fmt.Println(string(out))
	}
	return err
}
func (r *AuthPostgres) RegisterUserEmail(user structures.Params) error {
	user.EmailReg.Email = strings.ToLower(user.EmailReg.Email)
	query := fmt.Sprintf("call auth.set_reg_email('%s', '%s', '%s')", user.EmailReg.Login, user.EmailReg.Pass, user.EmailReg.Email)
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}
	{
		path := viper.GetString("tools.emailConfirm.path")
		out, err := exec.Command(path, "-email", user.EmailReg.Email).
			Output()
		if err != nil {
			return err
		}
		fmt.Println(string(out))
	}
	return err
}
func (r *AuthPostgres) ConfirmUserEmail(user structures.Params) error {
	query := fmt.Sprintf("call auth.check_auth_email_reg('%s', '%s')", user.EmailReg.Login, user.EmailReg.Hash)
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}

	return err
}
func (r *AuthPostgres) SignIn(user structures.Params) (structures.Tokens, error) {
	query := fmt.Sprintf("select * from auth.authentication_user('%s', '%s', '%s')", user.User.Login, user.User.Pass, user.User.DevInfo)
	var rows structures.Tokens
	err := r.db.Get(&rows, query)
	if err != nil {
		return rows, err
	}

	return rows, err
}

//	func (r *AuthPostgres) ActionUserAccess(user structures.Tokens) error {
//		return nil
//	}
//
//	func (r *AuthPostgres) ActionUserRefresh(user structures.Tokens) (structures.Tokens, error) {
//		return structures.Tokens{}, nil
//	}

//	func (r *AuthPostgres) AuthenticationUser(user structures.User) (structures.Tokens, error) {
//		query := fmt.Sprintf("select * from auth.authentication_user('%s', '%s', '%s')", user.Name, user.Password, user.DeviceId)
//		var rows structures.Tokens
//		fmt.Println(query)
//		err := r.db.Get(&rows, query)
//		if err != nil {
//			return rows, err
//		}
//
//		return rows, err
//	}
func (r *AuthPostgres) ActionUserRefresh(token structures.Tokens) (structures.Tokens, error) {
	query := fmt.Sprintf("call auth.action_user_refresh('%s', '%s', '%s')", token.Access, token.Refresh, token.DevInfo)
	var rows structures.Tokens
	err := r.db.Get(&rows, query)
	if err != nil {
		return rows, err
	}

	return rows, err
}
func (r *AuthPostgres) ActionUserAccess(tokens structures.Tokens) error {
	query := fmt.Sprintf("call auth.action_user_access('%s', '%s')", tokens.Access, tokens.DevInfo)
	fmt.Println(query)
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}

	return err
}
