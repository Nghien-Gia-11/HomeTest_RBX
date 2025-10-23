package service

import (
	"HomeTestRBX/internal/domain"
	"HomeTestRBX/internal/util"
)

type CustomerService interface {
	GetCustomerLoyal(logDay1, logDay2 []domain.Log) ([]int, error)
}

type CustomerServiceImpl struct{}

func NewCustomerServiceImpl() CustomerService {
	return &CustomerServiceImpl{}
}

func ExtractCustomerPage(logs []domain.Log) map[int]util.Set[int] {
	customerPage := make(map[int]util.Set[int]) // tạo 1 set có key là 1 customerID

	for _, log := range logs {
		if _, ok := customerPage[log.CustomerId]; !ok {
			customerPage[log.CustomerId] = util.NewSet[int]()
		}

		customerPage[log.CustomerId][log.PageId] = struct{}{}
	}
	return customerPage
}

func (c CustomerServiceImpl) GetCustomerLoyal(logDay1, logDay2 []domain.Log) ([]int, error) {
	customerDay1 := ExtractCustomerPage(logDay1)

	customerDay2 := ExtractCustomerPage(logDay2)

	var loyalCustomer []int

	for customer, pageDay1s := range customerDay1 {
		if pageDay2s, ok := customerDay2[customer]; ok {
			if pageDay1s.Len() == 1 && pageDay2s.Len() == 1 {
				pageDay1 := pageDay1s.Values()[0]
				pageDay2 := pageDay2s.Values()[0]
				if pageDay2 != pageDay1 {
					loyalCustomer = append(loyalCustomer, customer)
				}
			} else if pageDay2s.Len() > 1 {
				loyalCustomer = append(loyalCustomer, customer)
			}
		}
	}

	return loyalCustomer, nil
}
