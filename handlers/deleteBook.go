package handlers

import (
	"context"
	"github.com/ivansukach/book-service/protocol"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *BookAccounting) Delete(c echo.Context) error {
	log.Info("DeleteBook")
	book := new(BookModel)
	if err := c.Bind(book); err != nil {
		log.Errorf("echo.Context Error Delete %s", err)
		return err
	}
	_, err := a.client.Delete(context.Background(), &protocol.DeleteRequest{Id: book.Id})
	if err != nil {
		log.Errorf("GRPC Error DeleteClaims %s", err)
		return echo.ErrBadRequest
	}

	return c.String(http.StatusOK, "success")
}
