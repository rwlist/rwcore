package auth

import (
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/rwlist/rwcore/app/db"
	"github.com/rwlist/rwcore/app/model"
	"golang.org/x/crypto/bcrypt"
)

const (
	AdminRole = "admin"
	UserRole  = "user"
)

type LoginForm struct {
	Username string
	Password string
}

type SignUpForm struct {
	Username   string
	Email      string
	Password   string
	FirstName  string
	SecondName string
}

type Users struct{}

func (u Users) FindUser(r *http.Request) (*model.User, error) {
	decoder := json.NewDecoder(r.Body)
	var form LoginForm
	err := decoder.Decode(&form)
	if err != nil {
		return nil, err
	}
	return u.HandleLogin(r, form)
}

func (u Users) CreateUser(r *http.Request) (*model.User, error) {
	decoder := json.NewDecoder(r.Body)
	var form SignUpForm
	err := decoder.Decode(&form)
	if err != nil {
		return nil, err
	}
	return u.HandleSignUp(r, form)
}

func (u Users) HandleLogin(r *http.Request, form LoginForm) (*model.User, error) {
	db := db.From(r)

	user, err := db.Users().FindByUsername(form.Username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(form.Password))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u Users) HandleSignUp(r *http.Request, form SignUpForm) (*model.User, error) {
	db := db.From(r)

	hashed, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	roles := make(model.Roles).AddRole(UserRole)

	count, err := db.Users().Count()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		roles = roles.AddRole(AdminRole) // Set up admin if no one is there
	}

	user := &model.User{
		ID:             bson.NewObjectId(),
		Username:       form.Username,
		HashedPassword: hashed,
		Email:          form.Email,
		FirstName:      form.FirstName,
		SecondName:     form.SecondName,
		Roles:          roles,
	}

	err = db.Users().InsertOne(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
