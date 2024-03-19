package entity

import "github.com/gin-gonic/gin"

type OrderRepositoryInterface interface {
	CreateOrder(orderCore Order) (Order, error)
	GetAllOrders(userID string) ([]Order, error)
	GetOrderByID(orderID string) (Order, error)
	UpdateOrderByID(orderID string, orderCore Order) (Order, error)
	DeleteOrderByID(orderID string) error
	CreateItem(itemCore Item) (Item, error)
	GetAllItems() ([]Item, error)
	GetItemByID(itemID string) (Item, error)
	UpdateItemByID(itemID string, itemCore Item) (Item, error)
	DeleteItemByID(itemID string) error
}

type OrderServiceInterface interface {
	CreateOrder(orderCore Order) (Order, error)
	GetAllOrders(userID string) ([]Order, error)
	GetOrderByID(orderID string) (Order, error)
	UpdateOrderByID(orderID string, orderCore Order) (Order, error)
	DeleteOrderByID(orderID string) error
	CreateItem(itemCore Item) (Item, error)
	GetAllItems() ([]Item, error)
	GetItemByID(itemID string) (Item, error)
	UpdateItemByID(itemID string, itemCore Item) (Item, error)
	DeleteItemByID(itemID string) error
}

type OrderHandlerInterface interface {
	CreateOrder(c gin.Context)
	GetAllOrders(c gin.Context)
	GetOrderByID(c gin.Context)
	UpdateOrderByID(c gin.Context)
	DeleteOrderByID(c gin.Context)
}
