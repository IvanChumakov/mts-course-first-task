package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"second-task/task2/internal/server"
	"time"
)

func RequestVersion(client *http.Client) (string, error) {
	request, err := http.NewRequest(http.MethodGet, "http://localhost:8001/version", nil)
	if err != nil {
		return "", errors.New("ошибка при создании запроса")
	}
	httpData, err := client.Do(request)
	if err != nil {
		return "", errors.New("ошибка при отправке запроса")
	}

	return httpData.Header.Get("Version"), nil
}

func DecodeRequest(client *http.Client) (string, error) {
	input := server.InputJson{
		InputString: "YWJjMTIzIT8kKiYoKSctPUB+",
	}
	jsonData, _ := json.Marshal(input)
	request, err := http.NewRequest(http.MethodPost, "http://localhost:8001/decode", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", errors.New("ошибка при создании запроса")
	}
	httpData, err := client.Do(request)
	if err != nil {
		return "", errors.New("ошибка при отправке запроса")
	}
	var output server.OutputJson
	data, _ := io.ReadAll(httpData.Body)
	_ = json.Unmarshal(data, &output)
	return output.OutputString, nil
}

func HardOpRequest(client *http.Client) (bool, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	requestWithTimeout, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8001/hard-op", nil)
	if err != nil {
		return false, 0, errors.New("ошибка при создании запроса")
	}
	httpData, err := client.Do(requestWithTimeout)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return false, -1, nil
		} else {
			return false, 0, err
		}
	}
	return true, httpData.StatusCode, nil
}
