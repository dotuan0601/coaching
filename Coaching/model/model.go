package model

type Transaction struct {
	Id                       int64
	TraceNumber              string
	Block_id                 int64
	Txhash                   string
	Txcount                  int64
	Chaincodename            string
	Status                   int64
	Creator_msp_id           string
	Endorser_msp_id          string
	Chaincode_id             string
	Write_set                string
	Block_hash               string
	Channel_genesis_hash     string
	Validation_code          string
	envelope_signature       string
	Payload_extension        string
	Creator_id_bytes         string
	Creator_nonce            string
	Chaincode_proposal_input string
	Tx_response              string
	Payload_proposal_hash    string
	Endorser_id_bytes        string
	Endorser_signature       string
	Network_name             string
}

// receiver function
// func (tran *Transaction) UpdateTran(txtCount int64) {
// 	tran.Txcount = txtCount
// }
