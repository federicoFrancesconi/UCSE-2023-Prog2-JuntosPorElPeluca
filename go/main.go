package main

import (
	"log"

	"TPIntegrador/clients"
	"TPIntegrador/database"
	"TPIntegrador/handlers"
	"TPIntegrador/middlewares"
	"TPIntegrador/repositories"
	"TPIntegrador/services"

	"github.com/gin-gonic/gin"
)

var (
	camionHandler   *handlers.CamionHandler
	pedidoHandler   *handlers.PedidoHandler
	productoHandler *handlers.ProductoHandler
	envioHandler    *handlers.EnvioHandler

	router *gin.Engine
)

func main() {
	router = gin.Default()

	//Iniciar objetos de handler
	dependencies()
	//Iniciar rutas
	mappingRoutes()

	log.Println("Iniciando el servidor...")
	router.Run(":8080")
}

func mappingRoutes() {
	//implementa el metodo NewAuthMiddleware
	authClient := clients.NewAuthClient()
	authMiddleware := middlewares.NewAuthMiddleware(authClient)

	router.Use(middlewares.CORSMiddleware())
	router.Use(authMiddleware.ValidateToken)

	//Rutas de pedidos
	router.GET("/pedidos", pedidoHandler.ObtenerPedidos)
	router.GET("/pedidos/cantidadPorEstado", pedidoHandler.ObtenerCantidadPedidosPorEstado)
	router.POST("/pedidos", pedidoHandler.CrearPedido)
	router.PUT("/pedidos/:id/aceptar", pedidoHandler.AceptarPedido)
	router.PUT("/pedidos/:id/cancelar", pedidoHandler.CancelarPedido)

	//Rutas de envios
	router.GET("/envios", envioHandler.ObtenerEnvios)
	router.GET("/envios/:id", envioHandler.ObtenerEnvioPorId)
	router.GET("/envios/beneficioEntreFechas", envioHandler.ObtenerBeneficioEntreFechas)
	router.GET("/envios/cantidadPorEstado", envioHandler.ObtenerCantidadEnviosPorEstado)
	router.POST("/envios", envioHandler.CrearEnvio)
	router.POST("/envios/nuevaParada", envioHandler.AgregarParada)
	router.PUT("/envios/cambiarEstado", envioHandler.CambiarEstadoEnvio)

	//Rutas de camiones
	router.GET("/camiones", camionHandler.ObtenerCamiones)
	router.GET("/camiones/:patente", camionHandler.ObtenerCamionPorPatente)
	router.POST("/camiones", camionHandler.CrearCamion)
	router.PUT("/camiones", camionHandler.ActualizarCamion)
	router.DELETE("/camiones/:patente", camionHandler.EliminarCamion)

	//Rutas de productos
	router.GET("/productos", productoHandler.ObtenerProductos)
	router.GET("/productos/:codigo", productoHandler.ObtenerProductoPorCodigo)
	router.POST("/productos", productoHandler.CrearProducto)
	router.PUT("/productos", productoHandler.ActualizarProducto)
	router.DELETE("/productos/:codigo", productoHandler.EliminarProducto)
}

func dependencies() {
	database := database.NewMongoDB()

	//Iniciar repositorios
	camionRepository := repositories.NewCamionRepository(database)
	pedidoRepository := repositories.NewPedidoRepository(database)
	productoRepository := repositories.NewProductoRepository(database)
	envioRepository := repositories.NewEnvioRepository(database)

	//Iniciar servicios
	camionService := services.NewCamionService(camionRepository, envioRepository)
	pedidoService := services.NewPedidoService(pedidoRepository, envioRepository, productoRepository)
	productoService := services.NewProductoService(productoRepository, pedidoRepository)
	envioService := services.NewEnvioService(envioRepository, camionRepository, pedidoRepository, productoRepository)

	//Iniciar handlers
	camionHandler = handlers.NewCamionHandler(camionService)
	pedidoHandler = handlers.NewPedidoHandler(pedidoService)
	productoHandler = handlers.NewProductoHandler(productoService)
	envioHandler = handlers.NewEnvioHandler(envioService)
}
