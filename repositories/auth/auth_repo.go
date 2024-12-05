package auth

import (
	"errors"
	"fmt"
	"lokajatim/entities"

	"gorm.io/gorm"
)

func NewAuthRepo(db *gorm.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

type AuthRepo struct {
	db *gorm.DB
}

func (authRepo AuthRepo) Login(user entities.User) (entities.User, error) {
	userDb := FromEntities(user)
	result := authRepo.db.First(&userDb, "email = ? AND password = ?", userDb.Email, userDb.Password)
	if result.Error != nil {
		return entities.User{}, result.Error
	}
	return userDb.ToEntities(), nil
}

func (authRepo AuthRepo) Register(user entities.User) (entities.User, error) {
	if err := authRepo.db.Create(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (authRepo AuthRepo) GetUserByEmail(email string) (entities.User, error) {
	var user entities.User
	if err := authRepo.db.First(&user, "email = ?", email).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (authRepo AuthRepo) GetLastUserID() (int, error) {
	var user entities.User
	err := authRepo.db.Last(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 1, nil
		}
		return 0, err
	}
	return user.ID, nil
}

func (r *AuthRepo) GetUserByID(userID int) (entities.User, error) {
	var user entities.User
	result := r.db.First(&user, userID)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return entities.User{}, fmt.Errorf("user not found")
	}
	if result.Error != nil {
		return entities.User{}, result.Error
	}

	return user, nil
}

func (authRepo *AuthRepo) StoreOTP(email string, otp string) error {
	var user entities.User
	if err := authRepo.db.First(&user, "email = ?", email).Error; err != nil {
		return err
	}
	user.OTP = otp
	if err := authRepo.db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (authRepo *AuthRepo) VerifyOTP(email string, otp string) (bool, error) {
	var user entities.User
	if err := authRepo.db.First(&user, "email = ?", email).Error; err != nil {
		return false, err
	}
	if user.OTP != otp {
		return false, fmt.Errorf("invalid OTP")
	}
	return true, nil
}

func (authRepo *AuthRepo) UpdatePassword(user entities.User) error {
	if err := authRepo.db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
