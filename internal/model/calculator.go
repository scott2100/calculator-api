package model

type Numbers struct {
	Number1 int `json:"a"`
	Number2 int `json:"b"`
}

type SumNumbers struct {
	Values []int `json:"values"`
}

type ResultResponse struct {
	Result int `json:"result"` // Or float64, etc.
}

type OperationFunc func(numbers Numbers) (error, int)
