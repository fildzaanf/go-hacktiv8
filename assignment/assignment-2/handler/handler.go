package handler

import (
	"assignment-2/dto/request"
	"assignment-2/dto/response"
	"assignment-2/entity"
	"assignment-2/utils/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	serviceInterface entity.ServiceInterface
}

func NewHandler(serviceInterface entity.ServiceInterface) *handler {
	return &handler{
		serviceInterface: serviceInterface,
	}
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order with the provided order data
// @Tags orders
// @Accept json
// @Produce json
// @Param orderRequest body request.OrderRequest true "Order data"
// @Success 201 {object} response.OrderResponse
// @Failure 400 {object} responses.TErrorResponse 
// @Router /orders [post]
func (oh *handler) CreateOrder(c *gin.Context) {
	orderRequest := request.OrderRequest{}

	errBind := c.ShouldBind(&orderRequest)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
		return
	}

	request := request.OrderRequestToOrderCore(orderRequest)

	order, errCreate := oh.serviceInterface.CreateOrder(request)
	if errCreate != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errCreate.Error()))
		return
	}

	response := response.OrderCoreToOrderResponse(order)

	c.JSON(http.StatusCreated, responses.SuccessResponse("data created successfully", response))
}

// GetAllOrders godoc
// @Summary Get all orders
// @Description Retrieve all orders
// @Tags orders
// @Produce json
// @Success 200 {object} responses.TSuccessResponse
// @Failure 500 {object} responses.TErrorResponse
// @Router /orders [get]
func (oh *handler) GetAllOrders(c *gin.Context) {
	orders, errGetAll := oh.serviceInterface.GetAllOrders()
	if errGetAll != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse(errGetAll.Error()))
		return
	}

	if len(orders) == 0 {
		c.JSON(http.StatusOK, responses.SuccessResponse("data is empty", nil))
		return
	}

	response := response.ListOrdeCoreToOrderResponse(orders)

	c.JSON(http.StatusOK, responses.SuccessResponse("data retrieved successfully", response))
}

// GetOrderByID godoc
// @Summary Get order by ID
// @Description Retrieve order details by order ID
// @Tags orders
// @Produce json
// @Param order_id path string true "Order ID"
// @Success 200 {object} response.OrderResponse
// @Failure 500 {object} responses.TErrorResponse
// @Router /orders/{order_id} [get]
func (oh *handler) GetOrderByID(c *gin.Context) {
	orderID := c.Param("order_id")

	order, errGet := oh.serviceInterface.GetOrderByID(orderID)
	if errGet != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse(errGet.Error()))
		return
	}

	response := response.OrderCoreToOrderResponse(order)

	c.JSON(http.StatusOK, responses.SuccessResponse("data retrieved successfully", response))
}

// UpdateOrderByID godoc
// @Summary Update order by ID
// @Description Update order details by order ID
// @Tags orders
// @Accept json
// @Produce json
// @Param order_id path string true "Order ID"
// @Param updateOrderRequest body request.UpdateOrderRequest true "Order details"
// @Success 200 {object} response.OrderResponse
// @Failure 400 {object} responses.TErrorResponse
// @Failure 500 {object} responses.TErrorResponse
// @Router /orders/{order_id} [put]
func (oh *handler) UpdateOrderByID(c *gin.Context) {
	orderRequest := request.UpdateOrderRequest{}

	errBind := c.ShouldBind(&orderRequest)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
		return
	}

	request := request.UpdateOrderRequestToOrderCore(orderRequest)

	orderID := c.Param("order_id")
	order, errUpdate := oh.serviceInterface.UpdateOrderByID(orderID, request) 
	if errUpdate != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errUpdate.Error()))
		return
	}

	response := response.OrderCoreToOrderResponse(order)

	c.JSON(http.StatusOK, responses.SuccessResponse("data updated successfully", response))
}

// DeleteOrderByID godoc
// @Summary Delete order by ID
// @Description Delete order by order ID
// @Tags orders
// @Produce json
// @Param order_id path string true "Order ID"
// @Success 200 {object} responses.TSuccessResponse
// @Failure 500 {object} responses.TErrorResponse
// @Router /orders/{order_id} [delete]
func (oh *handler) DeleteOrderByID(c *gin.Context) {
	orderID := c.Param("order_id")
	errDelete := oh.serviceInterface.DeleteOrderByID(orderID)
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse(errDelete.Error()))
		return
	}
	c.JSON(http.StatusOK, responses.SuccessResponse("data deleted successfully", nil))
}

func (ih *handler) CreateItem(c *gin.Context) {
	itemRequest := request.ItemRequest{}

	errBind := c.ShouldBind(&itemRequest)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
		return
	}

	request := request.ItemRequestToItemCore(itemRequest)

	item, errCreate := ih.serviceInterface.CreateItem(request)
	if errCreate != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errCreate.Error()))
		return
	}

	response := response.ItemCoreToItemResponse(item)

	c.JSON(http.StatusCreated, responses.SuccessResponse("data created successfully", response))
}

func (ih *handler) GetAllItems(c *gin.Context) {
	items, errGetAll := ih.serviceInterface.GetAllItems()
	if errGetAll != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse(errGetAll.Error()))
		return
	}

	if len(items) == 0 {
		c.JSON(http.StatusOK, responses.SuccessResponse("data is empty", nil))
		return
	}

	response := response.ListItemCoreToItemResponse(items)

	c.JSON(http.StatusOK, responses.SuccessResponse("data retrieved successfully", response))
}

func (ih *handler) GetItemByID(c *gin.Context) {
	itemID := c.Param("item_id")

	item, errGet := ih.serviceInterface.GetItemByID(itemID)
	if errGet != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse(errGet.Error()))
		return
	}

	response := response.ItemCoreToItemResponse(item)

	c.JSON(http.StatusOK, responses.SuccessResponse("data retrieved successfully", response))
}

func (ih *handler) UpdateItemByID(c *gin.Context) {
	itemRequest := request.ItemRequest{}

	errBind := c.ShouldBind(&itemRequest)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
		return
	}

	request := request.ItemRequestToItemCore(itemRequest)

	itemID := c.Param("item_id")
	item, errUpdate := ih.serviceInterface.UpdateItemByID(itemID, request)
	if errUpdate != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errUpdate.Error()))
		return
	}

	response := response.ItemCoreToItemResponse(item)

	c.JSON(http.StatusOK, responses.SuccessResponse("data updated successfully", response))
}

func (ih *handler) DeleteItemByID(c *gin.Context) {
	itemID := c.Param("item_id")
	errDelete := ih.serviceInterface.DeleteItemByID(itemID)
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse(errDelete.Error()))
		return
	}
	c.JSON(http.StatusOK, responses.SuccessResponse("data deleted successfully", nil))
}
