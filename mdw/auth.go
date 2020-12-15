package mdw

import "github.com/labstack/echo/v4"

func BasicAuth1(username string, password string, c echo.Context) (bool, error) {
	// if username != "admin" || password != "789" {
	// 	return false, nil
	// }

	if username == "admin" && password == "789" {
		c.Set("username", username)
		c.Set("admin", true)

		return true, nil
	}

	if username == "Khoa" && password == "789" {
		c.Set("username", username)
		c.Set("admin", false)

		return true, nil
	}

	// Set key "username" in c equal to username that user Post through API

	return false, nil
}
