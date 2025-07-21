package main

type service struct {
	store OrderStore
}

func NewService(store OrdersStore) *service {
	return &service{store}
}

func (s *service) CreateOrder(context.Context) {
	return nil
}