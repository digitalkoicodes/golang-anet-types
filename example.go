package example

import (
	"anet"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"fmt"
)

//test
main(){
	opaqueData := &anet.OpaqueData{} //parse post request into this
	request := &anet.CreateTransaction{
		CreateTransactionRequest: &anet.CreateTransactionRequest{
			MerchantAuthentication: &anet.MerchantAuthentication{
				Name:           "API NAME",
				TransactionKey: "API KEY",
			},
			TransactionRequest: &anet.TransactionRequest{
				TransactionType: "authCaptureTransaction",
				Amount:          1,
				Payment: &anet.Payment{
					OpaqueData: opaqueData,
				},
				Order: &anet.Order{
					InvoiceNumber: "XXX1232",
				},
				Customer: &anet.Customer{
					Email: "example@example.org"
				},
			},
		},
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(request)
	req, _ := http.NewRequest("POST", "https://apitest.authorize.net/xml/v1/request.api", buf)
	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		return e
	}
	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	response := &anet.Response{}
	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println(response);
}