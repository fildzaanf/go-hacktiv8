package response

import "assignment-2/entity"

func OrderCoreToOrderResponse(orderCore entity.Order) OrderResponse {
	return OrderResponse{
		ID:           orderCore.ID,
		CustomerName: orderCore.CustomerName,
		OrderedAt:    orderCore.OrderedAt,
		Items:        ListItemCoreToItemResponse(orderCore.Items),
	}
}


func ListOrdeCoreToOrderResponse(orderCore []entity.Order) []OrderResponse {
	listOrderResponse := []OrderResponse{}
	for _, order := range orderCore {
		OrderResponse := OrderCoreToOrderResponse(order)
		listOrderResponse = append(listOrderResponse, OrderResponse)
	}
	return listOrderResponse
}


func ItemCoreToItemResponse(itemCore entity.Item) ItemResponse {
	return ItemResponse{
		ID:          itemCore.ID,
		ItemCode:    itemCore.ItemCode,
		Description: itemCore.Description,
		Quantity:    itemCore.Quantity,
	}
}

func ListItemCoreToItemResponse(itemCore []entity.Item) []ItemResponse {
	listItemResponse := []ItemResponse{}
	for _, item := range itemCore {
		itemResponse := ItemCoreToItemResponse(item)
		listItemResponse = append(listItemResponse, itemResponse)
	}
	return listItemResponse
}

