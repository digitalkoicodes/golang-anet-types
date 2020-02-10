package anet

type GetMerchantDetails struct {
	GetMerchantDetailsRequest *GetMerchantDetailsRequest `json:"getMerchantDetailsRequest"`
}

type GetMerchantDetailsResponse struct {
	PublicClientKey string `json:"publicClientKey"`
	LoginId         string `json:"loginId"`
}

type GetMerchantDetailsRequest struct {
	MerchantAuthentication *MerchantAuthentication `json:"merchantAuthentication"`
}

type MerchantAuthentication struct {
	Name           string `json:"name"`
	TransactionKey string `json:"transactionKey"`
}

type Payment struct {
	OpaqueData *OpaqueData `json:"opaqueData"`
}

type OpaqueData struct {
	DataDescriptor string `json:"dataDescriptor"`
	DataValue      string `json:"dataValue"`
}

type CreateTransaction struct {
	CreateTransactionRequest *CreateTransactionRequest `json:"createTransactionRequest"`
}

type CreateTransactionRequest struct {
	MerchantAuthentication *MerchantAuthentication `json:"merchantAuthentication"`
	TransactionRequest     *TransactionRequest     `json:"transactionRequest"`
}

type TransactionRequest struct {
	TransactionType string    `default:"authCaptureTransaction" json:"transactionType"`
	Amount          float64   `json:"amount"`
	Payment         *Payment  `json:"payment"`
	Order           *Order    `json:"order"`
	Customer        *Customer `json:"customer"`
}

type Order struct {
	InvoiceNumber string `json:"invoiceNumber"`
}

type Customer struct {
	Email string `json:"email"`
}

type TransactionResponse struct {
	ResponseCode  string   `json:"responseCode"`
	TransId       string   `json:"transId"`
	AccountNumber string   `json:"accountNumber"`
	AccountType   string   `json:"accountType"`
	Errors        []*Error `json:"errors"`
}

type Error struct {
	ErrorCode string `json:"errorCode"`
	ErrorText string `json:"errorText"`
}

type Response struct {
	TransactionResponse *TransactionResponse `json:"transactionResponse"`
	Messages            *Messages            `json:"messages"`
}

type Messages struct {
	ResultCode string     `json:"resultCode"`
	Message    []*Message `json:"message"`
}

type Message struct {
	Code string `json:"code"`
	Text string `json:"text"`
}

func NewGetMerchantDetails(name string, key string) *GetMerchantDetails {
	return &GetMerchantDetails{
		GetMerchantDetailsRequest: &GetMerchantDetailsRequest{
			MerchantAuthentication: &MerchantAuthentication{
				Name:           name,
				TransactionKey: key,
			},
		},
	}
}
