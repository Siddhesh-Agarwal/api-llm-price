package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ProviderPrice struct {
	InputCostInDollarsPerMillionTokens  float64 `json:"inputCostInDollarsPerMillionTokens"`
	OutputCostInDollarsPerMillionTokens float64 `json:"outputCostInDollarsPerMillionTokens"`
}

type ProviderDetails struct {
	Name  string        `json:"name"`
	Model string        `json:"model"`
	Price ProviderPrice `json:"price"`
}

var Providers = []ProviderDetails{
	{Name: "OpenAI", Model: "gpt-3.5-turbo", Price: ProviderPrice{InputCostInDollarsPerMillionTokens: 1.00, OutputCostInDollarsPerMillionTokens: 2.00}},
	{Name: "OpenAI", Model: "gpt-4o-mini", Price: ProviderPrice{InputCostInDollarsPerMillionTokens: 0.15, OutputCostInDollarsPerMillionTokens: 0.60}},
	{Name: "OpenAI", Model: "gpt-4o", Price: ProviderPrice{InputCostInDollarsPerMillionTokens: 2.50, OutputCostInDollarsPerMillionTokens: 10.00}},
	{Name: "OpenAI", Model: "gpt-4o-2024-05-13", Price: ProviderPrice{InputCostInDollarsPerMillionTokens: 5.00, OutputCostInDollarsPerMillionTokens: 20.00}},
	{Name: "OpenAI", Model: "gpt-4", Price: ProviderPrice{InputCostInDollarsPerMillionTokens: 30.00, OutputCostInDollarsPerMillionTokens: 60.00}},
	{Name: "OpenAI", Model: "gpt-4-32k", Price: ProviderPrice{InputCostInDollarsPerMillionTokens: 60.00, OutputCostInDollarsPerMillionTokens: 120.00}},
	{Name: "OpenAI", Model: "o1-mini", Price: ProviderPrice{InputCostInDollarsPerMillionTokens: 3.00, OutputCostInDollarsPerMillionTokens: 12.00}},
	{Name: "OpenAI", Model: "o1-preview", Price: ProviderPrice{InputCostInDollarsPerMillionTokens: 15.00, OutputCostInDollarsPerMillionTokens: 60.00}},
	{Name: "OpenAI", Model: "o1", Price: ProviderPrice{InputCostInDollarsPerMillionTokens: 15.00, OutputCostInDollarsPerMillionTokens: 60.00}},
	{Name: "Claude", Model: "3 Opus", Price: ProviderPrice{InputCostInDollarsPerMillionTokens: 15.00, OutputCostInDollarsPerMillionTokens: 75.00}},
	{Name: "Claude", Model: "3.5 Haiku", Price: ProviderPrice{InputCostInDollarsPerMillionTokens: 0.80, OutputCostInDollarsPerMillionTokens: 4.00}},
	{Name: "Claude", Model: "3.5 Sonnet", Price: ProviderPrice{InputCostInDollarsPerMillionTokens: 3.00, OutputCostInDollarsPerMillionTokens: 15.00}},
	{Name: "Google", Model: "1.0 Pro", Price: ProviderPrice{InputCostInDollarsPerMillionTokens: 0.50, OutputCostInDollarsPerMillionTokens: 1.50}},
	{Name: "Google", Model: "1.5 Flash-8b", Price: ProviderPrice{InputCostInDollarsPerMillionTokens: 0.0375, OutputCostInDollarsPerMillionTokens: 0.15}},
	{Name: "Google", Model: "1.5 Flash", Price: ProviderPrice{InputCostInDollarsPerMillionTokens: 0.075, OutputCostInDollarsPerMillionTokens: 0.30}},
	{Name: "Google", Model: "1.5 Pro", Price: ProviderPrice{InputCostInDollarsPerMillionTokens: 1.25, OutputCostInDollarsPerMillionTokens: 5.00}},
}

func getAllPrices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	json.NewEncoder(w).Encode(Providers)
}

func main() {
	port := 10000
	r := mux.NewRouter()
	r.HandleFunc("/", getAllPrices).Methods("GET")
	fmt.Printf("Server is running on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
