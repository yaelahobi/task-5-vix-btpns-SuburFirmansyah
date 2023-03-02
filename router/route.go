package router

import (
	"task-5-vix-btpns-SuburFirmansyah/controllers"
	"task-5-vix-btpns-SuburFirmansyah/middlewares"
)

func InitRoutes(s *controllers.Server) {
	s.Router.POST("register", s.CreateUser)
	s.Router.GET("login", s.Login)

	s.Router.GET("photos", s.GetPhotos)
	needAuth := s.Router.Group("/")
	needAuth.Use(middlewares.AuthMiddleware(s.DB))
	{
		needAuth.PUT("users/:userId", s.UpdateUser)
		needAuth.DELETE("users/:userId", s.DeleteUser)

		needAuth.POST("photos", s.CreatePhoto)
		needAuth.PUT("photos/:photoId", s.UpdatePhoto)
		needAuth.DELETE("photos/:photoId", s.DeletePhoto)
	}
}
