package service

import (
	"assignment-2/modules/order/entity"
	"assignment-2/utils/validator"
	"errors"
)

type orderService struct {
	repositoryInterface entity.OrderRepositoryInterface
}

func NewOrderService(repositoryInterface entity.OrderRepositoryInterface) entity.OrderServiceInterface {
	return &orderService{
		repositoryInterface: repositoryInterface,
	}
}

func (os *orderService) CreateOrder(orderCore entity.Order) (entity.Order, error) {
	errDataEmpty := validator.IsDataEmpty(orderCore.CustomerName, orderCore.OrderedAt, orderCore.Items)
	if errDataEmpty != nil {
		return entity.Order{}, errDataEmpty
	}

	for _, items := range orderCore.Items {
		errDataEmpty = validator.IsDataEmpty(items.ItemCode, items.Description, items.Quantity)
		if errDataEmpty != nil {
			return entity.Order{}, errDataEmpty
		}
	}

	createOrder, errCreateOrderRepo := os.repositoryInterface.CreateOrder(orderCore)
	if errCreateOrderRepo != nil {
		return entity.Order{}, errCreateOrderRepo
	}

	return createOrder, nil
}

func (os *orderService) GetAllOrders() ([]entity.Order, error) {
	getAllOrders, errGetAllOrdersRepo := os.repositoryInterface.GetAllOrders()
	if errGetAllOrdersRepo != nil {
		return nil, errGetAllOrdersRepo
	}

	return getAllOrders, nil
}

func (os *orderService) GetOrderByID(orderID string) (entity.Order, error) {
	if orderID == "" {
		return entity.Order{}, errors.New("id not found")
	}

	getOrderByID, errGetOrderByID := os.repositoryInterface.GetOrderByID(orderID)
	if errGetOrderByID != nil {
		return entity.Order{}, errGetOrderByID
	}

	return getOrderByID, nil
}

func (os *orderService) UpdateOrderByID(orderID string, orderCore entity.Order) (entity.Order, error) {
	errDataEmpty := validator.IsDataEmpty(orderCore.CustomerName, orderCore.OrderedAt, orderCore.Items)
	if errDataEmpty != nil {
		return entity.Order{}, errDataEmpty
	}

	for _, items := range orderCore.Items {
		errDataEmpty = validator.IsDataEmpty(items.ItemCode, items.Description, items.Quantity)
		if errDataEmpty != nil {
			return entity.Order{}, errDataEmpty
		}
	}

	updateOrderByID, errUpdateOrderByIDRepo := os.repositoryInterface.UpdateOrderByID(orderID, orderCore)
	if errUpdateOrderByIDRepo != nil {
		return entity.Order{}, errUpdateOrderByIDRepo
	}

	return updateOrderByID, nil
}

func (os *orderService) DeleteOrderByID(orderID string) error {
	if orderID == "" {
		return errors.New("id not found")
	}

	errDeleteOrderByID := os.repositoryInterface.DeleteOrderByID(orderID)
	if errDeleteOrderByID != nil {
		return errDeleteOrderByID
	}

	return nil
}

func (is *orderService) CreateItem(itemCore entity.Item) (entity.Item, error) {
	errDataEmpty := validator.IsDataEmpty(itemCore.ItemCode, itemCore.Description, itemCore.Quantity)
	if errDataEmpty != nil {
		return entity.Item{}, errDataEmpty
	}

	createItem, errCreateItemRepo := is.repositoryInterface.CreateItem(itemCore)
	if errCreateItemRepo != nil {
		return entity.Item{}, errCreateItemRepo
	}

	return createItem, nil
}

func (is *orderService) GetAllItems() ([]entity.Item, error) {
	getAllItems, errGetAllItemsRepo := is.repositoryInterface.GetAllItems()
	if errGetAllItemsRepo != nil {
		return nil, errGetAllItemsRepo
	}

	return getAllItems, nil
}

func (is *orderService) GetItemByID(itemID string) (entity.Item, error) {
	if itemID == "" {
		return entity.Item{}, errors.New("id not found")
	}

	getItemByID, errGetItemByID := is.repositoryInterface.GetItemByID(itemID)
	if errGetItemByID != nil {
		return entity.Item{}, errGetItemByID
	}

	return getItemByID, nil
}

func (is *orderService) UpdateItemByID(itemID string, itemCore entity.Item) (entity.Item, error) {
	errDataEmpty := validator.IsDataEmpty(itemCore.ID, itemCore.ItemCode, itemCore.Description, itemCore.Quantity)
	if errDataEmpty != nil {
		return entity.Item{}, errDataEmpty
	}

	updateItemByID, errUpdateItemByIDRepo := is.repositoryInterface.UpdateItemByID(itemID, itemCore)
	if errUpdateItemByIDRepo != nil {
		return entity.Item{}, errUpdateItemByIDRepo
	}

	return updateItemByID, nil
}

func (is *orderService) DeleteItemByID(itemID string) error {
	if itemID == "" {
		return errors.New("id not found")
	}

	errDeleteItemByID := is.repositoryInterface.DeleteItemByID(itemID)
	if errDeleteItemByID != nil {
		return errDeleteItemByID
	}

	return nil
}
