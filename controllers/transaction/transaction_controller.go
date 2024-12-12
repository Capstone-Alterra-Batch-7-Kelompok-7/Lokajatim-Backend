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
	TransactionService *transaction.TransactionService
}

func NewTransactionController(transactionService *transaction.TransactionService) *TransactionController {
	return &TransactionController{TransactionService: transactionService}
}

func (controller *TransactionController) CreateTransaction(c echo.Context) error {
	req := new(request.TransactionRequest)
	if err := c.Bind(req); err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to bind request",
		})
	}

	transaction, _ := req.ToEntities()

	created, err := controller.TransactionService.CreateTransaction(transaction.UserID, transaction.CartID)
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

	transactionResponses := make([]response.TransactionResponse, 0)
	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses, response.TransactionFromEntity(transaction))
	}

	return pagination.SuccessPaginatedResponse(c, transactionResponses, 1, 10, int64(len(transactionResponses)))
}

func (controller *TransactionController) UpdateTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	req := new(request.TransactionRequest)

	if err := c.Bind(req); err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to bind request",
		})
	}

	updatedTransaction, err := req.ToEntities()
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Invalid request data",
		})
	}

	transaction, err := controller.TransactionService.UpdateTransaction(id, updatedTransaction)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to update transaction",
		})
	}

	return base.SuccesResponse(c, response.TransactionFromEntity(transaction))
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
	return base.SuccesResponse(c, map[string]string{
		"transaction_id": strconv.Itoa(id),
		"status":         status,
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
