package schema

import (
	"time"
)

// Code defines the structure for code information.
type Code struct {
	CodeId            int       `bson:"code_id"`
	Contract          string    `bson:"contract"`
	DataHash          string    `bson:"data_hash"`
	CreatedAt         time.Time `bson:"created_at"`
	Creator           string    `bson:"creator"`
	InstantiateCount  int       `bson:"instantiate_count"`
	Permission        string    `bson:"permission"`
	PermittedAddress  string    `bson:"permitted_address"`
	TxHash            string    `bson:"txhash"`
	Version           string    `bson:"version"`
	FirstContractTime time.Time `bson:"first_contract_time"`
}

// NewCode returns a new Code.
func NewCode() *Code {
	return &Code{}
}

func (code *Code) SetCode(codeId int) *Code {
	code.CodeId = codeId
	return code
}

func (code *Code) SetContract(contract string) *Code {
	code.Contract = contract
	return code
}

func (code *Code) SetDataHash(dataHash string) *Code {
	code.DataHash = dataHash
	return code
}

func (code *Code) SetCreatedAt(createdAt time.Time) *Code {
	code.CreatedAt = createdAt
	return code
}

func (code *Code) SetCreator(creator string) *Code {
	code.Creator = creator
	return code
}

func (code *Code) SetInstantiateCount(count int) *Code {
	code.InstantiateCount = count
	return code
}

func (code *Code) SetPermission(permission string) *Code {
	code.Permission = permission
	return code
}

func (code *Code) SetPermittedAddress(permittedAddress string) *Code {
	code.PermittedAddress = permittedAddress
	return code
}

func (code *Code) SetVersion(version string) *Code {
	code.Version = version
	return code
}

func (code *Code) SetTxhash(txHash string) *Code {
	code.TxHash = txHash
	return code
}

func (code *Code) SetFirstContractTime(time time.Time) *Code {
	code.FirstContractTime = time
	return code
}
