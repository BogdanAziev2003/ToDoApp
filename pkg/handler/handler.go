package handler

import (
	"github.com/gin-gonic/gin"
)


type Handler struct {

}

func (h *Handler) InitRoutes() *gin.Engine {
	 router := gin.New()

	 auth := router.Group("/auth", h.signUp)
	 {
		 auth.POST("/sign-upp", h.signUp)
		 auth.POST("/sgn-in", h.signUp)
	 }

	 api := router.Group("/api", h.signUp)
	 {

		lists := api.Group("/lists", h.signUp)
		{
			lists.POST("/", h.signUp)
			lists.GET("/", h.signUp)
			lists.GET("/:id", h.signUp)
			lists.PUT("/:id", h.signUp)
			lists.DELETE("/:id", h.signUp)
   
			items := router.Group("/:id/items", h.signUp)
			{
			   items.POST("/", h.signUp)
			   items.GET("/", h.signUp)
			   items.GET("/:item_id", h.signUp)
			   items.PUT("/:item_id", h.signUp)
			   items.DELETE("/:item_id", h.signUp)
			}
		}
		 
	 }
	 return router
}