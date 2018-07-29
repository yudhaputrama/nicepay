package nicepay

import (
	"net/http"
	"net/url"
	"io"
	"bytes"
	"encoding/json"
	"crypto/sha256"
	"strconv"
	"encoding/hex"
	"os"
)

type ErrorCredential struct {
	Message string
}

func (e *ErrorCredential) Error() string {
	return e.Message
}

type ClientEnv string


const (
	EnvProduction  ClientEnv = "https://api.nicepay.co.id"
	EnvDevelopment ClientEnv = "https://qa.nicepay.co.id"
)

type Client struct {
	BaseUrl     ClientEnv
	MerchantId  string
	MerchantKey string
	Client      *http.Client
}

type BaseRequest struct {
	TimeStamp	string `json:"timeStamp"`
	MerchantId	string `json:"iMid"`
	ReferenceNo string `json:"referenceNo"`
	TransactionID           string      `json:"tXid"`
	Amount		int `json:"amt"`
	MerchantToken string `json:"merchantToken"`
}

func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error){
	u, err := url.Parse(string(c.BaseUrl) + path)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		b := new(bytes.Buffer)
		err := json.NewEncoder(b).Encode(body)
		if err != nil {
			return nil, err
		}
		buf = b
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	return req, nil
}

func (c *Client) Do(r *http.Request, v interface{}) (*http.Response, error) {
	// send the request
	resp, err := c.Client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&v)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	envUrl := EnvDevelopment

	if os.Getenv("NICEPAY_ENV") == "PRODUCTION" {
		envUrl = EnvProduction
	}

	c := &Client{
		BaseUrl: envUrl,
		Client:httpClient,
	}

	return c
}

func (c *Client) checkCredential() error {
	if c.MerchantKey == "" {
		return &ErrorCredential{
			Message: "MerchantKey must be defined",
		}
	}

	if c.MerchantId == "" {
		return &ErrorCredential{
			Message: "MerchantId must be defined",
		}
	}

	return nil
}

func (c *Client) Registration(request *RegistrationRequest) (*RegistrationResponse, *http.Response, error) {
	u, err := url.Parse("/nicepay/direct/v2/registration")
	if err != nil {
		return nil, nil, err
	}


	// Skip generate merchant token if defined manual
	if request.MerchantToken == "" {
		err = c.GenerateMerchantToken(request)
		if err != nil {
			return nil, nil, err
		}
	}

	req, err := c.NewRequest("POST", u.String(), request)
	if err != nil {
		return nil, nil, err
	}

	createRes := &RegistrationResponse{}
	resp, err := c.Do(req, createRes)
	if err != nil {
		return nil, nil, err
	}

	return createRes, resp, err
}

func (c *Client) PaymentCreditCard(request *PaymentRequest) (*http.Response, error){
	u, err := url.Parse("/nicepay/direct/v2/payment")
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	baseUrl := string(c.BaseUrl) + u.String()

	resp, err := c.Client.PostForm(baseUrl, url.Values{
		"timeStamp": 		{request.TimeStamp},
		"tXid": 			{request.TransactionID},
		"merchantToken":	{request.MerchantToken},
		"cardNo": 			{request.CardNo},
		"cardExpYymm":		{request.CardExp},
		"cardCvv":			{request.CardCvv},
		"cardHolderNm":		{request.CardHolderName},
		"recurringToken":	{request.CardRecurringToken},
		"preauthToken":		{request.CardPreauthToken},
		"clickPayNo":		{request.ClickPayNo},
		"dataField3":		{request.ClickToken3},
		"clickPayToken":	{request.ClickPayToken},
		"callBackUrl":		{request.CallBackURL},
	})

	return resp, err
}

type ErrorGenerateMerchantToken struct {
	Message string
}

func (e *ErrorGenerateMerchantToken) Error() string {
	return e.Message
}

func hashSHA256(raw string) string {
	hasher := sha256.New()
	hasher.Write([]byte(raw))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (c *Client) GenerateMerchantToken(request *RegistrationRequest) error {
	err := c.checkCredential()
	if err != nil {
		return err
	}
	if request.TimeStamp == "" {
		return &ErrorGenerateMerchantToken{Message: "Timestamp must be defined"}
	}
	if request.ReferenceNo == "" {
		return &ErrorGenerateMerchantToken{Message: "ReferenceNo must be defined"}
	}
	if request.Amount < MINIMUM_TRANSACTION_AMOUNT {
		return &ErrorGenerateMerchantToken{Message: "Amount must be greater than minimum transaction amout"}
	}

	rawToken := request.TimeStamp + c.MerchantId + request.ReferenceNo + strconv.Itoa(request.Amount) + c.MerchantKey
	token := hashSHA256(rawToken)

	request.MerchantId = c.MerchantId
	request.MerchantToken = token
	return nil
}

type ErrorResponse struct {
	ResultCode string `json:"resultCd"`
	ResultMessage string `json:"resultMsg"`
}

func (e *ErrorResponse) Error() string {
	msg := "Nicepay returned those error messages:  "
	return msg + e.ResultCode + " - " +e.ResultMessage
}

func CheckResponse(r *http.Response) error {

	switch r.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusInternalServerError:
		return &ErrorResponse{
			ResultMessage: "Internal Server Error",
			ResultCode: "9999",
		}
	default:
		var errResp ErrorResponse
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(errResp)
		if err != nil {
			errResp.ResultMessage = "Couldn't decode response body JSON"
		}
		return &errResp
	}
}