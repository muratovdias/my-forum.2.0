package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/muratovdias/my-forum.2.0/models"
	"gorm.io/gorm"
)

type AuthRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (r *AuthRepo) CreateUser(user models.User) error {
	result := r.db.Create(&user)
	if result.Error != nil {
		log.Printf("repo: create user: %s", result.Error)
		return fmt.Errorf(path+"create user: %w", result.Error)
	}
	return nil
}

func (r *AuthRepo) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	row := r.db.Select("username", "password").Where("email= ? ", email).Find(&user)
	if row.Error != nil {
		return user, fmt.Errorf(path+"get user by email: %w", err)
	}
	return user, nil
}

func (r *AuthRepo) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	row := r.db.Select("username", "password").Where("username= ? ", username).Find(&user)
	if row.Error != nil {
		return user, fmt.Errorf(path+"get user by username: %w", err)
	}
	return user, nil
}

func (r *AuthRepo) GetUserByToken(token string) (models.User, error) {
	var user models.User
	row := r.db.Where("token= ? ", token).Find(&user)
	if row.Error != nil {
		return user, fmt.Errorf(path+"get user by token: %w", row.Error)
	}
	return user, nil
}

func (r *AuthRepo) SaveToken(username, token string, duration time.Time) error {
	row := r.db.Model(&models.User{}).Where("username= ?", username).Updates(models.User{Token: token, TokenDuration: duration})
	if row.Error != nil {
		return fmt.Errorf("ERROR: /repository save token: %w", row.Error)
	}
	return nil
}

func (r *AuthRepo) DeleteToken(token string) error {
	row := r.db.Model(&models.User{}).Where("token= ?", token).Updates(map[string]interface{}{"token": "", "token_duration": time.Time{}})
	if row.Error != nil {
		return row.Error
	}
	return nil
}
