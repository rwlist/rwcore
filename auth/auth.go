package auth

import (
	"github.com/globalsign/mgo/bson"
	"github.com/rwlist/rwcore/users"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct{
	jwt *JWT
	stdClaims *StdClaims
}

func NewAuth(jwt *JWT, stdClaims *StdClaims) *Auth {
	return &Auth{
		jwt: jwt,
		stdClaims: stdClaims,
	}
}

func (a *Auth) Login(db users.Store, form LoginForm) (users.User, error) {
	user, err := db.FindByUsername(form.Username)
	if err != nil {
		return users.User{}, err
	}

	err = bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(form.Password))
	if err != nil {
		return users.User{}, err
	}

	return user, nil
}

func (a *Auth) SignUp(db users.Store, form SignUpForm) (users.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		return users.User{}, err
	}

	roles := make(users.Roles).AddRole(users.UserRole)

	count, err := db.Count()
	if err != nil {
		return users.User{}, err
	}

	if count == 0 {
		roles = roles.AddRole(users.AdminRole) // Set up admin if no one is there
	}

	user := users.User{
		ID:             bson.NewObjectId(),
		Username:       form.Username,
		HashedPassword: hashed,
		Email:          form.Email,
		FirstName:      form.FirstName,
		SecondName:     form.SecondName,
		Roles:          roles,
	}

	err = db.InsertOne(user)
	if err != nil {
		return users.User{}, err
	}

	return user, nil
}

func (a *Auth) CreateClaims(user users.User) Claims {
	return Claims{
		User: user,
		StandardClaims: a.stdClaims.Create(),
	}
}