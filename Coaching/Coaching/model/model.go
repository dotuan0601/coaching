package model

import "time"

type Transaction struct {
	Id                       int64     `json:"id"`
	TraceNumber              string    `json:"tracenumber"`
	Blockid                  int64     `json:"blockid"`
	Txhash                   string    `json:"txhash"`
	Createdt                 time.Time `json:"createdt"`
	Chaincodename            string    `json:"chaincodename"`
	Status                   int64     `json:"status"`
	Creator_msp_id           string    `json:"creator_msp_id"`
	Endorser_msp_id          string    `json:"endorser_msp_id"`
	Chaincode_id             string    `json:"chaincode_id"`
	Type                     string    `json:"type"`
	Read_set                 string    `json:"read_set`
	Write_set                string    `json:"write_set"`
	Block_hash               string    `json:"block_hash"`
	Channel_genesis_hash     string    `json:"channel_genesis_hash"`
	Validation_code          string    `json:"validation_code"`
	Envelope_signature       string    `json:"envelope_signature"`
	Payload_extension        string    `json:"payload_extension"`
	Creator_id_bytes         string    `json:"creator_id_bytes"`
	Creator_nonce            string    `json:"creator_nonce "`
	Chaincode_proposal_input string    `json:"chaincode_proposal_input"`
	Tx_response              string    `json:"tx_response"`
	Payload_proposal_hash    string    `json:"payload_proposal_hash"`
	Endorser_id_bytes        string    `json:"endorser_id_bytes"`
	Endorser_signature       string    `json:"endorser_signature"`
	Network_name             string    `json:"network_name"`
}
