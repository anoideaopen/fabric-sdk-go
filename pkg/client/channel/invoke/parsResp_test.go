package invoke

import (
	"bytes"
	"encoding/base64"
	"testing"

	prt "github.com/golang/protobuf/proto"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric-sdk-go/internal/github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/rwsetutil"
	"github.com/hyperledger/fabric-sdk-go/internal/github.com/hyperledger/fabric/protoutil"
	"github.com/stretchr/testify/assert"
	fproto "gitlab.n-t.io/atmz/foundation/proto"
)

var (
	ChaincodeActionStr1 = ""
	ChaincodeActionStr2 = ""
	BatchResponseStr1   = ""
	BatchResponseStr2   = ""
)

func TestParsing(t *testing.T) {
	t.Parallel()

	BatchResponseByte1, err1 := base64.StdEncoding.DecodeString(BatchResponseStr1)
	assert.NoError(t, err1)

	BatchResponseByte2, err2 := base64.StdEncoding.DecodeString(BatchResponseStr2)
	assert.NoError(t, err2)

	if !bytes.Equal(BatchResponseByte1, BatchResponseByte2) {
		out1 := &fproto.BatchResponse{}
		err1 = prt.Unmarshal(BatchResponseByte1, out1)
		assert.NoError(t, err1)

		out2 := &fproto.BatchResponse{}
		err2 = prt.Unmarshal(BatchResponseByte2, out2)
		assert.NoError(t, err2)

		if len(out1.SwapResponses) != len(out2.SwapResponses) {
			t.Log("1")
		} else if len(out1.SwapResponses) != 0 {
			for i := range out1.SwapResponses {
				switch {
				case out1.SwapResponses[i].Error != out2.SwapResponses[i].Error:
					t.Log("error")
				case !bytes.Equal(out1.SwapResponses[i].Id, out2.SwapResponses[i].Id):
					t.Log("error2")
				case len(out1.SwapResponses[i].Writes) != len(out2.SwapResponses[i].Writes):
					t.Log("error3")
				default:
					for j := range out1.SwapResponses[i].Writes {
						switch {
						case out1.SwapResponses[i].Writes[j].Key != out2.SwapResponses[i].Writes[j].Key:
							t.Log("error4")
						case out1.SwapResponses[i].Writes[j].IsDeleted != out2.SwapResponses[i].Writes[j].IsDeleted:
							t.Log("error5")
						case !bytes.Equal(out1.SwapResponses[i].Writes[j].Value, out2.SwapResponses[i].Writes[j].Value):
							t.Log("error6")
						}
					}
				}
			}
		}

		if len(out1.SwapKeyResponses) != len(out2.SwapKeyResponses) {
			t.Log("2")
		} else if len(out1.SwapKeyResponses) != 0 {
			for i := range out1.SwapKeyResponses {
				switch {
				case out1.SwapKeyResponses[i].Error != out2.SwapKeyResponses[i].Error:
					t.Log("error20")
				case !bytes.Equal(out1.SwapKeyResponses[i].Id, out2.SwapKeyResponses[i].Id):
					t.Log("error22")
				case len(out1.SwapKeyResponses[i].Writes) != len(out2.SwapKeyResponses[i].Writes):
					t.Log("error23")
				default:
					for j := range out1.SwapKeyResponses[i].Writes {
						switch {
						case out1.SwapKeyResponses[i].Writes[j].Key != out2.SwapKeyResponses[i].Writes[j].Key:
							t.Log("error24")
						case out1.SwapKeyResponses[i].Writes[j].IsDeleted != out2.SwapKeyResponses[i].Writes[j].IsDeleted:
							t.Log("error25")
						case !bytes.Equal(out1.SwapKeyResponses[i].Writes[j].Value, out2.SwapKeyResponses[i].Writes[j].Value):
							t.Log("error26")
						}
					}
				}
			}
		}

		if len(out1.CreatedSwaps) != len(out2.CreatedSwaps) {
			t.Log("3")
		} else if len(out1.CreatedSwaps) != 0 {
			for i := range out1.CreatedSwaps {
				switch {
				case out1.CreatedSwaps[i].Timeout != out2.CreatedSwaps[i].Timeout:
					t.Log("error30")
				case out1.CreatedSwaps[i].From != out2.CreatedSwaps[i].From:
					t.Log("error31")
				case out1.CreatedSwaps[i].To != out2.CreatedSwaps[i].To:
					t.Log("error32")
				case out1.CreatedSwaps[i].Token != out2.CreatedSwaps[i].Token:
					t.Log("error33")
				case !bytes.Equal(out1.CreatedSwaps[i].Creator, out2.CreatedSwaps[i].Creator):
					t.Log("error34")
				case !bytes.Equal(out1.CreatedSwaps[i].Id, out2.CreatedSwaps[i].Id):
					t.Log("error35")
				case !bytes.Equal(out1.CreatedSwaps[i].Amount, out2.CreatedSwaps[i].Amount):
					t.Log("error36")
				case !bytes.Equal(out1.CreatedSwaps[i].Hash, out2.CreatedSwaps[i].Hash):
					t.Log("error37")
				case !bytes.Equal(out1.CreatedSwaps[i].Owner, out2.CreatedSwaps[i].Owner):
					t.Log("error38")
				}
			}
		}

		if len(out1.TxResponses) != len(out2.TxResponses) {
			t.Log("4")
		} else if len(out1.TxResponses) != 0 {
			for i := range out1.TxResponses {
				switch {
				case out1.TxResponses[i].Error != out2.TxResponses[i].Error:
					t.Log("error40")
				case !bytes.Equal(out1.TxResponses[i].Id, out2.TxResponses[i].Id):
					t.Log("error42")
				case len(out1.TxResponses[i].Writes) != len(out2.TxResponses[i].Writes):
					t.Log("error43")
				default:
					for j := range out1.TxResponses[i].Writes {
						switch {
						case out1.TxResponses[i].Writes[j].Key != out2.TxResponses[i].Writes[j].Key:
							t.Log("error44")
						case out1.TxResponses[i].Writes[j].IsDeleted != out2.TxResponses[i].Writes[j].IsDeleted:
							t.Log("error45")
						case !bytes.Equal(out1.TxResponses[i].Writes[j].Value, out2.TxResponses[i].Writes[j].Value):
							t.Log("error46")
						}
					}
				}
			}
		}
	}

	respPayload1 := &pb.ChaincodeAction{}
	respPayload2 := &pb.ChaincodeAction{}

	ChaincodeActionByte1, err1 := base64.StdEncoding.DecodeString(ChaincodeActionStr1)
	assert.NoError(t, err1)

	ChaincodeActionByte2, err2 := base64.StdEncoding.DecodeString(ChaincodeActionStr2)
	assert.NoError(t, err2)

	if !bytes.Equal(ChaincodeActionByte1, ChaincodeActionByte2) {
		ChaincodeAction1, err1 := protoutil.UnmarshalProposalResponsePayload(ChaincodeActionByte1)
		assert.NoError(t, err1)

		if ChaincodeAction1.Extension != nil {
			respPayload1, err1 = protoutil.UnmarshalChaincodeAction(ChaincodeAction1.Extension)
			assert.NoError(t, err1)
		}

		ChaincodeAction2, err2 := protoutil.UnmarshalProposalResponsePayload(ChaincodeActionByte2)
		assert.NoError(t, err2)

		if ChaincodeAction2.Extension != nil {
			respPayload2, err2 = protoutil.UnmarshalChaincodeAction(ChaincodeAction2.Extension)
			assert.NoError(t, err2)
		}

		txRWSet1 := &rwsetutil.TxRwSet{}
		txRWSet2 := &rwsetutil.TxRwSet{}

		err1 = txRWSet1.FromProtoBytes(respPayload1.Results)
		assert.NoError(t, err1)

		err2 = txRWSet2.FromProtoBytes(respPayload2.Results)
		assert.NoError(t, err2)
	}
}
