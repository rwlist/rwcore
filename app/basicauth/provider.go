package basicauth

import (
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo/bson"

	"golang.org/x/crypto/bcrypt"

	"github.com/rwlist/rwcore/app/db"
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

type Provider struct {
	baseDB *db.Provider
}

func New(db *db.Provider) *Provider {
	return &Provider{db}
}

func (a Provider) FindUser(r *http.Request) (interface{}, error) {
	decoder := json.NewDecoder(r.Body)
	var form LoginForm
	err := decoder.Decode(&form)
	if err != nil {
		return nil, err
	}
	return a.HandleLogin(form)
}

func (a Provider) CreateUser(r *http.Request) (interface{}, error) {
	decoder := json.NewDecoder(r.Body)
	var form SignUpForm
	err := decoder.Decode(&form)
	if err != nil {
		return nil, err
	}
	return a.HandleSignUp(form)
}

func (a Provider) HandleLogin(form LoginForm) (interface{}, error) {
	DB := a.baseDB.Copy()
	defer DB.Close()

	user, err := DB.Users().FindByUsername(form.Username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(form.Password))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (a Provider) HandleSignUp(form SignUpForm) (interface{}, error) {
	DB := a.baseDB.Copy()
	defer DB.Close()

	hashed, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &db.User{
		ID:             bson.NewObjectId(),
		Username:       form.Username,
		HashedPassword: hashed,
		Email:          form.Email,
		FirstName:      form.FirstName,
		SecondName:     form.SecondName,
	}

	err = DB.Users().InsertOne(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
