package repositories

import (
	"errors"
	

	"github.com/NIROOZbx/project-server/internal/auth/models"
	"github.com/NIROOZbx/project-server/internal/shared/database"
	"gorm.io/gorm"
)

// FindDuplicateUser checks if a user with the same email OR name already exists.
// It returns an error ONLY when a record *is found*.
// If no record exists, GORM returns gorm.ErrRecordNotFound → meaning it's safe to create the user.
func FindDuplicateUser(email, name string) bool {
	
	var db = database.DB()
	var user models.User

	err := db.Where("email = ? OR name = ?", email, name).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return err == nil
}

// CreateUser inserts a new user into the database.
// Returns nil if user was successfully created.
// Returns an error if the insert fails (validation errors, duplicate email, etc).
func CreateUser(newUser *models.User) error {
	var db = database.DB()
	return db.Create(newUser).Error

}

// FindUserByEmail retrieves a user by their email address.
// If the user exists → returns (user, nil)
// If user not found → returns (nil, gorm.ErrRecordNotFound)
// If any database error occurs → returns (nil, error)
func FindUserByEmail(email string) (*models.User, error) {

	var db = database.DB()
	var user models.User

	err := db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil

}

// FindUserById retrieves a single user by ID.
// Used for authenticated operations (fetch profile, reset password, etc).
// Returns a pointer to User struct or an error.
func FindUserById(userId uint) (*models.User, error) {

	var db = database.DB()
	var user models.User

	err := db.Where("id = ?", userId).First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil

}

// UpdateTokenVersion increments the user’s token_version field by 1.
// This is used for **Token Versioning** — a method to instantly logout a user everywhere.
// Once the version changes, all old tokens become invalid.
func UpdateTokenVersion(userId uint) error {

	var db = database.DB()

	return db.Model(&models.User{}).Where("id = ?", userId).Update("token_version", gorm.Expr("token_version + 1")).Error

}

// SetUserVerified marks the user's email as verified.
// Called after OTP verification during registration.
// After this, user can login normally.

func SetUserVerified(userId uint) error {

	var db = database.DB()

	return db.Model(&models.User{}).Where("id = ?", userId).Update("is_verified", true).Error

}

// Used to update the token version and user password when uer clicks forgot password
func UpdatePasswordAndVersion(userId uint, newHashedPassword string) error {

	var db = database.DB()

	return db.Model(&models.User{}).Where("id = ?", userId).Updates(map[string]any{
		"password":      newHashedPassword,
		"token_version": gorm.Expr("token_version + 1"),
	}).Error

}

func UpdatePassword(userId uint, newHashedPassword string) error {

	var db = database.DB()

	return db.Model(&models.User{}).Where("id = ?", userId).Update("password", newHashedPassword).Error

}
