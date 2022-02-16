package services

import (
	"basic_golang/internal/domain/auth"
)

type fetchDomain struct {
	authDomain auth.AuthDomainInterface
}

func NewFetchServices(authDomain auth.AuthDomainInterface) *fetchDomain {
	return &fetchDomain{
		authDomain,
	}
}
