package api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTransactions(t *testing.T) {
	req, err := http.NewRequest("GET", "/get", nil)
	if err != nil {
		t.Fatal(err)
	}
	record := httptest.NewRecorder()
	handler := http.HandlerFunc(GetTransaction)
	handler.ServeHTTP(record, req)
	fmt.Println(req)
	if status := record.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	// expected := `[{

	// 	"tracenumber" : " test",
	// 	"blockid" : 2,
	// 	 "txhash" : "test",
	// 	 "creatdt":"12:00:00",
	// 	 "chaincodename" : "test",
	// 	 "status" : 3 ,
	// 	 "creator_msp_id" : "test",
	// 	 "endorser_msp_id" : " test",
	// 	 "chaincode_id" : "test" ,
	// 	 "type" :" test" ,
	// 	 "read_set" : "test",
	// 	 "write_set" : "test",
	// 	 "block_hash" : "test" ,
	// 	 "channel_genesis_hash" : "test",
	// 	 "validation_code" : "test" ,
	// 	 "envelope_signature" : "test" ,
	// 	 "payload_extension": "test",
	// 	 "creator_id_bytes": "test",
	// 	 "creator_nonce" : "test",
	// 	 "chaincode_proposal_input": "test" ,
	// 	 "tx_response"  : "test",
	// 	 "payload_proposal_hash" : "test",
	// 	 "endorser_id_bytes" : "test",
	// 	 "endorser_signature":"test",
	// 	 "network_name" :" test"

	//  }]`
	// if strings.TrimSpace(record.Body.String()) != expected {
	// 	t.Errorf("handler returned unexpected bot: got %v want %v", record.Body.String(), expected)
	// }
}
func TestCreateTransaction(t *testing.T) {
	var jsonStr = []byte(`{
		"tracenumber" : "test3",
		"blockid" : 4,
		 "txhash" : "test",
		 "creatdt":"12:00:00",
		 "chaincodename" : "test",
		 "status" : 3 ,
		 "creator_msp_id" : "test",
		 "endorser_msp_id" : " test",
		 "chaincode_id" : "test" ,
		 "type" :" test" ,
		 "read_set" : "test",
		 "write_set" : "test",
		 "block_hash" : "test" ,
		 "channel_genesis_hash" : "test",
		 "validation_code" : "test" ,
		 "envelope_signature" : "test" ,
		 "payload_extension": "test",
		 "creator_id_bytes": "test",
		 "creator_nonce" : "test",
		 "chaincode_proposal_input": "test" ,
		 "tx_response"  : "test",
		 "payload_proposal_hash" : "test",
		 "endorser_id_bytes" : "test",
		 "endorser_signature":"test",
		 "network_name" :" test"
	 
	 }`)
	req, err := http.NewRequest("POST", "/create", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)

	}
	req.Header.Set("Content-Type", "application/json")
	record := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTransaction)
	handler.ServeHTTP(record, req)
	if status := record.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
