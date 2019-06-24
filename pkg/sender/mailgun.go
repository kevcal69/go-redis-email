package sender

import (
	"context"
	"fmt"
	"time"

	"github.com/kevcal69/go-redis-email/app/utils"
	"github.com/kevcal69/go-redis-email/pkg/models"
	"github.com/mailgun/mailgun-go"
)

var domain = utils.GetEnv("MAILGUN_DOMAIN", "")                 // e.g. mg.yourcompany.com
var privateAPIKey = utils.GetEnv("MAILGUN_PRIVATE_API_KEY", "") // mailgun api private key

func send(email *models.Email, callback func(success bool)) (string, error) {
	mg := mailgun.NewMailgun(domain, privateAPIKey)
	m := mg.NewMessage(
		email.FromAddress,
		email.Subject,
		email.Body,
		email.ToAddress,
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, id, err := mg.Send(ctx, m)
	fmt.Println(err)
	if err != nil {
		callback(false)
	} else {
		callback(true)
	}
	return id, err
}

// Mailgun : wrap email sender with go routine
func Mailgun(email *models.Email, callback func(success bool)) {
	go func(e *models.Email) {
		send(email, callback)
	}(email)
}
