package server

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

type InputJson struct {
	InputString string `json:"inputString"`
}

type OutputJson struct {
	OutputString string `json:"outputString"`
}

func GetVersion(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Error:", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Version", "1.0.0")
	fmt.Println(w.Header().Get("Version"))
}

func Decode(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Error:", http.StatusMethodNotAllowed)
		return
	}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}
	var input InputJson
	err = json.Unmarshal(data, &input)
	if err != nil {
		http.Error(w, "Error:", 500)
		return
	}
	decoded, _ := b64.StdEncoding.DecodeString(input.InputString)
	output := OutputJson{
		OutputString: string(decoded),
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)

	if err != nil {
		http.Error(w, "Error:", 500)
		return
	}
}

func HardOp(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	randomTime := rand.Intn(11) + 10
	time.Sleep(time.Duration(randomTime) * time.Second)

	badStatuses := []int{
		http.StatusBadGateway,
		http.StatusGatewayTimeout,
		http.StatusHTTPVersionNotSupported,
		http.StatusInsufficientStorage,
		http.StatusInternalServerError,
		http.StatusLoopDetected,
		http.StatusNetworkAuthenticationRequired,
	}
	randomTmp := rand.Intn(2)
	if randomTmp == 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		randomBadStatus := rand.Intn(len(badStatuses))
		w.WriteHeader(badStatuses[randomBadStatus])
	}
}
