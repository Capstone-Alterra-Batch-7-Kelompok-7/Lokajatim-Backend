package transaction

import (
	"lokajatim/controllers/base"
	"lokajatim/controllers/pagination"
	"lokajatim/controllers/transaction/request"
	"lokajatim/controllers/transaction/response"
	"lokajatim/services/transaction"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	TransactionService transaction.TransactionService
}

func NewTransactionController(transactionService transaction.TransactionService) *TransactionController {
	return &TransactionController{TransactionService: transactionService}
}

func (controller *TransactionController) CreateTransaction(c echo.Context) error {
	req := new(request.TransactionRequest)
	if err := c.Bind(req); err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to bind request",
		})
	}

	transaction, err := req.ToEntities()
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Invalid request", 
		})
	}

	created, err := controller.TransactionService.CreateTransaction(transaction)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to create transaction",
		})
	}
	return base.SuccesResponse(c, response.TransactionFromEntity(created))
}

func (controller *TransactionController) GetTransactionByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transaction, err := controller.TransactionService.GetTransactionByID(id)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to get transaction",
		})
	}
	return base.SuccesResponse(c, response.TransactionFromEntity(transaction))
}

func (controller *TransactionController) GetAllTransactions(c echo.Context) error {
	transactions, err := controller.TransactionService.GetAllTransactions()
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to get transactions",
		})
	}
	return pagination.SuccessPaginatedResponse(c, transactions, 1, 10, int64(len(transactions)))
}

func (controller *TransactionController) UpdateTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	req := new(request.TransactionRequest)
	if err := c.Bind(req); err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to bind request",
		})
	}

	transaction, err := req.ToEntities()
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Invalid request",
		})
	}

	updated, err := controller.TransactionService.UpdateTransaction(id, transaction)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to update transaction",
		})
	}
	return base.SuccesResponse(c, response.TransactionFromEntity(updated))
}

func (controller *TransactionController) UpdateTransactionStatus(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	status := c.Param("status")
	err := controller.TransactionService.UpdateTransactionStatus(id, status)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to update transaction status",
		})
	}
	return base.SuccesResponse(c, response.TransactionResponse{
		TransactionID: strconv.Itoa(id),
		Status: status,
	})
}

func (controller *TransactionController) DeleteTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := controller.TransactionService.DeleteTransaction(id)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to delete transaction",
		})
	}
	return base.SuccesResponse(c, "Transaction deleted successfully")
}