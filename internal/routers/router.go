package routers

import (
	v1 "GameAdmin/api/v1"
	"GameAdmin/internal/middleware"
	"GameAdmin/internal/middleware/jwt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var _ IRouter = (*Router)(nil)

var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

type IRouter interface {
	Register(app *gin.Engine) error
	Prefixes() []string
}

type Router struct {
	MyJwt     *jwt.JWT
	UserApi   *v1.UserApi
	SystemApi *v1.SystemApi
	FileApi   *v1.FileApi
	MsgApi    *v1.MsgApi
}

func (a *Router) Register(app *gin.Engine) error {
	a.RegisterAPI(app)
	return nil
}

func (a *Router) Prefixes() []string {
	return []string{
		"/api/",
	}
}

// RegisterAPI register api group router
func (a *Router) RegisterAPI(app *gin.Engine) {
	g := app.Group("/api")
	g.Use(middleware.RateLimiterMiddleware())

	v1 := g.Group("/v1")
	v1.POST("/uploadFile", a.FileApi.UploadFile)

	var authMiddleware = a.MyJwt.GinJWTMiddlewareInit(&jwt.AllUserAuthorizator{})
	app.NoRoute(authMiddleware.MiddlewareFunc(), middleware.NoRouteHandler())

	v1.POST("/login", authMiddleware.LoginHandler)
	v1.GET("/refreshToken", authMiddleware.RefreshHandler)

	gUser := v1.Group("/user")
	gUser.Use(authMiddleware.MiddlewareFunc())
	{
		gUser.GET("/info", a.UserApi.GetUserInfo)
		gUser.POST("/logout", authMiddleware.LogoutHandler)
	}

	gSystem := v1.Group("system")
	gSystem.Use(authMiddleware.MiddlewareFunc())
	{
		gSystem.GET("/serverInfo", a.SystemApi.GetServerInfo)
		gSystem.GET("/reloadSystem", a.SystemApi.ReloadSystem)
	}

	var adminMiddleware = a.MyJwt.GinJWTMiddlewareInit(&jwt.AdminAuthorizator{})
	gUser.Use(adminMiddleware.MiddlewareFunc())
	{
		gUser.GET("/list", a.UserApi.GetUsers)
		gUser.POST("", a.UserApi.AddUser)
		gUser.PUT("", a.UserApi.UpdateUser)
		gUser.DELETE(":id", a.UserApi.DeleteUser)
	}

	v1.Group("/deleteFile").Use(adminMiddleware.MiddlewareFunc()).DELETE(":id", a.FileApi.DeleteFile)

	gMsg := v1.Group("msg")
	gMsg.GET("/notice", a.MsgApi.GetNotices)
	gMsg.Use(authMiddleware.MiddlewareFunc())
	{
		gMsg.POST("/notice", a.MsgApi.AddNotice)
		gMsg.PUT("/notice", a.MsgApi.UpdateNotice)
		gMsg.DELETE("/notice/:id", a.MsgApi.DeleteNotice)
	}
}
