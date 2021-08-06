package document

import (
	"fmt"
	"time"

	"github.com/cosmos-gaminghub/explorer-backend/orm"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
)

const (
	CollectionNmCommonTx = "txs"
	TxStatusSuccess      = "success"
	TxStatusFail         = "fail"

	Tx_Field_Time        = "timestamp"
	Tx_Field_Height      = "height"
	Tx_Field_Hash        = "txhash"
	Tx_Field_From        = "from"
	Tx_Field_To          = "to"
	Tx_Field_Signers     = "signers"
	Tx_Field_Amount      = "amount"
	Tx_Field_Type        = "type"
	Tx_Field_Fee         = "fee"
	Tx_Field_Memo        = "memo"
	Tx_Field_Status      = "status"
	Tx_Field_Code        = "code"
	Tx_Field_Log         = "log"
	Tx_Field_GasUsed     = "gas_used"
	Tx_Field_GasPrice    = "gas_price"
	Tx_Field_ActualFee   = "actual_fee"
	Tx_Field_ProposalId  = "proposal_id"
	Tx_Field_Tags        = "tags"
	Tx_Field_Msgs        = "msgs"
	Tx_Field_Value       = "logs.events.attributes.value"
	Tx_Field_Event_Type  = "logs.events.type"
	Tx_Field_Event_Value = "logs.events.attributes.value"
	Tx_Field_Event_Key   = "logs.events.attributes.key"

	Tx_Field_Msgs_UdInfo         = "msgs.msg.ud_info.source"
	Tx_Field_Msgs_Moniker        = "msgs.msg.moniker"
	Tx_Field_Msgs_UdInfo_Symbol  = "msgs.msg.ud_info.symbol"
	Tx_Field_Msgs_UdInfo_Gateway = "msgs.msg.ud_info.gateway"
	Tx_Field_Msgs_Hashcode       = "msgs.msg.hash_lock"
	Tx_AssetType_Native          = "native"
	Tx_AssetType_Gateway         = "gateway"

	Tx_Asset_TxType_Issue                = "IssueToken"
	Tx_Asset_TxType_Edit                 = "EditToken"
	Tx_Asset_TxType_Mint                 = "MintToken"
	Tx_Asset_TxType_TransferOwner        = "TransferTokenOwner"
	Tx_Asset_TxType_TransferGatewayOwner = "TransferGatewayOwner"
)

type Signer struct {
	AddrHex    string `bson:"addr_hex"`
	AddrBech32 string `bson:"addr_bech32"`
}

type Coin struct {
	Denom  string  `bson:"denom"`
	Amount float64 `bson:"amount"`
}

func (c Coin) Add(a Coin) Coin {
	if c.Denom == a.Denom {
		return Coin{
			Denom:  c.Denom,
			Amount: c.Amount + a.Amount,
		}
	}
	return c
}

type Coins []Coin

type Fee struct {
	Amount Coins `bson:"amount"`
	Gas    int64 `bson:"gas"`
}

type ActualFee struct {
	Denom  string  `bson:"denom"`
	Amount float64 `bson:"amount"`
}

type CommonTx struct {
	Time       time.Time         `bson:"time"`
	Height     int64             `bson:"height"`
	TxHash     string            `bson:"tx_hash"`
	From       string            `bson:"from"`
	To         string            `bson:"to"`
	Amount     Coins             `bson:"amount"`
	Type       string            `bson:"type"`
	Fee        Fee               `bson:"fee"`
	Memo       string            `bson:"memo"`
	Status     string            `bson:"status"`
	Code       uint32            `bson:"code"`
	Log        string            `bson:"log"`
	GasUsed    int64             `bson:"gas_used"`
	GasWanted  int64             `bson:"gas_wanted"`
	GasPrice   float64           `bson:"gas_price"`
	ActualFee  ActualFee         `bson:"actual_fee"`
	ProposalId uint64            `bson:"proposal_id"`
	Tags       map[string]string `bson:"tags"`
	Msgs       []MsgItem         `bson:"msgs"`
	Signers    []Signer          `bson:"signers"`
}

func (tx CommonTx) String() string {
	return ""

}

type (
	Msg interface {
		Type() string
		String() string
	}
	MsgItem struct {
		Type    string      `bson:"type"`
		MsgData interface{} `bson:"msg"`
	}
	StakeCreateValidator struct {
		PubKey      string         `bson:"pub_key"`
		Description ValDescription `bson:"description"`
		Commission  CommissionMsg  `bson:"commission"`
	}
	CommissionMsg struct {
		Rate          string `bson:"rate"`            // the commission rate charged to delegators
		MaxRate       string `bson:"max_rate"`        // maximum commission rate which validator can ever charge
		MaxChangeRate string `bson:"max_change_rate"` // maximum daily increase of the validator commission
	}
	StakeEditValidator struct {
		CommissionRate string         `bson:"commission_rate"`
		Description    ValDescription `bson:"description"`
	}
)

// Description
type ValDescription struct {
	Moniker  string `bson:"moniker"`
	Identity string `bson:"identity"`
	Website  string `bson:"website"`
	Details  string `bson:"details"`
}

type Counter []struct {
	Type  string `bson:"_id,omitempty"`
	Count int
}

func (d CommonTx) Name() string {
	return CollectionNmCommonTx
}

func (d CommonTx) PkKvPair() map[string]interface{} {
	return bson.M{Tx_Field_Hash: d.TxHash}
}

func (d CommonTx) EnsureIndexes() []mgo.Index {
	indexes := []mgo.Index{
		{
			Key:        []string{Tx_Field_Time},
			Background: true,
		},
		{
			Key:        []string{Tx_Field_Height},
			Background: true,
		},
		{
			Key:        []string{Tx_Field_Hash},
			Background: true,
			Unique:     true,
		},
		{
			Key:        []string{Tx_Field_Event_Type, Tx_Field_Event_Key, Tx_Field_Event_Value},
			Background: true,
		},
	}

	return indexes
}

func (d CommonTx) Batch(txs []txn.Op) error {
	return orm.Batch(txs)
}

func (cArr Counter) String() string {
	res := ""
	for k, v := range cArr {
		res += fmt.Sprintf("idx: %v Type  :%v  \t	Count :%v \n", k, v.Type, v.Count)
	}
	return res
}
