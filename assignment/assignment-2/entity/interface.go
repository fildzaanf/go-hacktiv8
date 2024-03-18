package entity

import "github.com/gin-gonic/gin"

type RepositoryInterface interface {
	CreateOrder(orderCore Order) (Order, error)
	GetAllOrders() ([]Order, error)
	GetOrderByID(orderID string) (Order, error)
	UpdateOrderByID(orderID string, orderCore Order) (Order, error)
	DeleteOrderByID(orderID string) error
	CreateItem(itemCore Item) (Item, error)
	GetAllItems() ([]Item, error)
	GetItemByID(itemID string) (Item, error)
	UpdateItemByID(itemID string, itemCore Item) (Item, error)
	DeleteItemByID(itemID string) error
}

type ServiceInterface interface {
	CreateOrder(orderCore Order) (Order, error)
	GetAllOrders() ([]Order, error)
	GetOrderByID(orderID string) (Order, error)
	UpdateOrderByID(orderID string, orderCore Order) (Order, error)
	DeleteOrderByID(orderID string) error
	CreateItem(itemCore Item) (Item, error)
	GetAllItems() ([]Item, error)
	GetItemByID(itemID string) (Item, error)
	UpdateItemByID(itemID string, itemCore Item) (Item, error)
	DeleteItemByID(itemID string) error
}

type HandlerInterface interface {
	CreateOrder(c gin.Context)
	GetAllOrders(c gin.Context)
	GetOrderByID(c gin.Context)
	UpdateOrderByID(c gin.Context)
	DeleteOrderByID(c gin.Context)
	CreateItem(c gin.Context)
	GetAllItems(c gin.Context)
	GetItemByID(c gin.Context)
	UpdateItemByID(c gin.Context)
	DeleteItemByID(c gin.Context)
}
