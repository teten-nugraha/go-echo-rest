package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})

func JWTBypass(next echo.HandlerFunc) echo.HandlerFunc {

	//A dummy handler func
	handler := func(c echo.Context) error {
		return nil
	}

	return func(c echo.Context) error {
		//use the default JWT middleware and feed in our
		//dummy handler
		handler = middleware.JWTWithConfig(getJWTConfig())(handler)
		//Execute the JWT validation. Since this will return nil
		//and without errors, we can continue
		handler(c)

		//Read the value set by the JWT middleware and authenticate by
		//checking if the jwtConfig object is set or not in the context
		storeVal := c.Get("jwtConfig")

		//Go to the next handler
		return next(c)
	}
}