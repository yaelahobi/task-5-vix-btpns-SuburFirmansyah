package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"task-5-vix-btpns-SuburFirmansyah/database"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func (s *Server) Initialize() {
	s.Router = gin.Default()
	s.DB = database.ConnDb()
}

func (s *Server) Run(port int) {
	s.Router.Run(fmt.Sprintf(":%d", port))
}
