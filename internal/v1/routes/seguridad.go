package routesv1

import (
	"github.com/buaazp/fasthttprouter"
	seguridadservicesv1 "github.com/dechevarrieta1/hrhelpers/internal/v1/services/v1/seguridad"
)

func Seguridad(router *fasthttprouter.Router) {
	router.GET("/alumnos", seguridadservicesv1.GetAlumnos)
	router.POST("/alumnos", seguridadservicesv1.CreateAlumno)
	// router.POST("/credenciales", SeguridadHandler)
}
