package basicauth

import (
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/rwlist/rwcore/app/db"
	"github.com/rwlist/rwcore/app/model"
	"golang.org/x/crypto/bcrypt"
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

type Provider struct{}

func (a Provider) FindUser(r *http.Request) (interface{}, error) {
	decoder := json.NewDecoder(r.Body)
	var form LoginForm
	err := decoder.Decode(&form)
	if err != nil {
		return nil, err
	}
	return a.HandleLogin(r, form)
}

func (a Provider) CreateUser(r *http.Request) (interface{}, error) {
	decoder := json.NewDecoder(r.Body)
	var form SignUpForm
	err := decoder.Decode(&form)
	if err != nil {
		return nil, err
	}
	return a.HandleSignUp(r, form)
}

func (a Provider) HandleLogin(r *http.Request, form LoginForm) (interface{}, error) {
	db := r.Context().Value(db.DBKey).(*db.Provider)

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

func (a Provider) HandleSignUp(r *http.Request, form SignUpForm) (interface{}, error) {
	db := r.Context().Value(db.DBKey).(*db.Provider)

	hashed, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	roles := make(model.Roles).AddRole("user")

	count, err := db.Users().Size()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		roles = roles.AddRole("admin")
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
