package utils

import (
	"log"
	"lokajatim/entities"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

var MidtransClient snap.Client

func InitMidtrans() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	if serverKey == "" {
		log.Fatal("MIDTRANS_SERVER_KEY not set in environment")
	}

	isProduction := false
	env := midtrans.Sandbox
	if isProduction {
		env = midtrans.Production
	}

	MidtransClient.New(serverKey, env)
	log.Println("Midtrans initialized with server key from environment")
}

func CreateTransaction(orderID string, grossAmount int64, Name, Email, PhoneNumber, Address string, cartItems []entities.CartItem) (string, error) {
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: grossAmount,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: Name,
			Email:     Email,
			Phone:     PhoneNumber,
			BillAddr: &midtrans.CustomerAddress{
				Address: Address,
			},
		},
		Items: convertCartItemsToSnapItems(cartItems),
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
	}
	resp, err := MidtransClient.CreateTransaction(req)
	if err != nil {
		return "", err
	}
	return resp.RedirectURL, nil
}

func convertCartItemsToSnapItems(cartItems []entities.CartItem) *[]midtrans.ItemDetails {
	items := make([]midtrans.ItemDetails, len(cartItems))
	for i, item := range cartItems {
		items[i] = midtrans.ItemDetails{
			ID:       strconv.Itoa(item.ProductID),
			Name:     item.Product.Name,
			Price:    int64(item.Product.Price),
			Qty:      int32(item.Quantity),
		}
	}
	return &items
}


func GenerateUniqueID() string {
	rand.Seed(time.Now().UnixNano())
	return time.Now().Format("20060102150405") + "-" + randomString(6)
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func GetCurrentTime() time.Time {
	return time.Now().UTC()
}
