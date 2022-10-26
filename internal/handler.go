package internal

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Handler struct {
	service *service
}

func NewHandler(sr *service) *Handler {
	return &Handler{service: sr}
}

func (h *Handler) Deposit(c *gin.Context) {
	var transfer transferRequest
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "incorrect Body: "+err.Error())
	}
	err = json.Unmarshal(jsonData, &transfer)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "incorrect Body: "+err.Error())
	}
	u, err := h.service.Deposit(transfer.Fromuserid, transfer.Touserid, transfer.Orderid, transfer.Serviceid, transfer.Amount)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "bd bad answer: "+err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, u)
}

func (h *Handler) Book(c *gin.Context) {
	var book bookRequest
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "incorrect Body: "+err.Error())
	}
	err = json.Unmarshal(jsonData, &book)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "incorrect Body: "+err.Error())
	}
	u, err := h.service.Book(book.Id, book.Amount)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "bd bad answer: "+err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, u)
}

func (h *Handler) UnBook(c *gin.Context) {
	var book bookRequest
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "incorrect Body: "+err.Error())
	}
	err = json.Unmarshal(jsonData, &book)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "incorrect Body: "+err.Error())
	}
	u, err := h.service.UnBook(book.Id, book.Amount)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "bd bad answer: "+err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, u)
}

func (h *Handler) Withdraw(c *gin.Context) {
	var transfer transferRequest
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "incorrect Body: "+err.Error())
	}
	err = json.Unmarshal(jsonData, &transfer)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "incorrect Body: "+err.Error())
	}
	u, err := h.service.Withdraw(transfer.Fromuserid, transfer.Touserid, transfer.Orderid, transfer.Serviceid, transfer.Amount)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "bd bad answer: "+err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, u)
}

func (h *Handler) Balance(c *gin.Context) {
	userid, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "incorrect user id")
		return
	}
	balance, err := h.service.Balance(userid)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "bd bad answer: "+err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, balance)
}

func (h *Handler) Report(c *gin.Context) {
	var report reportRequest
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "incorrect Body: "+err.Error())
	}
	err = json.Unmarshal(jsonData, &report)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "incorrect Body: "+err.Error())
	}

	url, err := h.service.Report(report.Month, report.Year)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "bd bad answer: "+err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, url)
}

func (h *Handler) Transactions(c *gin.Context) {
	var history HistoryRequest
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "incorrect Body: "+err.Error())
	}
	err = json.Unmarshal(jsonData, &history)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "incorrect Body: "+err.Error())
	}

	orders, err := h.service.Transactions(history.Id, history.SortOrder)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "bd bad answer "+err.Error())
	}
	c.IndentedJSON(http.StatusOK, orders)
}
