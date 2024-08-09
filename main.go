package main

import (
	"log"
	"time"

	"github.com/AldiOktavianto/go-domain/module/proc/client"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	// PURCHASE REQUEST
	_purchaseRequestHandler "prisca-orchestrator/module/purchase_request/delivery/rest"
	_purchaseRequestRepository "prisca-orchestrator/module/purchase_request/repository"
	_purchaseRequestUseCase "prisca-orchestrator/module/purchase_request/usecase"
)

var (
	timeoutContext = time.Duration(2) * time.Second
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	procClient := client.NewProcClient(client.ProcClientParams{URL: "http://localhost:9090"})

	log.Println(procClient.PrClient)

	// Respositories
	purchaseRequestRepository := _purchaseRequestRepository.NewPurchaseRequestRepository(procClient.PrClient)

	// Usercases
	purchaseRequestUseCase := _purchaseRequestUseCase.NewPurchaseRequestUseCase(purchaseRequestRepository, timeoutContext)

	// Handlers
	_purchaseRequestHandler.NewPurchaseRequestHandler(e, purchaseRequestUseCase)

	e.Logger.Info(e.Start(":" + "8080"))
}

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	procClient "prisca-domain/module/proc/client"
// )

// func main() {
// 	client := procClient.NewProcClient(procClient.ProcClientParams{URL: "http://localhost:9090"})

// 	pr := procClient.PostPrRequest{
// 		Name: "Adam",
// 	}
// 	respPost, err := client.PrClient.PostPr(pr)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(StructToJson(respPost))

// 	respGet, err := client.PrClient.GetPr()

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(StructToJson(respGet))

// }

// func StructToJson(val interface{}) string {
// 	b, err := json.Marshal(val)
// 	if err != nil {
// 		return err.Error()
// 	}

// 	return string(b)
// }
