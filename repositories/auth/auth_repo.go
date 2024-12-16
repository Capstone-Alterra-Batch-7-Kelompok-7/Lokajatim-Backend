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

func (authRepo *AuthRepo) StoreOTP(email, otp string) error {
	// Assuming the user table has fields "otp" and "otp_expiration"
	return authRepo.db.Model(&User{}).Where("email = ?", email).Updates(map[string]interface{}{
		"otp":             otp,
	}).Error
}


func (authRepo *AuthRepo) VerifyOTP(email, otp string) (bool, error) {
    var userOTP entities.User // Assuming UserOTP is the struct that holds the OTP and expiration time
    err := authRepo.db.Where("email = ?", email).First(&userOTP).Error
    if err != nil {
        return false, err // Will return if no record is found
    }


    // Verify if the OTP matches
    if userOTP.OTP != otp {
        return false, errors.New("invalid OTP")
    }

    return true, nil
}


func (authRepo *AuthRepo) UpdatePassword(user entities.User) error {
	if err := authRepo.db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (authRepo *AuthRepo) GetAllUsers() ([]entities.User, error) {
	var users []entities.User
	err := authRepo.db.Find(&users).Error
	return users, err
}

func (authRepo *AuthRepo) UpdateUser(user entities.User) (entities.User, error) {
	err := authRepo.db.Model(&entities.User{}).Where("id = ?", user.ID).Updates(&user).Error
	return user, err
}

func (authRepo *AuthRepo) DeleteUser(userID int) error {
	err := authRepo.db.Delete(&entities.User{}, userID).Error
	return err
}
