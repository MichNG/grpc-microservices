package ports

import (
	"github.com/MichNG/grpc-microservices/payment/internal/application/domain"
)

type DBPort interface {
	Get(id string) (domain.Payment, error)
	Save(payment *domain.Payment) error
}
