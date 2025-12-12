package email

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/NIROOZbx/project-server/internal/shared/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendVerificationEmail(ctx context.Context,toEmail, otp string) error {


	cfg := config.GetConfig()
	from:=mail.NewEmail(cfg.SenderName,cfg.SenderMail)
	to:=mail.NewEmail("",toEmail)

	subject := "Your OTP Verification Code"

	plainText := fmt.Sprintf("Your OTP code is: %s", otp)

	htmlContent := fmt.Sprintf("<strong>Your verification code is: %s</strong><p>It will expire in 1 minutes.</p>", otp)

	message := mail.NewSingleEmail(from, subject, to, plainText, htmlContent)

	client := sendgrid.NewSendClient(cfg.SendGridAPIKey)

   response, err := client.SendWithContext(ctx,message)

	if err != nil {
		return err
	}

	if response.StatusCode!=202{
		log.Printf("Failed to send email: %s", response.Body)
		return errors.New("failed to send verification email")
	}


	log.Printf("Verification email sent to %s", toEmail)
	return nil



}
