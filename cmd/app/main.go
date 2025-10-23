package main

import (
	"HomeTestRBX/internal/repository"
	"HomeTestRBX/internal/service"
	"fmt"
)

// home = 1, about = 2, services = 3, contact = 4, blog = 5
// 101 103 104 105 106 107 108 304

// customerID
// 101: {1,2} {1,2}
// 103: {4,5} {2,4}
// 104: {1,3} {1,5}
// 105: {2,4} {3,4}
// 106: {1,5} {1,2}
// 107: {3,4} {1,3}
// 108: {1,2} {5,4}
// 304: {1} {3,4}

func main() {

	logRepo := repository.NewLogRepositoryImpl()
	customerService := service.NewCustomerServiceImpl()

	log1, errLog1 := logRepo.ReadLogs("log_day_1.txt")
	log2, errLog2 := logRepo.ReadLogs("log_day_2.txt")

	if errLog1 != nil {
		panic(errLog1)
	}

	if errLog2 != nil {
		panic(errLog2)
	}

	loyalCustomer, err := customerService.GetCustomerLoyal(log1, log2)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Tổng số khách hàng trung thành là: %d\n", len(loyalCustomer))
	fmt.Println("Danh sách khách hàng trung thành:")
	for _, customer := range loyalCustomer {
		fmt.Println(customer)
	}
}
