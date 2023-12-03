package handlers

import (
	"TPIntegrador/dto"
	"TPIntegrador/services"
	"TPIntegrador/utils"
	"TPIntegrador/utils/logging"
	"TPIntegrador/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductoHandler struct {
	productoService services.ProductoServiceInterface
}

func NewProductoHandler(productoService services.ProductoServiceInterface) *ProductoHandler {
	return &ProductoHandler{productoService: productoService}
}

// Obtiene los productos, pudiendo filtrarlos por stock minimo y tipo de producto
func (handler *ProductoHandler) ObtenerProductos(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	//Pregunta si desea filtrar por stock minimo o no
	filtrarPorStockMinimoStr := c.DefaultQuery("filtrarPorStockMinimo", "false")
	filtrarPorStockMinimo, err := strconv.ParseBool(filtrarPorStockMinimoStr)

	if err != nil {
		logging.LoggearErrorYResponder(c, "ProductoHandler", "ObtenerProductos", err, &user)
		return
	}

	//Obtiene el tipo de producto por el que se desea filtrar
	tipoProducto := c.DefaultQuery("tipoProducto", "")

	//Armamos el filtro
	filtroProducto := utils.FiltroProducto{
		FiltrarPorStockMinimo: filtrarPorStockMinimo,
		TipoProducto:          model.TipoProducto(tipoProducto),
	}

	productos, err := handler.productoService.ObtenerProductos(filtroProducto)

	//Si hay un error, lo devolvemos
	if err != nil {
		logging.LoggearErrorYResponder(c, "ProductoHandler", "ObtenerProductos", err, &user)
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	logging.LoggearResultadoYResponder(c, "ProductoHandler", "ObtenerProductos", productos, &user)
}

func (handler *ProductoHandler) ObtenerProductoPorCodigo(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	//Recibimos el codigo del producto a buscar
	codigo := c.Param("codigo")

	//Creamos el objeto producto
	productoConCodigo := dto.Producto{CodigoProducto: codigo}

	producto, err := handler.productoService.ObtenerProductoPorCodigo(&productoConCodigo, &user)

	//Si hay un error, lo devolvemos
	if err != nil {
		logging.LoggearErrorYResponder(c, "ProductoHandler", "ObtenerProductoPorCodigo", err, &user)
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	logging.LoggearResultadoYResponder(c, "ProductoHandler", "ObtenerProductoPorCodigo", producto, &user)
}

func (handler *ProductoHandler) CrearProducto(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	var producto dto.Producto

	//Parseamos el body del request y lo guardamos en el objeto producto
	err := c.ShouldBindJSON(&producto)
	if err != nil {
		logging.LoggearErrorYResponder(c, "ProductoHandler", "CrearProducto", err, &user)
		return
	}

	//Creamos el producto en la base de datos
	err = handler.productoService.CrearProducto(&producto, &user)
	if err != nil {
		logging.LoggearErrorYResponder(c, "ProductoHandler", "CrearProducto", err, &user)
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	logging.LoggearResultadoYResponder(c, "ProductoHandler", "CrearProducto", true, &user)
}

// Handler para actualizar un producto
func (handler *ProductoHandler) ActualizarProducto(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	var producto dto.Producto

	//Parseamos el body del request y lo guardamos en el objeto producto
	err := c.ShouldBindJSON(&producto)
	if err != nil {
		logging.LoggearErrorYResponder(c, "ProductoHandler", "ActualizarProducto", err, &user)
		return
	}

	//Actualizamos el producto en la base de datos
	err = handler.productoService.ActualizarProducto(&producto, &user)
	if err != nil {
		logging.LoggearErrorYResponder(c, "ProductoHandler", "ActualizarProducto", err, &user)
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	logging.LoggearResultadoYResponder(c, "ProductoHandler", "ActualizarProducto", true, &user)
}

// Handler para eliminar un producto
func (handler *ProductoHandler) EliminarProducto(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	//Recibimos el codigo del producto a eliminar
	codigo := c.Param("codigo")

	//Creamos el objeto producto
	producto := dto.Producto{CodigoProducto: codigo}

	//Eliminamos el producto de la base de datos
	err := handler.productoService.EliminarProducto(&producto, &user)
	if err != nil {
		logging.LoggearErrorYResponder(c, "ProductoHandler", "EliminarProducto", err, &user)
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	logging.LoggearResultadoYResponder(c, "ProductoHandler", "EliminarProducto", true, &user)
}
