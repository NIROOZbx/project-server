package services

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
	"strconv"
	"time"

	
	"github.com/NIROOZbx/project-server/internal/auth/repositories"
	"github.com/NIROOZbx/project-server/internal/shared/cache"
	"github.com/NIROOZbx/project-server/internal/shared/email"
	"github.com/redis/go-redis/v9"
)

func GenerateSecureOTP() (string, error) {
	otp := ""
	for i := 0; i < 6; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		otp += num.String()
	}
	return otp, nil
}

func SendOTP(ctx context.Context,toEmail string) error {

	user, err := repositories.FindUserByEmail(toEmail)
	if err != nil {
		return errors.New("failed to process request")
	}


	otp, err := GenerateSecureOTP()

	if err != nil {
		return err
	}

	rdb := cache.GetClient()

	key := "otp:" + strconv.Itoa(int(user.ID))


	err = rdb.Set(ctx, key, otp, 6*time.Minute).Err()

	if err != nil {
		return errors.New("failed to store OTP")
	}

	return email.SendVerificationEmail(ctx, user.Email, otp)

}

func VerifyOTP(ctx context.Context, email string, otp string) error {

	user, err := repositories.FindUserByEmail(email)
	if err != nil {
		return errors.New("failed to process request")
	}

	rdb := cache.GetClient()

	key := "otp:" + strconv.Itoa(int(user.ID))

	storedOTP, err := rdb.Get(ctx, key).Result()

	if err == redis.Nil {
		
		return errors.New("OTP expired or invalid")
	}
	if err != nil {
		return err 
	}

	if storedOTP != otp {
		return errors.New("OTP is incorrect")
	}

	if err := repositories.SetUserVerified(user.ID); err != nil {
		return err
	}

	rdb.Del(ctx, key)

	return nil

}
