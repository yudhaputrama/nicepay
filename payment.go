package nicepay

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