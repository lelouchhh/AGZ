package handler

import (
	"AGZ/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(CORSMiddleware())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		signUp := router.Group("auth/sign_up")
		{
			signUp.POST("/register_user_email", h.RegisterUserEmail)
			signUp.POST("/confirm_user_email", h.ConfirmUserEmail)
			signUp.POST("/reset_user_email_pass", h.ResetUserEmailPass)

		}
		auth.POST("/sign_in", h.SignIn)
		auth.POST("/action_access", h.ActionUserAccess)
		auth.POST("/action_refresh", h.ActionUserRefresh)
		//auth.POST("/sign_confirm", h.RefConfirm)
	}
	profile := router.Group("/profile")
	{
		profile.POST("/add_purchase", h.AddPurchase)
		profile.POST("/remove_purchase", h.RemovePurchase)
		profile.POST("/get_basket", h.GetBasket)
		profile.POST("/remove_link", h.RemoveLink)
		profile.POST("/add_link", h.AddLink)
		profile.POST("/get_link_basket", h.GetLinkBasket)
		profile.POST("/add_custom_note", h.AddCustomNote)
		profile.POST("/add_custom_status", h.AddCustomStatus)
		profile.POST("/get_profile_info", h.GetProfileInfo)
		profile.POST("/add_profile_info", h.AddProfileInfo)
	}
	//token := router.Group("/token")
	//{
	//	token.POST("/action_access", h.ActionUserAccess)
	//	token.POST("/action_refresh", h.ActionUserRefresh)
	//}
	err := router.RunTLS(":8080", "/etc/letsencrypt/live/www.969975-cv27771.tmweb.ru/fullchain.pem", "/etc/letsencrypt/live/www.969975-cv27771.tmweb.ru/privkey.pem")
	if err != nil {
		fmt.Println(err)
	}
	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
