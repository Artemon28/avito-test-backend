package internal

import (
	services "avito-test-backend/internal/services"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Handler struct {
	service *services.Service
}

func NewHandler(sr *services.Service) *Handler {
	return &Handler{service: sr}
}

// Deposit godoc
// @Summary      deposit
// @Description  top up the user's balance by id
// @Tags         User account
// @Accept       json
// @Produce      json
// @Param        input body transferRequest true  "User from whom this money, User to whom this money, order id, services id, money amount"
// @Success      200  {object}  structures.User
// @Failure      400
// @Router       /deposit [put]
func (h *Handler) Deposit(c *gin.Context) {
	var transfer transferRequest
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "incorrect Body: "+err.Error())
		return
	}
	err = json.Unmarshal(jsonData, &transfer)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "incorrect Body: "+err.Error())
		return
	}
	u, err := h.service.Deposit(transfer.Fromuserid, transfer.Touserid, transfer.Orderid, transfer.Serviceid, transfer.Amount)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "bd bad answer: "+err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, u)
}

// Book godoc
// @Summary      book
// @Description  book money on user account before withdraw operation
// @Tags         User account
// @Accept       json
// @Produce      json
// @Param        input body bookRequest true  "User id, money to book amount"
// @Success      200  {object}  structures.User
// @Failure      400
// @Router       /book [put]
func (h *Handler) Book(c *gin.Context) {
	var book bookRequest
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "incorrect Body: "+err.Error())
		return
	}
	err = json.Unmarshal(jsonData, &book)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "incorrect Body: "+err.Error())
		return
	}
	u, err := h.service.Book(book.Id, book.Amount)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "bd bad answer: "+err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, u)
}

// UnBook godoc
// @Summary      unbook
// @Description  unbook money on user account if withdraw fall
// @Tags         User account
// @Accept       json
// @Produce      json
// @Param        input body bookRequest true  "User id, money to unbook amount"
// @Success      200  {object}  structures.User
// @Failure      400
// @Router       /unbook [put]
func (h *Handler) UnBook(c *gin.Context) {
	var book bookRequest
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "incorrect Body: "+err.Error())
		return
	}
	err = json.Unmarshal(jsonData, &book)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "incorrect Body: "+err.Error())
		return
	}
	u, err := h.service.UnBook(book.Id, book.Amount)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "bd bad answer: "+err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, u)
}

// Withdraw godoc
// @Summary      withdraw
// @Description	 debiting money from a separate account
// @Tags         User account
// @Accept       json
// @Produce      json
// @Param        input body transferRequest true  "User from whom this money, User to whom this money, order id, services id, money amount"
// @Success      200  {object}  structures.User
// @Failure      400
// @Router       /withdraw [put]
func (h *Handler) Withdraw(c *gin.Context) {
	var transfer transferRequest
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "incorrect Body: "+err.Error())
		return
	}
	err = json.Unmarshal(jsonData, &transfer)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "incorrect Body: "+err.Error())
		return
	}
	u, err := h.service.Withdraw(transfer.Fromuserid, transfer.Touserid, transfer.Orderid, transfer.Serviceid, transfer.Amount)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "bd bad answer: "+err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, u)
}

// Balance godoc
// @Summary      balance
// @Description  get active balance of the user account
// @Tags         User account
// @Accept       json
// @Produce      json
// @Param        id   path      int true  "User id"
// @Success      200  {object}  structures.User
// @Failure      400
// @Router       /balance/{id} [get]
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

// Report godoc
// @Summary      report
// @Description	 create report CSV file with indicating the amount of revenue for each services, return url for the file
// @Tags         Accounting
// @Accept       json
// @Produce      json
// @Param        input body reportRequest true  "month MM, year YYYY"
// @Success      200  {object}  string
// @Failure      400
// @Router       /report [put]
func (h *Handler) Report(c *gin.Context) {
	var report reportRequest
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "incorrect Body: "+err.Error())
		return
	}
	err = json.Unmarshal(jsonData, &report)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "incorrect Body: "+err.Error())
		return
	}

	url, err := h.service.Report(report.Month, report.Year)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "bd bad answer: "+err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, url)
}

// Transactions godoc
// @Summary      transactions
// @Description	 get list of the transactions for user. User can set order by date or amount
// @Tags         User account
// @Accept       json
// @Produce      json
// @Param        input body HistoryRequest true  "user id, preferred order (date/amount/"")"
// @Success      200  {object}  structures.Order
// @Failure      400
// @Router       /transactions [get]
func (h *Handler) Transactions(c *gin.Context) {
	var history HistoryRequest
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "incorrect Body: "+err.Error())
		return
	}
	err = json.Unmarshal(jsonData, &history)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "incorrect Body: "+err.Error())
		return
	}

	orders, err := h.service.Transactions(history.Id, history.SortOrder)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "bd bad answer "+err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, orders)
}
