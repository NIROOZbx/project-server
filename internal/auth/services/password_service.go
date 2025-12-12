package services

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/NIROOZbx/project-server/internal/auth/repositories"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/NIROOZbx/project-server/internal/shared/cache"
	"github.com/NIROOZbx/project-server/internal/shared/email"
	"github.com/NIROOZbx/project-server/internal/shared/security"
	"github.com/redis/go-redis/v9"
)

// -----------------------------------------------------------
// ForgotPasswordService
// -----------------------------------------------------------
// STEP 1: User submits email to request a password reset.
// Security best practice: We NEVER reveal if the email exists or not.
//
// If email exists + verified → send OTP.
// If email does not exist or unverified → return nil (pretend success).
// ---
func ForgotPasswordService(ctx context.Context, toEmail string) error {

	otp, err := GenerateSecureOTP()

	if err != nil {
		return err
	}

	user, err := repositories.FindUserByEmail(toEmail)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user does not exist")
		}
		return err
	}

	if !user.IsVerified {
		return errors.New("account is not verified")
	}

	rdb := cache.GetClient()

	key := "reset:" + strconv.Itoa(int(user.ID))

	err = rdb.Set(ctx, key, otp, 3*time.Minute).Err()

	if err != nil {
		return errors.New("failed to store OTP")
	}

	return email.SendVerificationEmail(ctx, toEmail, otp)

}

// -----------------------------------------------------------
// VerifyResetPasswordService
// -----------------------------------------------------------
// STEP 2: User submits OTP for verification.
// If OTP is correct → return a RESET TOKEN (NOT an access token).
// This reset token will be used by ResetPasswordHandler.
// -----------------------------------------------------------
func VerifyResetPasswordService(ctx context.Context, email string, otp string) (string, error) {
	user, err := repositories.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid OTP or email")
	}

	rdb := cache.GetClient()

	key := "reset:" + strconv.Itoa(int(user.ID))

	storedOTP, err := rdb.Get(ctx, key).Result()

	if err == redis.Nil {
		// This means the key doesn't exist (it expired)
		return "", errors.New("OTP expired or invalid")
	}
	if err != nil {
		return "", err // A different server error
	}

	if storedOTP != otp {
		return "", errors.New("OTP is incorrect")
	}

	rdb.Del(ctx, key)

	resetToken, err := GenerateResetToken(email,user.ID)
	if err != nil {
		return "", err
	}
	

	return resetToken, nil

}

// -----------------------------------------------------------
// ChangePasswordService
// -----------------------------------------------------------
// STEP 3: Reset password using userID extracted from reset token.
//
// - Hash new password
// - Update DB
// - Increment token_version (logout all sessions)
// -----------------------------------------------------------

func ChangePasswordService(userId uint, password string) error {

	user, err := repositories.FindUserById(userId)

	if err != nil {

		if err == gorm.ErrRecordNotFound {
			return errors.New("user not found")
		}
		return err
	}
	compareErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if compareErr == nil {
		return errors.New("new password cannot be the same as the old password")
	}

	hashedPass, err := security.HashPassword(password)

	if err != nil {
		return err
	}

	err = repositories.UpdatePasswordAndVersion(userId, hashedPass)

	if err != nil {
		return err
	}

	return nil

}
