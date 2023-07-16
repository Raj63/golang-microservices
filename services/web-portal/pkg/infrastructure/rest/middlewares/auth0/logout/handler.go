package logout

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// DI is the logout handler entity
type DI struct {
	Auth0ClientID string
	Auth0Domain   string
}

// Handler for our logout.
func Handler(di *DI) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logoutURL, err := url.Parse("https://" + di.Auth0Domain + "/v2/logout")
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		scheme := "http"
		if ctx.Request.TLS != nil {
			scheme = "https"
		}

		returnTo, err := url.Parse(scheme + "://" + ctx.Request.Host + "/v1/login")
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		parameters := url.Values{}
		parameters.Add("returnTo", returnTo.String())
		parameters.Add("client_id", di.Auth0ClientID)
		logoutURL.RawQuery = parameters.Encode()

		ctx.Redirect(http.StatusTemporaryRedirect, logoutURL.String())
	}
}
