package app

import (
	"github.com/gin-gonic/gin"
)

type Application interface {
	Run()
}

type applicationImpl struct {
	r *gin.Engine
}

func NewApplication() Application {
	application := applicationImpl{
		r: RegisterHandler(),
	}
	return &application
}

func (m *applicationImpl) Run() {
	m.r.Run()
}

// Load resource

// Register handlers

// Run
