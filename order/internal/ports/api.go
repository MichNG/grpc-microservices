package ports

import "github.com/MichNG/grpc-microservices/order/internal/application/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
