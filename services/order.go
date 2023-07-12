package services

import "TavernDDD/domain/customer"

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.Repository
}
