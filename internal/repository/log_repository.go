package repository

import (
	"HomeTestRBX/internal/domain"
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

type LogRepository interface {
	ReadLogs(fileName string) ([]domain.Log, error)
}

type LogRepositoryImpl struct{}

func (l *LogRepositoryImpl) ReadLogs(fileName string) ([]domain.Log, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var logs []domain.Log
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

		logs = append(logs, domain.Log{
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

func NewLogRepositoryImpl() LogRepository {
	return &LogRepositoryImpl{}
}
