package schema

import (
	"time"
)

// Contract defines the structure for contract information.
type Contract struct {
	CodeId           int       `bson:"code_id"`
	Contract         string    `bson:"contract"`
	ContractAddress  string    `bson:"contract_address"`
	Admin            string    `bson:"admin"`
	Creator          string    `bson:"creator"`
	ExecutedCount    int       `bson:"executed_count"`
	InstantiatedAt   time.Time `bson:"instantiated_at"`
	Label            string    `bson:"label"`
	LastExecutedAt   time.Time `bson:"last_executed_at"`
	Permission       string    `bson:"permission"`
	PermittedAddress string    `bson:"permitted_address"`
	TxHash           string    `bson:"txhash"`
	Version          string    `bson:"version"`
}

// NewCode returns a new Code.
func NewContract() *Contract {
	return &Contract{}
}

func (c *Contract) SetCode(codeId int) *Contract {
	c.CodeId = codeId
	return c
}

func (c *Contract) SetContract(contract string) *Contract {
	c.Contract = contract
	return c
}

func (c *Contract) SetContractAddress(contractAddress string) *Contract {
	c.ContractAddress = contractAddress
	return c
}

func (c *Contract) SetAdmin(admin string) *Contract {
	c.Admin = admin
	return c
}

func (c *Contract) SetCreator(creator string) *Contract {
	c.Creator = creator
	return c
}

func (c *Contract) SetExecutedCount(count int) *Contract {
	c.ExecutedCount = count
	return c
}

func (c *Contract) SetInstantiatedAt(instantiatedAt time.Time) *Contract {
	c.InstantiatedAt = instantiatedAt
	return c
}

func (c *Contract) SetLabel(label string) *Contract {
	c.Label = label
	return c
}

func (c *Contract) SetLastExecutedAt(lastExecutedAt time.Time) *Contract {
	c.LastExecutedAt = lastExecutedAt
	return c
}

func (c *Contract) SetPermission(permission string) *Contract {
	c.Permission = permission
	return c
}

func (c *Contract) SetPermittedAddress(permittedAddress string) *Contract {
	c.PermittedAddress = permittedAddress
	return c
}

func (contract *Contract) SetTxhash(txHash string) *Contract {
	contract.TxHash = txHash
	return contract
}

func (c *Contract) SetVersion(version string) *Contract {
	c.Version = version
	return c
}
