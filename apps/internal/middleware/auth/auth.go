package auth

import (
	"github.com/Insiro/my_spotify/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type Middleware struct {
	privateData service.PrivateDataService
	user        service.UserService
}

func NewAuthMiddleware(privateData service.PrivateDataService, user service.UserService) *Middleware {
	return &Middleware{
		privateData: privateData,
		user:        user,
	}
}

func (m *Middleware) BaseLogged(c echo.Context, useQueryToken *bool) (*JwtUser, error) {
	// Get token from query parameter if enabled
	queryToken := c.QueryParam("token")

	if useQueryToken != nil && *useQueryToken && queryToken != "" {
		// TODO: Implement getUserFromField equivalent
		// user := getUserFromField("publicToken", queryToken, false)
		// if user != nil {
		//   return user
		// }
		return nil, nil
	}

	// Get token from cookie
	auth, err := c.Cookie("token")
	if err != nil {
		return nil, err
	}
	privateData, err := m.privateData.GetPrivateData()
	if err != nil {
		return nil, err
	}
	if privateData == nil || privateData.JwtPrivateKey == "" {
		return nil, errors.New("No private data found, cannot sign JWT")
	}
	jwtUser := &JwtUser{}
	claims, err := parseToken(auth.Value, privateData.JwtPrivateKey)
	if err != nil {
		return nil, err
	}

	jwtUser.UserId = claims["userId"].(string)

	if jwtUser.UserId == "" {
		return nil, nil
	}

	user, err := m.user.GetUserFromField("id", jwtUser.UserId, false)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}
	return jwtUser, nil
}
