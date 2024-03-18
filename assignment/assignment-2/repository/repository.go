package repository

import (
	"assignment-2/entity"
	"assignment-2/model"
	"errors"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) entity.RepositoryInterface {
	return &repository{
		db: db,
	}
}

func (or *repository) CreateOrder(orderCore entity.Order) (entity.Order, error) {
	orderRequest := entity.OrderCoreToOrderModel(orderCore)

	tx := or.db.Create(&orderRequest)
	if tx.Error != nil {
		return entity.Order{}, tx.Error
	}

	orderResponse := entity.OrderModelToOrderCore(orderRequest)
	return orderResponse, nil
}

func (or *repository) GetAllOrders() ([]entity.Order, error) {
	orderModel := []model.Order{}

	tx := or.db.Model(&model.Order{}).Preload("Items").Find(&orderModel)
	if tx.Error != nil {
		return nil, tx.Error
	}

	orderResponse := entity.ListOrderModelToOrderCore(orderModel)
	return orderResponse, nil
}

func (or *repository) GetOrderByID(orderID string) (entity.Order, error) {
	orderModel := model.Order{}

	tx := or.db.Where("id = ?", orderID).Preload("Items").First(&orderModel)
	if tx.Error != nil {
		return entity.Order{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.Order{}, errors.New("data not found")
	}

	orderResponse := entity.OrderModelToOrderCore(orderModel)
	return orderResponse, nil
}

func (or *repository) UpdateOrderByID(orderID string, orderCore entity.Order) (entity.Order, error) {
	orderModel := model.Order{}

	tx := or.db.Where("id = ?", orderID).Preload("Items").First(&orderModel)
	if tx.Error != nil {
		return entity.Order{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.Order{}, errors.New("data not found")
	}

    updatedItems := []entity.Item{}
    for _, items := range orderCore.Items {
        updateItem, err := or.UpdateItemByID(items.ID, items)
        if err != nil {
            return entity.Order{}, err
        }
        updatedItems = append(updatedItems, updateItem)
    }

    updatedItemsCore := entity.ListItemCoreToItemModel(updatedItems)
	orderModel.Items = updatedItemsCore

	updatedOrderModel := entity.OrderCoreToOrderModel(orderCore)

	tx = or.db.Model(&orderModel).Updates(updatedOrderModel)
	if tx.Error != nil {
		return entity.Order{}, tx.Error
	}

	orderResponse := entity.OrderModelToOrderCore(orderModel)
	return orderResponse, nil
}



func (or *repository) DeleteOrderByID(orderID string) error {
	orderModel := model.Order{}

	tx := or.db.Where("id = ?", orderID).Delete(&orderModel)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("data not found")
	}
	return nil
}

func (ir *repository) CreateItem(itemCore entity.Item) (entity.Item, error) {
	itemRequest := entity.ItemCoreToItemModel(itemCore)

	tx := ir.db.Create(&itemRequest)
	if tx.Error != nil {
		return entity.Item{}, tx.Error
	}

	itemResponse := entity.ItemModelToItemCore(itemRequest)
	return itemResponse, nil
}

func (ir *repository) GetAllItems() ([]entity.Item, error) {
	itemModel := []model.Item{}

	tx := ir.db.Find(&itemModel)
	if tx.Error != nil {
		return nil, tx.Error
	}

	itemResponse := entity.ListItemModelToItemCore(itemModel)
	return itemResponse, nil
}

func (ir *repository) GetItemByID(itemID string) (entity.Item, error) {
	itemModel := model.Item{}

	tx := ir.db.Where("id = ?", itemID).First(&itemModel)
	if tx.Error != nil {
		return entity.Item{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.Item{}, errors.New("data not found")
	}

	itemResponse := entity.ItemModelToItemCore(itemModel)
	return itemResponse, nil
}

func (ir *repository) UpdateItemByID(itemID string, itemCore entity.Item) (entity.Item, error) {
	itemModel := model.Item{}

	tx := ir.db.Where("id = ?", itemID).First(&itemModel)
	if tx.Error != nil {
		return entity.Item{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.Item{}, errors.New("data not found")
	}

	updatedItemModel := entity.ItemCoreToItemModel(itemCore)

	tx = ir.db.Model(&itemModel).Updates(updatedItemModel)
	if tx.Error != nil {
		return entity.Item{}, tx.Error
	}

	itemResponse := entity.ItemModelToItemCore(itemModel)
	return itemResponse, nil
}

func (ir *repository) DeleteItemByID(itemID string) error {
	itemModel := model.Item{}

	tx := ir.db.Where("id = ?", itemID).Delete(&itemModel)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("data not found")
	}
	return nil
}
