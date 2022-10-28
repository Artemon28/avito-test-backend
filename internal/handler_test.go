package internal

import (
	"avito-test-backend/internal/services"
	mock_service "avito-test-backend/internal/services/mocks"
	"avito-test-backend/internal/structures"
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestHandler_Deposit(t *testing.T) {
	type mockBehavior func(s *mock_service.MockUserInterface, transfer transferRequest)

	tests := []struct {
		name                 string
		inputBody            []byte
		inputRequest         transferRequest
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: []byte(`{"fromuserid": 1, "touserid": 2, "serviceid": 1, "orderid": 1, "amount": 100}`),
			inputRequest: transferRequest{
				Fromuserid: 1,
				Touserid:   2,
				Serviceid:  1,
				Orderid:    1,
				Amount:     100,
			},
			mockBehavior: func(s *mock_service.MockUserInterface, transfer transferRequest) {
				s.EXPECT().Deposit(transfer.Fromuserid, transfer.Touserid, transfer.Serviceid,
					transfer.Orderid, transfer.Amount).Return(structures.User{Id: 2, Amount: 100, Bookamount: 0}, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: `{
    "id": 2,
    "amount": 100,
    "bookamount": 0
}`,
		},
		{
			name:      "Same users",
			inputBody: []byte(`{"fromuserid": 1, "touserid": 1, "serviceid": 2, "orderid": 2, "amount": 100}`),
			inputRequest: transferRequest{
				Fromuserid: 1,
				Touserid:   1,
				Serviceid:  2,
				Orderid:    2,
				Amount:     100,
			},
			mockBehavior: func(s *mock_service.MockUserInterface, transfer transferRequest) {
				s.EXPECT().Deposit(transfer.Fromuserid, transfer.Touserid, transfer.Serviceid,
					transfer.Orderid, transfer.Amount).Return(structures.User{}, errors.New("unable to deposit money from the same user"))
			},
			expectedStatusCode:   400,
			expectedResponseBody: `"bd bad answer: unable to deposit money from the same user"`,
		},
		{
			name:                 "empty Body",
			inputBody:            []byte(``),
			inputRequest:         transferRequest{},
			mockBehavior:         func(s *mock_service.MockUserInterface, transfer transferRequest) {},
			expectedStatusCode:   400,
			expectedResponseBody: `"incorrect Body: unexpected end of JSON input"`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			user := mock_service.NewMockUserInterface(c)
			test.mockBehavior(user, test.inputRequest)

			services := &services.Service{UserInterface: user}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.PUT("/deposit", handler.Deposit)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "/deposit", bytes.NewBuffer(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, string(w.Body.Bytes()), test.expectedResponseBody)
		})
	}
}

func TestHandler_Withdraw(t *testing.T) {
	type mockBehavior func(s *mock_service.MockUserInterface, transfer transferRequest)

	tests := []struct {
		name                 string
		inputBody            []byte
		inputRequest         transferRequest
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: []byte(`{"fromuserid": 1, "touserid": 2, "serviceid": 1, "orderid": 1, "amount": 100}`),
			inputRequest: transferRequest{
				Fromuserid: 1,
				Touserid:   2,
				Serviceid:  1,
				Orderid:    1,
				Amount:     100,
			},
			mockBehavior: func(s *mock_service.MockUserInterface, transfer transferRequest) {
				s.EXPECT().Withdraw(transfer.Fromuserid, transfer.Touserid, transfer.Serviceid,
					transfer.Orderid, transfer.Amount).Return(structures.User{Id: 1, Amount: 0, Bookamount: 0}, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: `{
    "id": 1,
    "amount": 0,
    "bookamount": 0
}`,
		},
		{
			name:      "Same users",
			inputBody: []byte(`{"fromuserid": 1, "touserid": 1, "serviceid": 2, "orderid": 2, "amount": 100}`),
			inputRequest: transferRequest{
				Fromuserid: 1,
				Touserid:   1,
				Serviceid:  2,
				Orderid:    2,
				Amount:     100,
			},
			mockBehavior: func(s *mock_service.MockUserInterface, transfer transferRequest) {
				s.EXPECT().Withdraw(transfer.Fromuserid, transfer.Touserid, transfer.Serviceid,
					transfer.Orderid, transfer.Amount).Return(structures.User{}, errors.New("unable to withdraw money from the same user"))
			},
			expectedStatusCode:   400,
			expectedResponseBody: `"bd bad answer: unable to withdraw money from the same user"`,
		},
		{
			name:                 "empty Body",
			inputBody:            []byte(``),
			inputRequest:         transferRequest{},
			mockBehavior:         func(s *mock_service.MockUserInterface, transfer transferRequest) {},
			expectedStatusCode:   400,
			expectedResponseBody: `"incorrect Body: unexpected end of JSON input"`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			user := mock_service.NewMockUserInterface(c)
			test.mockBehavior(user, test.inputRequest)

			services := &services.Service{UserInterface: user}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.PUT("/withdraw", handler.Withdraw)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "/withdraw", bytes.NewBuffer(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, string(w.Body.Bytes()), test.expectedResponseBody)
		})
	}
}

func TestHandler_Book(t *testing.T) {
	type mockBehavior func(s *mock_service.MockUserInterface, book bookRequest)

	tests := []struct {
		name                 string
		inputBody            []byte
		inputRequest         bookRequest
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: []byte(`{"id": 1, "amount": 100}`),
			inputRequest: bookRequest{
				Id:     1,
				Amount: 100,
			},
			mockBehavior: func(s *mock_service.MockUserInterface, book bookRequest) {
				s.EXPECT().Book(book.Id, book.Amount).Return(structures.User{Id: 1, Amount: 0, Bookamount: 100}, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: `{
    "id": 1,
    "amount": 0,
    "bookamount": 100
}`,
		},
		{
			name:                 "wrong body",
			inputBody:            []byte(`}{`),
			inputRequest:         bookRequest{},
			mockBehavior:         func(s *mock_service.MockUserInterface, book bookRequest) {},
			expectedStatusCode:   400,
			expectedResponseBody: `"incorrect Body: invalid character '}' looking for beginning of value"`,
		},
		{
			name:      "error from service",
			inputBody: []byte(`{"id": 0, "amount": 100}`),
			inputRequest: bookRequest{
				Id:     0,
				Amount: 100,
			},
			mockBehavior: func(s *mock_service.MockUserInterface, book bookRequest) {
				s.EXPECT().Book(book.Id, book.Amount).Return(structures.User{}, errors.New("insufficient funds"))
			},
			expectedStatusCode:   400,
			expectedResponseBody: `"bd bad answer: insufficient funds"`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			user := mock_service.NewMockUserInterface(c)
			test.mockBehavior(user, test.inputRequest)

			services := &services.Service{UserInterface: user}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.PUT("/book", handler.Book)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "/book", bytes.NewBuffer(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, string(w.Body.Bytes()), test.expectedResponseBody)
		})
	}
}

func TestHandler_UnBook(t *testing.T) {
	type mockBehavior func(s *mock_service.MockUserInterface, book bookRequest)

	tests := []struct {
		name                 string
		inputBody            []byte
		inputRequest         bookRequest
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: []byte(`{"id": 1, "amount": 100}`),
			inputRequest: bookRequest{
				Id:     1,
				Amount: 100,
			},
			mockBehavior: func(s *mock_service.MockUserInterface, book bookRequest) {
				s.EXPECT().UnBook(book.Id, book.Amount).Return(structures.User{Id: 1, Amount: 100, Bookamount: 0}, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: `{
    "id": 1,
    "amount": 100,
    "bookamount": 0
}`,
		},
		{
			name:                 "wrong body",
			inputBody:            []byte(`}{`),
			inputRequest:         bookRequest{},
			mockBehavior:         func(s *mock_service.MockUserInterface, book bookRequest) {},
			expectedStatusCode:   400,
			expectedResponseBody: `"incorrect Body: invalid character '}' looking for beginning of value"`,
		},
		{
			name:      "error from service",
			inputBody: []byte(`{"id": 0, "amount": 100}`),
			inputRequest: bookRequest{
				Id:     0,
				Amount: 100,
			},
			mockBehavior: func(s *mock_service.MockUserInterface, book bookRequest) {
				s.EXPECT().UnBook(book.Id, book.Amount).Return(structures.User{}, errors.New("insufficient funds"))
			},
			expectedStatusCode:   400,
			expectedResponseBody: `"bd bad answer: insufficient funds"`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			user := mock_service.NewMockUserInterface(c)
			test.mockBehavior(user, test.inputRequest)

			services := &services.Service{UserInterface: user}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.PUT("/unbook", handler.UnBook)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "/unbook", bytes.NewBuffer(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, string(w.Body.Bytes()), test.expectedResponseBody)
		})
	}
}

func TestHandler_Balance(t *testing.T) {
	type mockBehavior func(s *mock_service.MockUserInterface, id int)

	tests := []struct {
		name                 string
		mockBehavior         mockBehavior
		headerName           string
		headerValue          string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "Ok",
			mockBehavior:         func(s *mock_service.MockUserInterface, id int) {},
			headerName:           "id",
			headerValue:          "1",
			expectedStatusCode:   400,
			expectedResponseBody: `"incorrect user id"`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)

			defer c.Finish()

			user := mock_service.NewMockUserInterface(c)
			f, _ := strconv.Atoi(test.headerValue)
			test.mockBehavior(user, f)

			services := &services.Service{UserInterface: user}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			w := httptest.NewRecorder()

			r.GET("/balance", handler.Balance, func(c *gin.Context) {
				c.Params = []gin.Param{gin.Param{Key: "id", Value: "1"}}
			})

			// Init Test Request

			req := httptest.NewRequest("GET", "/balance", nil)
			req.Header.Set(test.headerName, test.headerValue)

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, string(w.Body.Bytes()), test.expectedResponseBody)
		})
	}
}
