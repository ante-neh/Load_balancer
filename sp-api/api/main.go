package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/ante-neh/amazon-sp-api/internals/spapi"
)

func main() {
	lwaClient := &spapi.LWAClient{
		ClientID:     os.Getenv("LWA_CLIENT_ID"),
		ClientSecret: os.Getenv("LWA_CLIENT_SECRET"),
		RefreshToken: os.Getenv("LWA_REFRESH_TOKEN"),
	}

	spapiClient := spapi.NewSPAPIClient("eu-west-1", lwaClient)

	http.HandleFunc("/api/orders", func(w http.ResponseWriter, r *http.Request) {
		orders, err := spapiClient.GetOrders(time.Now().Add(-24 * time.Hour * 7)) 
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(orders)
	})

	fmt.Println("Go server running on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}