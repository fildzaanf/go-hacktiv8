package request

import "assignment-2/modules/order/entity"

func OrderRequestToOrderCore(orderRequest OrderRequest) entity.Order {
	return entity.Order{
		OrderedAt:    orderRequest.OrderedAt,
		CustomerName: orderRequest.CustomerName,
		Items:        ListItemRequestToItemCore(orderRequest.Items),
	}
}

func UpdateOrderRequestToOrderCore(orderRequest UpdateOrderRequest) entity.Order {
	return entity.Order{
		OrderedAt:    orderRequest.OrderedAt,
		CustomerName: orderRequest.CustomerName,
		Items:        ListUpdateItemRequestToItemCore(orderRequest.Items),
	}
}

func ItemRequestToItemCore(itemRequest ItemRequest) entity.Item {
	return entity.Item{
		ItemCode:    itemRequest.ItemCode,
		Description: itemRequest.Description,
		Quantity:    itemRequest.Quantity,
	}
}

func UpdateItemRequestToItemCore(itemRequest UpdateItemRequest) entity.Item {
	return entity.Item{
		ID:          itemRequest.ID,
		ItemCode:    itemRequest.ItemCode,
		Description: itemRequest.Description,
		Quantity:    itemRequest.Quantity,
	}
}

func ListItemRequestToItemCore(itemRequest []ItemRequest) []entity.Item {
	listItemCore := []entity.Item{}
	for _, item := range itemRequest {
		itemCore := ItemRequestToItemCore(item)
		listItemCore = append(listItemCore, itemCore)
	}
	return listItemCore
}

func ListUpdateItemRequestToItemCore(itemRequest []UpdateItemRequest) []entity.Item {
	listItemCore := []entity.Item{}
	for _, item := range itemRequest {
		itemCore := UpdateItemRequestToItemCore(item)
		listItemCore = append(listItemCore, itemCore)
	}
	return listItemCore
}
