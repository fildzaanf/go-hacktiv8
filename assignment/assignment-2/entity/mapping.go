package entity

import "assignment-2/model"

func OrderCoreToOrderModel(orderCore Order) model.Order {
	orderModel := model.Order{
		ID:           orderCore.ID,
		CustomerName: orderCore.CustomerName,
		Items:        ListItemCoreToItemModel(orderCore.Items),
		OrderedAt:    orderCore.OrderedAt,
		CreatedAt:    orderCore.CreatedAt,
		UpdatedAt:    orderCore.UpdatedAt,
		DeletedAt:    orderCore.DeletedAt,
	}
	return orderModel
}

func ListOrderCoreToOrderModel(orderCore []Order) []model.Order {
	listOrderModel := []model.Order{}
	for _, order := range orderCore {
		orderModel := OrderCoreToOrderModel(order)
		listOrderModel = append(listOrderModel, orderModel)
	}
	return listOrderModel
}

func OrderModelToOrderCore(orderModel model.Order) Order {
	orderCore := Order{
		ID:           orderModel.ID,
		CustomerName: orderModel.CustomerName,
		Items:        ListItemModelToItemCore(orderModel.Items),
		OrderedAt:    orderModel.OrderedAt,
		CreatedAt:    orderModel.CreatedAt,
		UpdatedAt:    orderModel.UpdatedAt,
		DeletedAt:    orderModel.DeletedAt,
	}
	return orderCore
}

func ListOrderModelToOrderCore(orderModel []model.Order) []Order {
	listOrderCore := []Order{}
	for _, order := range orderModel {
		orderCore := OrderModelToOrderCore(order)
		listOrderCore = append(listOrderCore, orderCore)
	}
	return listOrderCore
}

func ItemCoreToItemModel(itemCore Item) model.Item {
	itemModel := model.Item{
		ID:          itemCore.ID,
		ItemCode:    itemCore.ItemCode,
		Description: itemCore.Description,
		Quantity:    itemCore.Quantity,
		CreatedAt:   itemCore.CreatedAt,
		UpdatedAt:   itemCore.UpdatedAt,
		DeletedAt:   itemCore.DeletedAt,
	}
	return itemModel
}

func ListItemCoreToItemModel(itemCore []Item) []model.Item {
	listItemModel := []model.Item{}
	for _, item := range itemCore {
		itemModel := ItemCoreToItemModel(item)
		listItemModel = append(listItemModel, itemModel)
	}
	return listItemModel
}

func ItemModelToItemCore(itemModel model.Item) Item {
	itemCore := Item{
		ID:          itemModel.ID,
		ItemCode:    itemModel.ItemCode,
		Description: itemModel.Description,
		Quantity:    itemModel.Quantity,
		CreatedAt:   itemModel.CreatedAt,
		UpdatedAt:   itemModel.UpdatedAt,
		DeletedAt:   itemModel.DeletedAt,
	}
	return itemCore
}

func ListItemModelToItemCore(itemModel []model.Item) []Item {
	listItemCore := []Item{}
	for _, item := range itemModel {
		itemCore := ItemModelToItemCore(item)
		listItemCore = append(listItemCore, itemCore)
	}
	return listItemCore
}
