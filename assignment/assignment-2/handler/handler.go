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
