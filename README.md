# Nicepay Client

## Usage
```go
merchantId := os.Getenv("NICEPAY_MERCHANT_ID")
merchantKey := os.Getenv("NICEPAY_MERCHANT_KEY")

client := nicepay.NewClient(nil)
client.BaseUrl = nicepay.EnvProduction
client.MerchantId = merchantId
client.MerchantKey = merchantKey
```

## Supported
|Methods|Value|Description|
|------------------------|----|-------------|
| MethodCreditCard | 01 |  |
| MethodVirtualAccount | 02 |  |
| MethodConvenienceStore | 03 |  |
| MethodClickPay | 04 | Coming Soon |
| MethodEWallet | 05 | Coming Soon |

### Registration

#### Credit Card
|Property|Type Data|Description|
|------------------|----|-------------|
| InstallmentType  | ||
| InstallmentMonth | ||
| RecurringOption  | ||

```go

// Registrantion
request := &nicepay.RegistrationRequest{
    TimeStamp: time.Now().Format(nicepay.FORMAT_TIMESTAMP),
    MerchantId: c.MerchantId,
    Amount: 20000,
    Currency: "IDR",
    ReferenceNo: "1234",
    GoodsName: "Nicepay Pillow",
    Description: "Jalan Jalan",
    BillingName:"Nicepay",
    BillingPhone:"0812345678",
    BillingEmail:"hello@test.com",
    BillingAddress: "Suka Suka",
    BillingCity: "Bandung",
    BillingState: "Jawa Barat",
    BillingPostalCode: "55555",
    BillingCountry: "Indonesia",
    NotificationUrl: "http://ptsv2.com/t/s6v4a-1530243879/post",
    CartData: "{}",
    UserIP: "127.0.0.1",

    PayMethod: nicepay.MethodCreditCard,
    InstallmentType: nicepay.MerchantCharge,
    InstallmentMonth: 1,
    RecurringOption: nicepay.ReccAutomaticCancel,
}

client.GenerateMerchantToken(request)

res, _, err := c.Registration(request)

// Payment
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
```

#### Virtual Account
|Property|Type Data|Description|
|------------------|----|-------------|
| InstallmentType  | ||
| InstallmentMonth | ||
| RecurringOption  | ||

```go
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
    BillingPhone:"0812345678",
    BillingEmail:"hello@test.com",
    BillingAddress: "Bandung",
    BillingCity: "Bandung",
    BillingState: "Jawa Barat",
    BillingPostalCode: "55555",
    BillingCountry: "Indonesia",
    NotificationUrl: "http://requestbin.fullcontact.com/1fexw2h1",
    CartData: "{}",

    PayMethod: nicepay.MethodVirtualAccount,
    Bank: nicepay.BankBNI,
    VirtualAccountValidDate: validVA.Format(nicepay.FORMAT_DATE),
    VirtualAccountValidTime: validVA.Format(nicepay.FORMAT_TIME),
})
```

#### ConvenienceStore
|Property|Type Data|Description|
|------------------|----|-------------|
| Bank  | ||
| VirtualAccountValidDate | ||
| VirtualAccountValidTime  | ||
| MerchantReservedVAID  | ||

### Type Data
##### InstallmentType
|Property   |Value|Description|
|----------------|---|-----------|
| CustomerCharge | 1 ||
| MerchantCharge | 2 ||

##### Recuring Option
|Property|Value|Description|
|------------------------|----|-------------|
| ReccAutomaticCancel | 1 ||
| ReccDoNotCancel | 2 ||
| ReccDoNotMakeToken | 3 ||

##### Bank
|Property|Value|Description|
|------------------------|----|-------------|
| BankMandiri | BMRI ||
| BankMaybank | IBBK ||
| BankPermata | BBBA ||
| BankBCA | CENA ||
| BankBNI | BNIN ||
| BankHana | HNBN ||
| BankBRI | BRIN ||
| BankCimbNiaga | BNIA ||
| BankDanamon | BDIN ||
| BankOther | OTHR ||

##### Mitra
|Property|Value|Description|
|------------------------|----|-------------|
| MitraAlfamart | ALMA ||
| MitraIndomaret | INDO ||
| MitraLawson | LOSN ||
| MitraAlfaMidi | ALMI ||
| MitraDanDan | DNDN ||
| MitraClickPayMandiri | MDRC ||
| MitraClickPayBCA | BCAC ||
| MitraClickPayCimb | CIMC ||
| MitraWalletMandiri | MDRE ||
| MitraWalletSakuku | BCAE ||