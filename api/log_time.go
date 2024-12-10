package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type LogTimeResponse struct {
	HydraMember []struct {
		TotalHours int `json:"totalHours"`
	} `json:"hydra:member"`
}

func GetLogTime(name string) (int, error) {
	now := time.Now()
	var startDate, endDate time.Time
	if now.Day() > 28 {
		startDate = time.Date(now.Year(), now.Month(), 29, 0, 0, 0, 0, now.Location())
		endDate = time.Date(now.Year(), now.Month()+1, 0, 0, 0, 0, 0, now.Location())
	} else {
		startDate = time.Date(now.Year(), now.Month()-1, 29, 0, 0, 0, 0, now.Location())
		endDate = time.Date(now.Year(), now.Month(), 28, 0, 0, 0, 0, now.Location())
	}
	startDateStr := startDate.Format("2006-01-02")
	endDateStr := endDate.Format("2006-01-02")
	requestBody := map[string]string{
		"startDate": startDateStr,
		"endDate":   endDateStr,
		"login":     name,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal request body: %v", err)
	}
	url := "https://logtime-med.1337.ma/api/get_log"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()
	var logTimeResp LogTimeResponse
	if err := json.NewDecoder(resp.Body).Decode(&logTimeResp); err != nil {
		return 0, fmt.Errorf("failed to decode response: %v", err)
	}
	if len(logTimeResp.HydraMember) == 0 {
		return 0, fmt.Errorf("no data found for user: %s", name)
	}
	return logTimeResp.HydraMember[0].TotalHours, nil
}
