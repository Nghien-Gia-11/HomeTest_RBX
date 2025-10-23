package service

import (
	"HomeTestRBX/internal/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// helper
func makeLog(customerID, pageID int, day string) domain.Log {
	t, _ := time.Parse("2006-01-02", day)
	return domain.Log{
		Timestamp:  t,
		PageId:     pageID,
		CustomerId: customerID,
	}
}

func TestGetLoyalCustomers_SinglePageDifferentDays(t *testing.T) {
	analyzer := NewCustomerServiceImpl()

	day1 := []domain.Log{
		makeLog(1, 100, "2025-10-21"),
		makeLog(2, 200, "2025-10-21"),
	}

	day2 := []domain.Log{
		makeLog(1, 101, "2025-10-22"), // customer 1 truy cập page khác => loyal
		makeLog(2, 200, "2025-10-22"), // customer 2 vẫn truy cập cùng page => không loyal
	}

	customerLoyals, _ := analyzer.GetCustomerLoyal(day1, day2)

	assert.NotEmpty(t, customerLoyals)
}

func TestGetLoyalCustomers_MultiplePagesInDay2(t *testing.T) {
	analyzer := NewCustomerServiceImpl()

	day1 := []domain.Log{
		makeLog(3, 300, "2025-10-21"),
	}

	day2 := []domain.Log{
		makeLog(3, 301, "2025-10-22"),
		makeLog(3, 302, "2025-10-22"), // customer 3 có nhiều page trong day2 => loyal
	}

	customerLoyals, _ := analyzer.GetCustomerLoyal(day1, day2)

	assert.NotEmpty(t, customerLoyals)
}

func TestGetLoyalCustomers_NoOverlap(t *testing.T) {
	analyzer := NewCustomerServiceImpl()

	day1 := []domain.Log{
		makeLog(1, 100, "2025-10-21"),
	}

	day2 := []domain.Log{
		makeLog(2, 200, "2025-10-22"), // khác customer => không ai loyal
	}

	customerLoyals, _ := analyzer.GetCustomerLoyal(day1, day2)

	assert.Empty(t, customerLoyals)
}

func TestGetLoyalCustomers_EmptyLogs(t *testing.T) {
	analyzer := NewCustomerServiceImpl()

	var day1, day2 []domain.Log

	customerLoyals, _ := analyzer.GetCustomerLoyal(day1, day2)

	assert.Empty(t, customerLoyals)
}

func TestGetLoyalCustomers_SamePageButMultipleVisits(t *testing.T) {
	analyzer := NewCustomerServiceImpl()

	day1 := []domain.Log{
		makeLog(1, 101, "2025-10-21"),
		makeLog(1, 101, "2025-10-21"), // trùng page
	}

	day2 := []domain.Log{
		makeLog(1, 101, "2025-10-22"),
	}

	customerLoyals, _ := analyzer.GetCustomerLoyal(day1, day2)
	assert.Empty(t, customerLoyals) // không trung thành vì chỉ xem 1 page
}
