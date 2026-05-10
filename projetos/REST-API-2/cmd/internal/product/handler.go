package product

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase *Usecase
}

func NewHandler(usecase *Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) GetProducts(ctx *gin.Context) {
	//test do handler
	// products := []Product{
	// 	{
	// 		ID:    1,
	// 		Name:  "Batata Frita",
	// 		Price: 20,
	// 	},
	// }

	products, err := h.usecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (h *Handler) CreateProduct(ctx *gin.Context) {

	var product Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := h.usecase.repo.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (h Handler) GetProductsById(ctx *gin.Context) {

	id := ctx.Param("productId")

	if id == "" {
		response := Response{
			Message: "Id do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	productId, err := strconv.Atoi(id)
	if err != nil {
		response := Response{
			Message: "Id do produto näo pode ser um caracter",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := h.usecase.GetProductById(productId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := Response{
			Message: "Produto não encontrado",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)

}

func (h Handler) DeleteProductByID(ctx *gin.Context) {

	productID, err := strconv.Atoi(ctx.Param("productId"))

	if err != nil {
		response := Response{
			Message: "Id do produto não pode ser um caractere",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.usecase.DeleteProductByID(productID)

	if err != nil {
		response := Response{
			Message: err.Error(),
		}

		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := Response{
		Message: "Produto deletado com sucesso",
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateProduct(ctx *gin.Context) {

	id := ctx.Param("productId")

	if id == "" {
		response := Response{
			Message: "Id do produto não pode ser nulo",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productID, err := strconv.Atoi(id)

	if err != nil {
		response := Response{
			Message: "Id do produto inválido",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var product Product

	err = ctx.BindJSON(&product)

	if err != nil {
		response := Response{
			Message: "Erro ao converter json",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product.ID = productID

	err = h.usecase.UpdateProduct(product)

	if err != nil {
		response := Response{
			Message: "Erro ao atualizar produto",
		}

		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := Response{
		Message: "Produto atualizado com sucesso",
	}

	ctx.JSON(http.StatusOK, response)
}
