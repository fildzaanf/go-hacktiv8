package entity

import "github.com/gin-gonic/gin"

type RepositoryInterface interface {
	CreateOrder(orderCore Order) (Order, error)
	GetAllOrders() ([]Order, error)
	GetOrderByID(orderID string) (Order, error)
	UpdateOrderByID(orderID string, orderCore Order) error
	DeleteOrderByID(orderID string) error
	CreateItem(itemCore Item) (Item, error)
	GetAllItems() ([]Item, error)
	GetItemByID(itemID string) (Item, error)
	UpdateItemByID(itemID string, itemCore Item) error
	DeleteItemByID(itemID string) error
}

type ServiceInterface interface {
	CreateOrder(orderCore Order) (Order, error)
	GetAllOrders() ([]Order, error)
	GetOrderByID(orderID string) (Order, error)
	UpdateOrderByID(orderID string, orderCore Order) error
	DeleteOrderByID(orderID string) error
	CreateItem(itemCore Item) (Item, error)
	GetAllItems() ([]Item, error)
	GetItemByID(itemID string) (Item, error)
	UpdateItemByID(itemID string, itemCore Item) error
	DeleteItemByID(itemID string) error
}

type HandlerInterface interface {
	CreateOrder(c gin.Context) error
	GetAllOrders(c gin.Context) error
	GetOrderByID(c gin.Context) error
	UpdateOrderByID(c gin.Context) error
	DeleteOrderByID(c gin.Context) error
	CreateItem(c gin.Context) error
	GetAllItems(c gin.Context) error
	GetItemByID(c gin.Context) error
	UpdateItemByID(c gin.Context) error
	DeleteItemByID(c gin.Context) error
}
