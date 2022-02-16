package fetch

import (
	"basic_golang/internal/domain/auth"
	"basic_golang/internal/domain/fetch/services"
)

type FetchDomainInterface interface {
	// ResourcesAdmin(ctx context.Context) (err error)
	// Resources(ctx context.Context) (err error)
}

func NewFetchDomain(authDomain auth.AuthDomainInterface) FetchDomainInterface {
	fetchServices := services.NewFetchServices(authDomain)
	return fetchServices
}
