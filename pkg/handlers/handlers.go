package handlers

import (
	"net/http"

	"github.com/kevcal69/go-redis-email/app/db"
	"github.com/kevcal69/go-redis-email/pkg/models"
	"github.com/kevcal69/go-redis-email/pkg/sender"
	"github.com/labstack/echo"
)

var client = db.RedisClient()

// api view handler
func sendEmailView(c echo.Context) (err error) {
	email := new(models.Email)
	if err = c.Bind(email); err != nil {
		return err
	}
	email.Save(client)
	sender.Mailgun(email, func(success bool) {
		if success {
			email.MarkSent(client)
		} else {
			email.MarkFail(client)
		}
	})
	return c.JSON(http.StatusOK, email)
}

func fetchEmails(c echo.Context) error {
	email := new(models.Email)
	data := email.GetAll(client)
	return c.JSON(200, data)
}

type apiView struct {
	Handler func(echo.Context) error
	Route   string
	Method  []string
}

// APIRoutes : list of registered api views
var APIRoutes = []apiView{
	apiView{
		Route:   "/emails/create",
		Handler: sendEmailView,
		Method:  []string{"POST"},
	},
	apiView{
		Route:   "/emails",
		Handler: fetchEmails,
		Method:  []string{"GET"},
	},
}
