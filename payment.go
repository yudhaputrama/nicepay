package nicepay

import "time"

type PaymentRequest struct {
	TimeStamp      string      `json:"timeStamp"`
	TransactionID           string      `json:"tXid"`
	MerchantToken  string      `json:"merchantToken"`
	CardNo         string      `json:"cardNo"`
	CardExp		   string      `json:"cardExpYymm"`
	CardCvv        string      `json:"cardCvv"`
	CardHolderName   string      `json:"cardHolderNm"` //Cimb Niaga
	CardRecurringToken string `json:"recurringToken"`
	CardPreauthToken   string `json:"preauthToken"`
	ClickPayNo     string `json:"clickPayNo"`
	ClickToken3     string `json:"dataField3"`
	ClickPayToken  string `json:"clickPayToken"`
	CallBackURL    string      `json:"callBackUrl"`
}


type PaymentRespond struct {
	ResultCd       string      `json:"resultCd"`
	ResultMsg      string      `json:"resultMsg"`
	TransactionID  string      `json:"tXid"`
	ReferenceNo    string      `json:"referenceNo"`
	PayMethod      string      `json:"payMethod"`
	Amount		string      `json:"amt"`
	TransactionDate string      `json:"transDt"`
	TransactionTime string      `json:"transTm"`
	Description    string      `json:"description"`
	AuthNo         string      `json:"authNo"`
	IssueBank     Bank	`json:"issuBankCd"`
	AcquireBank     Bank	`json:"acquBankCd"`
	CardNo         string      `json:"cardNo"`
	ReceiptCode    string `json:"receiptCode"`
	Mitra MitraCode `json:"mitraCd"`
	RecurringToken string `json:"recurringToken"`
	PreauthToken   string `json:"preauthToken"`
	Currency       string      `json:"currency"`
	GoodsNm        string      `json:"goodsNm"`
	BillingNm      string      `json:"billingNm"`
	CcTransType    string      `json:"ccTransType"`
	MReferenceNo         string `json:"mRefNo"`
	InstallmentType    string      `json:"instmntType"`
	InstallmentMonth     string      `json:"instmntMon"`
	CardExp    string      `json:"cardExpYymm"`
}

type StatusRequest struct {
	TimeStamp     string `json:"timeStamp"`
	TransactionID string `json:"tXid"`
	MerchantID    string `json:"iMid"`
	ReferenceNo   string `json:"referenceNo"`
	Amount        string `json:"amt"`
	MerchantToken string `json:"merchantToken"`
}

type StatusResponse struct {
	ResultCode       string      `json:"resultCd"`
	ResultMessage  string      `json:"resultMsg"`
	TransactionID        string      `json:"tXid"`
	MerchantID     string      `json:"iMid"`
	Currency       string      `json:"currency"`
	Amount       string      `json:"amt"`
	ReferenceNo    string      `json:"referenceNo"`
	GoodsNm        string      `json:"goodsNm"`
	PayMethod      string      `json:"payMethod"`
	BillingNm      string      `json:"billingNm"`
	RequestDate    string      `json:"reqDt"`
	RequestTime    string      `json:"reqTm"`
	Status         string      `json:"status"`
	CancelAmount   int `json:"cancelAmt"`
	TransactionDate time.Time `json:"transDt"`
	TransactionTime time.Time `json:"transTm"`
	MRefNo string `json:"mRefNo"`

	CardNo         string `json:"cardNo"`
	PreauthToken   string `json:"preauthToken"`
	AcquBankCode   string `json:"acquBankCd"`
	IssuBankCode   string `json:"issuBankCd"`
	AcquBankName   string `json:"acquBankNm"`
	IssuBankName   string `json:"issuBankNm"`
	InstmntMonth   int `json:"instmntMon"`
	InstmntType    string      `json:"instmntType"`

	VacctValidDt   string `json:"vacctValidDt"`
	VacctValidTm   string `json:"vacctValidTm"`
	VacctNo        string `json:"vacctNo"`
	CardExpYymm    string `json:"cardExpYymm"`
	BankCd         string `json:"bankCd"`
	RecurringToken string `json:"recurringToken"`
	CcTransType    string `json:"ccTransType"`
	AcquStatus     string `json:"acquStatus"`

	PayNo          string `json:"payNo"`
	MitraCode      string `json:"mitraCd"`
	ReceiptCode    string `json:"receiptCode"`
	PayValidDt     string `json:"payValidDt"`
	PayValidTm     string `json:"payValidTm"`
}

type CancelRequest struct {
	TimeStamp      string `json:"timeStamp"`
	TransactionID  string `json:"tXid"`
	MerchantID     string `json:"iMid"`
	PayMethod      string `json:"payMethod"`
	CancelType     string `json:"cancelType"`
	CancelMessage  string `json:"cancelMsg"`
	MerchantToken  string `json:"merchantToken"`
	PreauthToken   string `json:"preauthToken"`
	Amount         string `json:"amt"`
	CancelServerIP string `json:"cancelServerIp"`
	CancelUserID   string `json:"cancelUserId"`
	CancelUserIP   string `json:"cancelUserIp"`
	CancelUserInfo string `json:"cancelUserInfo"`
	CancelRetryCount string `json:"cancelRetryCnt"`
	Worker         string `json:"worker"`
}

type CancelResponse struct {
	TransactionID string `json:"tXid"`
	ReferenceNo string `json:"referenceNo"`
	ResultCode    string `json:"resultCd"`
	ResultMessage   string `json:"resultMsg"`
	TransactionDate     string `json:"transDt"`
	TransactionTime     string `json:"transTm"`
	Description string `json:"description"`
	Amount	string `json:"amt"`
}