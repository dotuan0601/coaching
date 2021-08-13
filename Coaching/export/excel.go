package export

import (
	"github.com/xuri/excelize/v2"
)

func excel(t *Transactions) {
	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", t.id)
	f.SetCellValue("Sheet1", "B1", t.trace_number)
	f.SetCellValue("Sheet1", "C1", t.blockid)
	f.SetCellValue("Sheet1", "D1", t.txhash)
	f.SetCellValue("Sheet1", "E1", t.txcount)
	f.SetCellValue("Sheet1", "G1", t.chaincodename)
	f.SetCellValue("Sheet1", "H1", t.status)
	f.SetCellValue("Sheet1", "I1", t.creator_msp_id)
	f.SetCellValue("Sheet1", "J1", t.endorser_msp_id)
	f.SetCellValue("Sheet1", "K1", t.chaincode_id)
	f.SetCellValue("Sheet1", "M1", t.block_hash)
	f.SetCellValue("Sheet1", "N1", t.channel_genesis_hash)
	f.SetCellValue("Sheet1", "O1", t.validation_code)
	f.SetCellValue("Sheet1", "P1", t.envelope_signature)
	f.SetCellValue("Sheet1", "Q1", t.payload_extension)
	f.SetCellValue("Sheet1", "R1", t.creator_id_bytes)
	f.SetCellValue("Sheet1", "S1", t.creator_nonce)
	f.SetCellValue("Sheet1", "T1", t.tx_response)
	f.SetCellValue("Sheet1", "U1", t.payload_proposal_hash)
	f.SetCellValue("Sheet1", "V1", t.endorser_id_bytes)
	f.SetCellValue("Sheet1", "W1", t.network_name)
	f.SetCellValue("Sheet1", "X1", t.endorser_signature)

}
