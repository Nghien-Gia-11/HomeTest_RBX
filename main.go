package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// home = 1, about = 2, services = 3, contact = 4, blog = 5
// 101 103 104 105 106 107 108 304
func main() {
	fileLog1 := "log_day_1.txt"
	fileLog2 := "log_day_2.txt"

	log1, errLog1 := readLogs(fileLog1)
	if errLog1 != nil {
		panic(errLog1)
	}
	customerDay1 := ExtractCustomerPage(log1)

	log2, errLog2 := readLogs(fileLog2)
	if errLog2 != nil {
		panic(errLog2)
	}
	customerDay2 := ExtractCustomerPage(log2)

	var loyalCustomer []int

	for customer, pageDay1 := range customerDay1 {
		if pageDay2, ok := customerDay2[customer]; ok {
			if len(pageDay2) == 1 && len(pageDay1) == 1 {
				if pageDay2[0] != pageDay1[0] {
					loyalCustomer = append(loyalCustomer, customer)
				}
			} else if len(pageDay2) > 1 {
				loyalCustomer = append(loyalCustomer, customer)
			}
		}
	}

	fmt.Printf("Tổng số khách hàng trung thành là: %d\n", len(loyalCustomer))
	fmt.Println("Danh sách khách hàng trung thành:")
	for _, customer := range loyalCustomer {
		fmt.Println(customer)
	}
}

func readLogs(fileName string) ([]Log, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var logs []Log
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue // bỏ qua những dòng nào thiếu or thừa các thành phần
		}

		timestampStr := parts[0]
		timestamp, errParseTime := time.Parse(time.RFC3339, timestampStr)
		pageId, errParsePage := strconv.Atoi(parts[1])
		customerId, errParseCustomer := strconv.Atoi(parts[2])

		if errParsePage != nil || errParseCustomer != nil || errParseTime != nil {
			continue
		}

		logs = append(logs, Log{
			Timestamp:  timestamp,
			PageId:     pageId,
			CustomerId: customerId,
		})

		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}
	return logs, nil
}

func ExtractCustomerPage(logs []Log) map[int]map[int]struct{} {
	customerPage := make(map[int]map[int]struct{}) // tạo 1 map có key là 1 customerID và value là 1 set các pageId (map[int]struct{} = set)

	for _, log := range logs {
		if _, ok := customerPage[log.CustomerId]; !ok {
			customerPage[log.CustomerId] = make(map[int]struct{})
		}

		customerPage[log.CustomerId][log.PageId] = struct{}{}
	}
	return customerPage
}
