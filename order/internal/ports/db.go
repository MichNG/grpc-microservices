package ports

import "github.com/MichNG/grpc-microservices/order/internal/application/domain"

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(order *domain.Order) error
}
