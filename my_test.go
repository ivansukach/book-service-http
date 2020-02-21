package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ivansukach/book-service-http/config"
	"github.com/ivansukach/book-service-http/handlers"
	"github.com/ivansukach/book-service/protocol"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func TestBookService(t *testing.T) {
	log.Println("Client started")
	cfg := config.Load()
	opts := grpc.WithInsecure()
	clientConnInterface, err := grpc.Dial(cfg.AuthGRPCEndpoint, opts)
	if err != nil {
		log.Fatal(err)
	}
	defer clientConnInterface.Close()
	//timeout := time.Duration(500 * time.Second)
	////clientHTTP := http.Client{Timeout: timeout}
	client := protocol.NewBookServiceClient(clientConnInterface)
	bs := handlers.NewHandler(client)
	e := echo.New()
	e.POST("/create", bs.Create)
	e.POST("/update", bs.Update)
	e.POST("/delete", bs.Delete)
	e.POST("/getById", bs.GetById)
	e.POST("/listing", bs.Listing)
	go func() {
		_ = e.Start(fmt.Sprintf(":%d", cfg.Port))
	}()
	defer e.Close()

	//Create
	id := time.Now().Unix()
	idR := id % 10000
	bookId := fmt.Sprintf("bookId%d", id)
	title := fmt.Sprintf("title%d", idR)
	requestBody, err := json.Marshal(protocol.Book{
		Id:            bookId,
		Title:         title,
		Author:        "William Shakespeare",
		Genre:         "Dramatic",
		Edition:       "London №8",
		NumberOfPages: 221,
		Year:          1887,
		Amount:        12100,
		IsPopular:     true,
		InStock:       true,
	})
	if err != nil {
		log.Error(err)
	}
	log.Println(string(requestBody))
	responseCreate, err := http.Post("http://localhost:"+strconv.Itoa(cfg.Port)+"/create",
		"application/json",
		bytes.NewBuffer(requestBody))
	if err != nil {
		log.Error(err)
	}
	defer responseCreate.Body.Close()
	bodyCreate, err := ioutil.ReadAll(responseCreate.Body)
	if err != nil {
		log.Error(err)
	}
	log.Println(string(bodyCreate))

	//Update
	requestUpdateBody, err := json.Marshal(protocol.Book{
		Id:            bookId,
		Title:         title,
		Author:        "Ivan Sukach",
		Genre:         "Criminal",
		Edition:       "London №8",
		NumberOfPages: 221,
		Year:          2020,
		Amount:        1,
		IsPopular:     false,
		InStock:       false,
	})

	log.Println(string(requestUpdateBody))
	responseUpdate, err := http.Post("http://localhost:"+strconv.Itoa(cfg.Port)+"/update",
		"application/json",
		bytes.NewBuffer(requestUpdateBody))
	if err != nil {
		log.Error(err)
	}
	defer responseUpdate.Body.Close()
	bodyUpdate, err := ioutil.ReadAll(responseUpdate.Body)
	if err != nil {
		log.Error(err)
	}
	log.Println(string(bodyUpdate))

	//GetByID

	requestGetBody, err := json.Marshal(map[string]string{
		"id": bookId,
	})
	responseGet, err := http.Post("http://localhost:"+strconv.Itoa(cfg.Port)+"/getById",
		"application/json",
		bytes.NewBuffer(requestGetBody))
	if err != nil {
		log.Error(err)
	}
	defer responseGet.Body.Close()
	bodyGet, err := ioutil.ReadAll(responseGet.Body)
	if err != nil {
		log.Error(err)
	}
	log.Println(string(bodyGet))

	//DeleteBook

	requestDeleteBody, err := json.Marshal(map[string]string{
		"id": bookId,
	})
	responseDelete, err := http.Post("http://localhost:"+strconv.Itoa(cfg.Port)+"/delete",
		"application/json",
		bytes.NewBuffer(requestDeleteBody))
	if err != nil {
		log.Error(err)
	}
	defer responseDelete.Body.Close()
	bodyDelete, err := ioutil.ReadAll(responseDelete.Body)
	if err != nil {
		log.Error(err)
	}
	log.Println(string(bodyDelete))

	//Listing
	responseListing, err := http.Post("http://localhost:"+strconv.Itoa(cfg.Port)+"/listing",
		"application/json", nil)
	if err != nil {
		log.Error(err)
	}
	defer responseListing.Body.Close()
	bodyListing, err := ioutil.ReadAll(responseListing.Body)
	if err != nil {
		log.Error(err)
	}
	log.Println(string(bodyListing))
}
