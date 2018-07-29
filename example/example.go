package main

import (
	"time"
	"os"
	"log"
	"fmt"
	"errors"
	"io/ioutil"
	"github.com/yudhaputrama/nicepay"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
)

var (
	client *nicepay.Client
	merchantId string
	merchantKey string
)

func registerVirtualAccount(c *nicepay.Client) {
	now := time.Now()
	validVA := now.Local().Add(time.Hour * time.Duration(5))

	res, _, err := c.Registration(&nicepay.RegistrationRequest{
		TimeStamp: now.Format(nicepay.FORMAT_TIMESTAMP),
		MerchantId: c.MerchantId,
		Amount: 25000,
		Currency: "IDR",
		ReferenceNo: "1234",
		GoodsName: "Nicepay Pillow",
		Description: "Jalan Sukasuka",
		BillingName:"Nicepay",
		BillingPhone:"087845654877",
		BillingEmail:"hello@test.com",
		BillingAddress: "Bandung",
		BillingCity: "Bandung",
		BillingState: "Jawa Barat",
		BillingPostalCode: "55763",
		BillingCountry: "Indonesia",
		NotificationUrl: "http://requestbin.fullcontact.com/1fexw2h1",
		CartData: "{}",

		PayMethod: nicepay.MethodVirtualAccount,
		Bank: nicepay.BankBNI,
		VirtualAccountValidDate: validVA.Format(nicepay.FORMAT_DATE),
		VirtualAccountValidTime: validVA.Format(nicepay.FORMAT_TIME),
	})


	if err != nil {
		log.Fatalln(err)
	}

	if res.ResultCd != "0000" {
		log.Fatal(res)
	}

	fmt.Println(res.TransactionID)
}

func registerCreditCard(c *nicepay.Client) error {
	now := time.Now()
	request := &nicepay.RegistrationRequest{
		TimeStamp: now.Format(nicepay.FORMAT_TIMESTAMP),
		MerchantId: c.MerchantId,
		Amount: 20000,
		Currency: "IDR",
		ReferenceNo: "1234",
		GoodsName: "Nicepay Pillow",
		Description: "Jalan Sukasuka",
		BillingName:"Nicepay",
		BillingPhone:"087845654877",
		BillingEmail:"hello@test.com",
		BillingAddress: "Dayeuhkolot Bandung",
		BillingCity: "Bandung",
		BillingState: "Jawa Barat",
		BillingPostalCode: "55763",
		BillingCountry: "Indonesia",
		NotificationUrl: "http://ptsv2.com/t/s6v4a-1530243879/post",
		CartData: "{}",
		UserIP: "127.0.0.1",

		PayMethod: nicepay.MethodCreditCard,
		InstallmentType: nicepay.MerchantCharge,
		InstallmentMonth: 1,
		RecurringOption: nicepay.ReccAutomaticCancel,
	}

	c.GenerateMerchantToken(request)

	res, _, err := c.Registration(request)

	if err != nil {
		log.Fatalln(err)
		return err
	}

	if res.ResultCd != "0000" {
		log.Fatal(res)
		return errors.New(res.ResultMsg)
	}

	fmt.Println(res.TransactionID)

	payment := &nicepay.PaymentRequest{
		TimeStamp: now.Format(nicepay.FORMAT_TIMESTAMP),
		TransactionID: res.TransactionID,
		CardNo: "4222222222222222",
		CardExp: "2006",
		CardCvv: "123",
		CardRecurringToken: "",
		CardPreauthToken: "",
		MerchantToken: request.MerchantToken,
		CallBackURL: "http://ptsv2.com/t/s6v4a-1530243879/post",
	}

	resHttp, err := c.PaymentCreditCard(payment)

	if err != nil {
		log.Fatal(err)
		return err
	}

	bodyBytes, _ := ioutil.ReadAll(resHttp.Body)
	fmt.Println(string(bodyBytes))

	return nil
}

func Register(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	request := &nicepay.RegistrationRequest{
		TimeStamp: now.Format(nicepay.FORMAT_TIMESTAMP),
		MerchantId: client.MerchantId,
		Amount: 20000,
		Currency: "IDR",
		ReferenceNo: "1234",
		GoodsName: "Pillow Nicepay",
		Description: "Jalan Riau",
		BillingName:"Nicepay",
		BillingPhone:"08123456789",
		BillingEmail:"test@test.com",
		BillingAddress: "Sukasuka",
		BillingCity: "Bandung",
		BillingState: "Jawa Barat",
		BillingPostalCode: "55555",
		BillingCountry: "Indonesia",
		NotificationUrl: "http://ptsv2.com/t/v2lr8-1532850779/post",
		CartData: "{}",
		UserIP: "127.0.0.1",

		PayMethod: nicepay.MethodCreditCard,
		InstallmentType: nicepay.MerchantCharge,
		InstallmentMonth: 1,
		RecurringOption: nicepay.ReccAutomaticCancel,
	}

	client.GenerateMerchantToken(request)

	res, _, err := client.Registration(request)

	if err != nil {
		log.Fatalln(err)
	}

	if res.ResultCd != "0000" {
		log.Fatal(res)
	}

	fmt.Fprint(w, res.TransactionID)
}

func Pay(w http.ResponseWriter, r *http.Request) {
	s, _ := json.Marshal(nicepay.RegistrationRequest{

	})
	fmt.Fprint(w, string(s))
}

func Cancel(w http.ResponseWriter, r *http.Request) {

}

func Status(w http.ResponseWriter, r *http.Request) {

}

func main() {
	merchantId = os.Getenv("NICEPAY_MERCHANT_ID")
	merchantKey = os.Getenv("NICEPAY_MERCHANT_KEY")

	client = nicepay.NewClient(nil)
	client.BaseUrl = nicepay.EnvProduction
	client.MerchantId = merchantId
	client.MerchantKey = merchantKey

	//registerCreditCard(client)

	router := mux.NewRouter()
	router.HandleFunc("/register", Register)
	router.HandleFunc("/pay", Pay)
	router.HandleFunc("/cancel", Cancel)
	router.HandleFunc("/status", Status)

	http.ListenAndServe(":8080", router)
}
