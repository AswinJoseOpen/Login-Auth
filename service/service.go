package service

import (
	"fmt"
	"time"

	"github.com/AswinJoseOpen/Login-Auth/config"
	"github.com/AswinJoseOpen/Login-Auth/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ServiceImpl struct {
	DB     *gorm.DB
	Config config.AppConfig
}

func NewServiceImpl(db *gorm.DB, config config.AppConfig) *ServiceImpl {
	testServiceImpl := &ServiceImpl{
		DB:     db,
		Config: config,
	}
	return testServiceImpl
}

func (s *ServiceImpl) TestService(c *gin.Context) *model.Message {

	fmt.Println("testing service")
	user, _ := c.Get("user")
	fmt.Println(user)
	resp := model.NewMessageResponse("testing")
	return resp
}

func (s *ServiceImpl) SignUp(c *gin.Context, req *model.Users) error {
	//password hashing]
	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	//Create the user
	user := &model.Users{
		Email:    req.Email,
		Password: string(pass),
	}
	result := s.DB.Create(user)
	if result.Error != nil {
		return fmt.Errorf("failed to create user")
	}
	return nil
}

func (s *ServiceImpl) Login(c *gin.Context, req *model.Users) (*model.TokenResponse, error) {
	var user model.Users
	// look up requested user
	s.DB.First(&user, "email = ?", req.Email)
	if user.ID == 0 {
		return nil, fmt.Errorf("user does not exist")
	}

	// Compare password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("incorrect password")
	}

	// Generate jwt Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(s.Config.Secret))
	if err != nil {
		return nil, fmt.Errorf("unable to generate token")
	}
	// c.SetSameSite(http.SameSiteLaxMode)
	// c.SetCookie()

	//respond
	return model.NewTokenResponse(tokenString), nil
}
