package main

import (
	"EchoFrameWork/handler"
	"EchoFrameWork/mdw"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	server := echo.New()

	server.Use(middleware.Logger())

	isLogedIn := middleware.JWT([]byte("mysecretkey"))
	isAdmin := mdw.IsAdminMdw

	server.GET("/", handler.Hello, isLogedIn)

	server.POST("/login", handler.Login, middleware.BasicAuth(mdw.BasicAuth1))

	server.GET("/admin", handler.Hello, isLogedIn, isAdmin)

	groupv2 := server.Group("/v2")
	groupv2.GET("/hello", handler.Hello2)

	groupUser := server.Group("/api/user", isLogedIn)
	groupUser.GET("/get", handler.GetUser)
	groupUser.GET("/update", handler.UpdateUser, isAdmin)
	groupUser.GET("/delete", handler.DeleteUser, isAdmin)
	server.Logger.Fatal(server.Start(":8888"))
}

// func hello(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello World!")
// }

// type LoginRequest struct {
// 	UserName string `json:"username" form:"username" xml:"username" query:"username"`
// 	PassWord string `json:"password" form:"password" xml:"password" query:"password"`
// }
