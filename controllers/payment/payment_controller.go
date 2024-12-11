package payment

import (
	"lokajatim/controllers/base"
	"lokajatim/controllers/payment/request"
	"lokajatim/controllers/payment/response"
	"lokajatim/services/payment"
	"lokajatim/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PaymentController struct {
	PaymentService payment.PaymentService
}

func NewPaymentController(paymentService payment.PaymentService) *PaymentController {
	return &PaymentController{PaymentService: paymentService}
}

func (controller *PaymentController) CreatePayment(c echo.Context) error {
	req := new(request.PaymentRequest)
	if err := c.Bind(req); err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to bind request",
		})
	}

	payment, err := req.ToEntities()
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Invalid request",
		})
	}

	created, err := controller.PaymentService.CreatePayment(payment)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to create payment",
		})
	}

	return base.SuccesResponse(c, response.PaymentFromEntity(created))
}

func (controller *PaymentController) GetPaymentByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	payment, err := controller.PaymentService.GetPaymentByID(id)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to get payment",
		})
	}

	return base.SuccesResponse(c, response.PaymentFromEntity(payment))
}

func (controller *PaymentController) UpdatePaymentStatus(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	status := c.Param("status")
	err := controller.PaymentService.UpdatePaymentStatus(id, status)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to update payment status",
		})
	}

	
	return base.SuccesResponse(c, response.PaymentResponse{
		ID:            id,
		PaymentStatus: status,
	})
}

func (controller *PaymentController) DeletePayment(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := controller.PaymentService.DeletePayment(id)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to delete payment",
		})
	}

	return base.SuccesResponse(c, "Payment deleted successfully")
}

func (controller *PaymentController) ProcessMidtransNotification(c echo.Context) error {
	var notification utils.NotificationRequest
	if err := c.Bind(&notification); err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to bind notification",
		})
	}

	err := controller.PaymentService.ProcessMidtransNotification(notification)
	if err != nil {
		return base.ErrorResponse(c, err, map[string]string{
			"error": "Failed to process payment notification",
		})
	}

	return base.SuccesResponse(c, "Payment notification processed successfully")
}
