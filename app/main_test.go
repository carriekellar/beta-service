package main

import (
	"beta_service/api/handlers"
	"beta_service/api/router_groups"
	"beta_service/db/data_access"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Test_GetAllAssets(t *testing.T) {
	godotenv.Load("../.env")

	dbAccess, err := data_access.NewDbAccess()
	if err != nil {
		t.Fatal(err)
	}

	assetHandler, err := handlers.NewAssetHandler(dbAccess)
	if err != nil {
		t.Fatal(err)
	}

	// Create test router
	test_router := gin.New()
	// Add /asset endpoints to test router
	router_groups.NewAssetRouter(test_router.Group("/assets"), assetHandler)

	// Create a new server using the test router
	ts := httptest.NewServer(test_router)
	defer ts.Close()

	// Ping the "/assets/all" endpoint
	res, err := http.Get(ts.URL + "/assets/all?pageSize=4&pageIdx=4")
	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}
}

func Test_GetRedemptionAssets(t *testing.T) {
	godotenv.Load("../.env")

	dbAccess, err := data_access.NewDbAccess()
	if err != nil {
		t.Fatal(err)
	}

	redemptionHandler, err := handlers.NewRedemptionHandler(dbAccess)
	if err != nil {
		t.Fatal(err)
	}

	// Create test router
	test_router := gin.New()
	// Add /asset endpoints to test router
	router_groups.NewRedemptionRouter(test_router.Group("/redeem"), redemptionHandler)

	// Create a new server using the test router
	ts := httptest.NewServer(test_router)
	defer ts.Close()

	// Ping the "/assets/all" endpoint
	res, err := http.Get(ts.URL + "/redeem/assets?mint=8UEf7A4EeK5pZuPbEjC7RtX6DVJ4Coy5GzJcXUavqs8e")
	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}
}
