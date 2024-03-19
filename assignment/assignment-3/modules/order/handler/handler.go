package handler

import (
	"assignment-2/middlewares"
	"assignment-2/modules/order/dto/request"
	"assignment-2/modules/order/dto/response"
	"assignment-2/modules/order/entity"
	"assignment-2/utils/constant"
	"assignment-2/utils/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	serviceInterface entity.OrderServiceInterface
}

func NewOrderHandler(serviceInterface entity.OrderServiceInterface) *orderHandler {
	return &orderHandler{
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
func (oh *orderHandler) CreateOrder(c *gin.Context) {

	userID, role, errExtract := middlewares.VerifyToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}
	if role != constant.USER {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	orderRequest := request.OrderRequest{}

	errBind := c.ShouldBind(&orderRequest)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
		return
	}

	request := request.OrderRequestToOrderCore(orderRequest)

	request.UserID = userID

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
func (oh *orderHandler) GetAllOrders(c *gin.Context) {

	userID, role, errExtract := middlewares.VerifyToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}
	if role != constant.USER {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	orders, errGetAll := oh.serviceInterface.GetAllOrders(userID)
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
func (oh *orderHandler) GetOrderByID(c *gin.Context) {
	orderID := c.Param("order_id")

	userID, role, errExtract := middlewares.VerifyToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}
	if role != constant.USER {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}
	
	order, errGet := oh.serviceInterface.GetOrderByID(orderID)
	if errGet != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse(errGet.Error()))
		return
	}

	if userID != order.UserID {
        c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
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
func (oh *orderHandler) UpdateOrderByID(c *gin.Context) {
	orderID := c.Param("order_id")

	userID, role, errExtract := middlewares.VerifyToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}
	if role != constant.USER {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	order, errGet := oh.serviceInterface.GetOrderByID(orderID)
    if errGet != nil {
        c.JSON(http.StatusInternalServerError, responses.ErrorResponse(errGet.Error()))
        return
    }

	if userID != order.UserID {
        c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
        return
    }


	orderRequest := request.UpdateOrderRequest{}

	errBind := c.ShouldBind(&orderRequest)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(errBind.Error()))
		return
	}

	request := request.UpdateOrderRequestToOrderCore(orderRequest)

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
func (oh *orderHandler) DeleteOrderByID(c *gin.Context) {
	orderID := c.Param("order_id")

	userID, role, errExtract := middlewares.VerifyToken(c)
	if errExtract != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse(errExtract.Error()))
		return
	}
	if role != constant.USER {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	order, errGet := oh.serviceInterface.GetOrderByID(orderID)
    if errGet != nil {
        c.JSON(http.StatusInternalServerError, responses.ErrorResponse(errGet.Error()))
        return
    }

	if userID != order.UserID {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse("not authorized to access this resource"))
		return
	}

	errDelete := oh.serviceInterface.DeleteOrderByID(orderID)
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse(errDelete.Error()))
		return
	}
	c.JSON(http.StatusOK, responses.SuccessResponse("data deleted successfully", nil))
}
