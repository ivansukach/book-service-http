package handlers

import (
	"context"
	"github.com/ivansukach/book-service/protocol"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *BookAccounting) Update(c echo.Context) error {
	log.Info("Update")
	book := new(BookModel)
	if err := c.Bind(book); err != nil {
		log.Errorf("echo.Context Error Update %s", err)
		return err
	}
	b := new(protocol.Book)
	b.Id = book.Id
	b.Title = book.Title
	b.Author = book.Author
	b.Genre = book.Genre
	b.Edition = book.Edition
	b.NumberOfPages = book.NumberOfPages
	b.Year = book.Year
	b.Amount = book.Amount
	b.IsPopular = book.IsPopular
	b.InStock = book.InStock
	_, err := a.client.Update(context.Background(), &protocol.UpdateRequest{
		Book: b})
	if err != nil {
		log.Errorf("GRPC Error SignUp %s", err)
		return err
	}
	return c.JSON(http.StatusOK, "success")
}
