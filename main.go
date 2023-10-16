package main

import (
	"log"

	"UCSE-2023-Prog2-TPIntegrador/database"
	"UCSE-2023-Prog2-TPIntegrador/handlers"
	"UCSE-2023-Prog2-TPIntegrador/repositories"
	"UCSE-2023-Prog2-TPIntegrador/services"

	"github.com/gin-gonic/gin"
	//"UCSE-2023-Prog2-TPIntegrador/middlewares"
	//"UCSE-2023-Prog2-TPIntegrador/clients"
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
	//-------------------- Middleware --------------------
	//middleware para permitir peticiones del mismo server localhost

	//cliente para api externa
	// var authClient clients.AuthClientInterface
	// authClient = clients.NewAuthClient()

	//creacion de middleware de autenticacion
	//authMiddleware := middlewares.NewAuthMiddleware(authClient)

	//Uso del middleware para todas las rutas del grupo
	//group.Use(authMiddleware.ValidateToken)

	//group.Use(middlewares.CORSMiddleware())

	//------------------------------------------------------

	//Listado de rutas
	pedidos := router.Group("/pedidos")
	envios := router.Group("/envios")
	camiones := router.Group("/camiones")
	productos := router.Group("/productos")

	//Rutas de pedidos
	pedidos.GET("", pedidoHandler.ObtenerPedidos)
	pedidos.POST("", pedidoHandler.CrearPedido)
	pedidos.PUT("/:id/aceptar", pedidoHandler.AceptarPedido)
	pedidos.PUT("/:id/cancelar", pedidoHandler.CancelarPedido)

	//Rutas de envios
	envios.GET("", envioHandler.ObtenerEnvios)
	envios.GET("/:id", envioHandler.ObtenerEnvioPorId)
	envios.POST("", envioHandler.CrearEnvio)
	//TODO: hay que modificar estos metodos en el repo, service y handler
	// envios.PUT("/:id/nuevaParada", envioHandler.AgregarParada)
	// envios.PUT("/:id/finalizar", envioHandler.FinalizarViaje)
	// envios.PUT("/:id/iniciar", envioHandler.IniciarViaje)

	//Rutas de camiones
	camiones.GET("", camionHandler.ObtenerCamiones)
	camiones.GET("/:patente", camionHandler.ObtenerCamionPorPatente)
	camiones.POST("", camionHandler.CrearCamion)
	camiones.PUT("", camionHandler.ActualizarCamion)
	camiones.DELETE("", camionHandler.EliminarCamion)

	//Rutas de productos
	productos.GET("", productoHandler.ObtenerProductos)
	productos.POST("", productoHandler.CrearProducto)
	productos.DELETE("", productoHandler.EliminarProducto)
}

func dependencies() {
	database := database.NewMongoDB()

	//Iniciar repositorios
	camionRepository := repositories.NewCamionRepository(database)
	pedidoRepository := repositories.NewPedidoRepository(database)
	productoRepository := repositories.NewProductoRepository(database)
	envioRepository := repositories.NewEnvioRepository(database)

	//Iniciar servicios
	camionService := services.NewCamionService(camionRepository)
	pedidoService := services.NewPedidoService(pedidoRepository)
	productoService := services.NewProductoService(productoRepository)
	conexionService := services.NewConexionService(camionService, pedidoService, productoService)
	envioService := services.NewEnvioService(envioRepository, conexionService)

	//Iniciar handlers
	camionHandler = handlers.NewCamionHandler(camionService)
	pedidoHandler = handlers.NewPedidoHandler(pedidoService)
	productoHandler = handlers.NewProductoHandler(productoService)
	envioHandler = handlers.NewEnvioHandler(envioService)
}
