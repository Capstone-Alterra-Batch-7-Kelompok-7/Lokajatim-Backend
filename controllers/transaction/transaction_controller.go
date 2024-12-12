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

// @Summary Create Transaction
// @Description Create new transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param transaction body request.TransactionRequest true "Transaction data"
// @Success 200 {object} response.TransactionResponse
// @Router /transactions [post]
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

// @Summary Get Transaction by ID
// @Description Get transaction by ID
// @Tags Transaction
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} response.TransactionResponse
// @Failure 404 {object} map[string]string
// @Router /transactions/{id} [get]
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

// @Summary Get All Transactions
// @Description Get all transactions
// @Tags Transaction
// @Accept json
// @Produce json
// @Success 200 {object} []response.TransactionResponse
// @Router /transactions [get]
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

// @Summary Update Transaction
// @Description Update transaction by ID
// @Tags Transaction
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} response.TransactionResponse
// @Router /transactions/{id} [put]
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

// @Summary Update Transaction Status
// @Description Update transaction status by ID
// @Tags Transaction
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Param status path string true "Transaction status"
// @Success 200 {object} map[string]string
// @Router /transactions/{id}/status/{status} [put]
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

// @Summary Delete Transaction
// @Description Delete transaction by ID
// @Tags Transaction
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} map[string]string
// @Router /transactions/{id} [delete]
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
