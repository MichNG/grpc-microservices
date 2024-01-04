package ports

import "github.com/MichNG/grpc-microservices/order/internal/application/domain"

type PaymentPort interface {
	Charge(order *domain.Order) error
}
