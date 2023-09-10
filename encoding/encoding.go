package main

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func encodeBase64(m string) {
	fmt.Println("plaintext:")
	fmt.Println(m)

	encodedString := base64.StdEncoding.EncodeToString([]byte(m))
	fmt.Println("base64 text:")
	fmt.Println(encodedString)

	fmt.Println("decoded base64 text:")
	decodedString, _ := base64.StdEncoding.DecodeString(encodedString)
	fmt.Println(string(decodedString))
}

func encodeHex(m string) {
	encodedString := hex.EncodeToString([]byte(m))
	fmt.Println("HEX text:")
	fmt.Println(encodedString)

	fmt.Println("decoded hex text:")
	decodedString, _ := hex.DecodeString(encodedString)
	fmt.Println(string(decodedString))
}

func jsonDemo() {
	type Customer struct {
		CustomerId   int    `json:"customer_id"`
		CustomerName string `json:"customer_name"`
		SSN          string `json:"ssn"`
	}
	fmt.Printf("\n\n============STRING TO JSON =========\n\n")
	c := Customer{1, "Sigma", "123-45-6789"}
	jsonBody, err := json.Marshal(&c)
	if err != nil {
		fmt.Println("Error in Marshalling")
		return
	} else {
		fmt.Println(string(jsonBody))
	}

	fmt.Printf("\n\n============JSON TO STRING=========\n\n")
	var cust Customer
	json.Unmarshal(jsonBody, &cust)
	fmt.Printf("%#v", cust)
}

func main() {
	message := "hello,go (*w3hu%#"
	encodeBase64(message)
	encodeHex(message)
	jsonDemo()
}
