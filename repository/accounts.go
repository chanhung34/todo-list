package repository

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
	"todo_list/model"
	"todo_list/storage"
)

type AccountRepository interface {
	Validate() (map[string]interface{}, bool)
	SignUp(ctx context.Context, newAccount *model.Account) (*model.Account, error)
	GetUser(ctx context.Context, userId uint) *model.Account
	Auth(ctx *gin.Context, username string, password string) (*model.Account, error)
	VerifyAuthenticate(ctx *gin.Context) (*model.Account, error)
}

type account struct {
	storage storage.Account
	logger  logrus.Logger
}

func NewAccountRepository(sa storage.Account, logger logrus.Logger) *account {
	return &account{
		storage: sa,
		logger:  logger,
	}
}

func (repo account) Validate() (map[string]interface{}, bool) {
	panic("implement me")
}

func (repo account) SignUp(ctx context.Context, newAccount *model.Account) (*model.Account, error) {
	isExistAccount, err := repo.storage.ExistUserName(ctx, newAccount.UserName)
	if err != nil {
		return nil, err
	}
	if isExistAccount {
		return nil, errors.New("User already ")
	}
	now := time.Now()
	newAccount.CreatedAt = &now
	newAccount.UpdatedAt = &now
	newAccount, err2 := repo.storage.Create(ctx, newAccount)
	if err2 != nil {
		return nil, err2
	}
	return newAccount, nil

}

func (repo account) GetUser(ctx context.Context, userId uint) *model.Account {
	panic("implement me")
}

func (repo account) VerifyAuthenticate(ctx *gin.Context) (*model.Account, error) {
	// We can obtain the session token from the requests cookies, which come with every request
	var authAccount model.Account
	tknStr, err := ctx.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			return nil, err
		}
		// For any other type of error, return a bad request status
		return nil, errors.New("bad request")
	}

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !tkn.Valid {
		return nil, errors.New("token invalid")
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("token invalid")
		}
		return nil, errors.New("token invalid")
	}

	// Finally, return the welcome message to the user, along with their
	// username given in the token
	authAccount.ID = claims.UserId
	return &authAccount, nil
}

func (repo account) Auth(ctx *gin.Context, username string, password string) (*model.Account, error) {

	acc, err := repo.storage.GetByUserNameAndPassword(context.Background(), username, password)
	if err != nil {
		return nil, errors.New("Wrong username or password")
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		UserId: acc.ID,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, errors.New("Internal Error")
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	ctx.SetCookie(
		"token",
		tokenString,
		int(expirationTime.Unix()),
		"/",
		"http://localhost:8080",
		false,
		false,
	)
	return acc, nil
}

var jwtKey = []byte("very_secret_key")

var users = map[string]string{
	"hungvtc": "123abc",
	"user2":   "password2",
}

// Create a struct to read the username and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	UserId int `json:"username"`
	jwt.StandardClaims
}
