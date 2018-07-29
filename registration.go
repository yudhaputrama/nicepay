package nicepay

import "encoding/json"

type PaymentMethod string

const MINIMUM_TRANSACTION_AMOUNT = 10000

const (
	FORMAT_TIMESTAMP = "20060102150405"
	FORMAT_DATE = "20060102"
	FORMAT_TIME = "150405"
)

const (
	MethodCreditCard 		PaymentMethod = "01"
	MethodVirtualAccount 	PaymentMethod = "02"
	MethodConvenienceStore	PaymentMethod = "03"
	MethodClickPay			PaymentMethod = "04"
	MethodEWallet			PaymentMethod = "05"
)

var AllPaymentMethod = []PaymentMethod{
	MethodCreditCard,
	MethodVirtualAccount,
	MethodConvenienceStore,
	MethodClickPay,
	MethodEWallet,
}

type MitraCode string

const (
	MitraAlfamart MitraCode = "ALMA"
	MitraIndomaret MitraCode = "INDO"
	MitraLawson MitraCode = "LOSN"
	MitraAlfaMidi MitraCode = "ALMI"
	MitraDanDan MitraCode = "DNDN"
	MitraClickPayMandiri = "MDRC"
	MitraClickPayBCA = "BCAC"
	MitraClickPayCimb = "CIMC"
	MitraWalletMandiri = "MDRE"
	MitraWalletSakuku = "BCAE"
)

type InstallmentType int

const (
	CustomerCharge InstallmentType = 1
	MerchantCharge InstallmentType = 2
)

type RecurringOption int

const (
	ReccAutomaticCancel RecurringOption = 0
	ReccDoNotCancel RecurringOption = 1
	ReccDoNotMakeToken RecurringOption = 2
)

type RegistrationRequest struct {
	TimeStamp				string `json:"timeStamp"`
	MerchantId				string `json:"iMid"`
	PayMethod				PaymentMethod `json:"payMethod"`
	Currency				string `json:"currency"`
	Amount					int `json:"amt"`
	ReferenceNo 			string `json:"referenceNo"`
	GoodsName				string `json:"goodsNm"`
	BillingName				string `json:"billingNm"`
	BillingPhone 			string `json:"billingPhone"`
	BillingEmail 			string `json:"billingEmail"`
	BillingAddress 			string `json:"billingAddr"`
	BillingCity				string `json:"billingCity"`
	BillingState			string `json:"billingState"`
	BillingPostalCode 		string `json:"billingPostCd"`
	BillingCountry			string `json:"billingCountry"`
	DeliveryName			string `json:"deliveryNm"`
	DeliveryPhone			string `json:"deliveryPhone"`
	DeliveryAddress			string `json:"deliveryAddr"`
	DeliveryCity			string `json:"deliveryCity"`
	DeliveryState			string `json:"deliveryState"`
	DeliveryPostalCode		string `json:"deliveryPostCd"`
	DeliveryCountry			string `json:"deliveryCountry"`
	NotificationUrl			string `json:"dbProcessUrl"`
	VAT						int `json:"vat"`
	Fee						int `json:"fee"`
	FreeTaxAmount			int `json:"notaxAmt"`
	Description				string `json:"description"`
	MerchantToken			string `json:"merchantToken"`
	RequestDate				string `json:"reqDt"`
	RequestTime				string `json:"reqTm"`
	RequestDomain			string `json:"reqDomain"`
	RequestServerIP			string `json:"reqServerIP"`
	RequestClientVersion	string `json:"reqClientVer"`
	UserIP 					string `json:"userIP"`
	UserSessionID 			string `json:"userSessionID"`
	UserAgent 				string `json:"userAgent"`
	UserLanguage 			string `json:"userLanguage"`
	CartData 				string `json:"cartData"`
	InstallmentType 		InstallmentType `json:"instmntType"`
	InstallmentMonth 		int `json:"instmntMon"`
	RecurringOption 		RecurringOption `json:"recurrOpt"`
	Bank					Bank `json:"bankCd"`
	VirtualAccountValidDate string `json:"vacctValidDt"`
	VirtualAccountValidTime string `json:"vacctValidTm"`
	MerchantReservedVAID	string `json:"merFixAcctId"`
	Mitra					MitraCode `json:"mitraCd"`
}

type CartItem struct {
	ImageUrl 	string 	`json:"img_url"`
	Name 		string 	`json:"goods_name"`
	Detail 		string 	`json:"goods_detail"`
	Amount 		int 	`json:"goods_amt"`
}

type CartData struct {
	Count 		int 		`json:"count"`
	Items 		[]CartItem 	`json:"item"`
}

func (c CartData) String() string {
	s, err := json.Marshal(c)
	if err != nil {
		return "{}"
	}
	return string(s)
}

type RegistrationResponse struct {
	ResultCd				string			`json:"resultCd"`
	ResultMsg				string			`json:"resultMsg"`
	TransactionID			string			`json:"tXid"`
	ReferenceNo				string			`json:"referenceNo"`
	PayMethod				PaymentMethod	`json:"payMethod"`
	Amount					string 			`json:"amt"`
	TransactionDate			string 			`json:"transDt"`
	TransactionTime			string 			`json:"transTm"`
	Description				string 			`json:"description"`
	Bank 					Bank 			`json:"bankCd"`
	VirtualAccountNumber	string      	`json:"vacctNo"`
	VirtualAccountValidDate string     		`json:"vacctValidDt"`
	VirtualAccountValidTime string     		`json:"vacctValidTm"`
	Mitra 					MitraCode 		`json:"mitraCd"`
	PayNumber        		string			`json:"payNo"`
	PayValidDate   			string 			`json:"payValidDt"`
	PayValidTime   			string			`json:"payValidTm"`
	Currency				string			`json:"currency"`
	GoodsName				string			`json:"goodsNm"`
	BillingName				string			`json:"billingNm"`
}