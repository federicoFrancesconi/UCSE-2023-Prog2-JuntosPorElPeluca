package handlers

import (
	"UCSE-2023-Prog2-TPIntegrador/dto"
	"UCSE-2023-Prog2-TPIntegrador/model"
	"UCSE-2023-Prog2-TPIntegrador/services"
	"UCSE-2023-Prog2-TPIntegrador/utils"
	"log"
	"net/http"
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
		log.Printf("[handler:ProductoHandler][method:ObtenerProductos][error:%s][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Obtiene el tipo de producto por el que se desea filtrar
	tipoProductoStr := c.DefaultQuery("tipoProducto", "-1")
	tipoProducto, err := strconv.Atoi(tipoProductoStr)

	if err != nil {
		log.Printf("[handler:ProductoHandler][method:ObtenerProductos][error:%s][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Armamos el filtro
	filtroProducto := utils.FiltroProducto{
		FiltrarPorStockMinimo: filtrarPorStockMinimo,
		TipoProducto:          model.TipoProducto(tipoProducto),
	}

	productos, err := handler.productoService.ObtenerProductosFiltrados(filtroProducto)

	//Si hay un error, lo devolvemos
	if err != nil {
		log.Printf("[handler:ProductoHandler][method:ObtenerProductos][error:%s][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:ProductoHandler][method:ObtenerProductos][cantidad:%d][user:%s]", len(productos), user.Codigo)

	c.JSON(http.StatusOK, productos)
}

func (handler *ProductoHandler) CrearProducto(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	var producto dto.Producto

	//Parseamos el body del request y lo guardamos en el objeto producto
	if err := c.ShouldBindJSON(&producto); err != nil {
		log.Printf("[handler:ProductoHandler][method:CrearProducto][error:%s][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Creamos el producto en la base de datos
	if err := handler.productoService.CrearProducto(&producto); err != nil {
		log.Printf("[handler:ProductoHandler][method:CrearProducto][error:%s][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:ProductoHandler][method:CrearProducto][user:%s]", user.Codigo)

	c.JSON(http.StatusCreated, true)
}

// Handler para eliminar un producto
func (handler *ProductoHandler) EliminarProducto(c *gin.Context) {
	user := dto.NewUser(utils.GetUserInfoFromContext(c))

	//Recibimos el codigo del producto a eliminar
	codigo := c.Param("codigo")

	//Convertimos el codigo a int
	codigoInt, err := strconv.Atoi(codigo)

	if err != nil {
		log.Printf("[handler:ProductoHandler][method:EliminarProducto][error:%s][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Creamos el objeto producto
	producto := dto.Producto{CodigoProducto: codigoInt}

	//Eliminamos el producto de la base de datos
	if err := handler.productoService.EliminarProducto(&producto); err != nil {
		log.Printf("[handler:ProductoHandler][method:EliminarProducto][error:%s][user:%s]", err.Error(), user.Codigo)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Agregamos un log para indicar información relevante del resultado
	log.Printf("[handler:ProductoHandler][method:EliminarProducto][user:%s]", user.Codigo)

	c.JSON(http.StatusOK, true)
}
