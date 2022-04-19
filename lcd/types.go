package lcd

import (
	"fmt"
	"time"

	"github.com/cosmos-gaminghub/explorer-backend/utils"
)

const (
	UrlAccount                = "%s/bank/accounts/%s"
	UrlBankTokenStats         = "%s/bank/token-stats"
	UrlValidator              = "%s/stake/validators/%s"
	UrlValidators             = "%s/cosmos/staking/v1beta1/validators?pagination.offset=%d"
	UrlDelegationByVal        = "%s/stake/validators/%s/delegations"
	UrlDelegationsByDelegator = "%s/stake/delegators/%s/delegations"
	//UrlDelegationsFromValidatorByDelegator = "%s/stake/delegators/%s/validators/%s"
	UrlUnbondingDelegationByDelegator            = "%s/stake/delegators/%s/unbonding-delegations"
	UrlDelegationsByValidator                    = "%s/stake/validators/%s/delegations"
	UrlUnbondingDelegationByValidator            = "%s/stake/validators/%s/unbonding-delegations"
	UrlRedelegationsByValidator                  = "%s/stake/validators/%s/redelegations"
	UrlSignInfo                                  = "%s/slashing/validators/%s/signing-info"
	UrlNodeInfo                                  = "%s/node-info"
	UrlNodeVersion                               = "%s/node-version"
	UrlGenesis                                   = "%s/genesis"
	UrlWithdrawAddress                           = "%s/distribution/%s/withdraw-address"
	UrlBlockLatest                               = "%s/cosmos/base/tendermint/v1beta1/blocks/latest"
	UrlBlock                                     = "%s/cosmos/base/tendermint/v1beta1/blocks/%d"
	UrlValidatorSet                              = "%s/cosmos/base/tendermint/v1beta1/validatorsets/%d?pagination.offset=%d"
	UrlValidatorSetLatest                        = "%s/cosmos/base/tendermint/v1beta1/validatorsets/latest"
	UrlStakePool                                 = "%s/stake/pool"
	UrlBlocksResult                              = "%s/block-results/%d"
	UrlTxsTxHeight                               = "%s/cosmos/tx/v1beta1/txs?events=tx.height=%d"
	UrlModuleParam                               = "%s/cosmos/%s/v1beta1/params"
	UrlGovParam                                  = "%s/cosmos/gov/v1beta1/params/%s"
	UrlDistributionRewardsByValidatorAcc         = "%s/distribution/%s/rewards"
	UrlValidatorsSigningInfoByConsensuPublicKey  = "%s/slashing/validators/%s/signing-info"
	UrlDistributionWithdrawAddressByValidatorAcc = "%s/distribution/%s/withdraw-address"
	UrlAssetTokens                               = "%s/asset/tokens"
	UrlAssetGateways                             = "%s/asset/gateways"
	UrlValSigningInfo                            = "%s/cosmos/slashing/v1beta1/signing_infos/%s"

	UrlProposal         = "%s/cosmos/gov/v1beta1/proposals"
	UrlProposalDeposit  = "%s/cosmos/gov/v1beta1/proposals/%d/deposits"
	UrlProposalProposer = "%s/gov/proposals/%d/proposer"
	UrlProposalVoters   = "%s/cosmos/gov/v1beta1/proposals/%d/votes?pagination.offset=%d"
	UrlProposalTally    = "%s/cosmos/gov/v1beta1/proposals/%d/tally"

	UrlWasmContractState = "%s/wasm/contract/%s/state"
	UrlWasmCode          = "%s/wasm/code"
	UrlWasmCodeContracts = "%s/wasm/code/%d/contracts"
	UrlWasmContract      = "%s/wasm/contract/%s"

	DefaultValidatorSetLimit = 100
	DefaultValidatorLimit    = 200
)

type AccountVo struct {
	Address string   `json:"address"`
	Coins   []string `json:"coins"`
	//PublicKey struct {
	//	Type  string `json:"type"`
	//	Value string `json:"value"`
	//} `json:"public_key"`
	AccountNumber string `json:"account_number"`
	Sequence      string `json:"sequence"`
}

type Validator struct {
	OperatorAddress string `json:"operator_address"`
	ConsensusPubkey struct {
		Type string `json:"@type"`
		Key  string `json:"key"`
	} `json:"consensus_pubkey"`
	Jailed            bool        `json:"jailed"`
	Status            string      `json:"status"`
	Tokens            string      `json:"tokens"`
	DelegatorShares   string      `json:"delegator_shares"`
	Description       Description `json:"description"`
	UnbondingHeight   string      `json:"unbonding_height"`
	UnbondingTime     time.Time   `json:"unbonding_time"`
	Commission        Commission  `json:"commission"`
	MinSelfDelegation string      `json:"min_self_delegation"`
}

type ValidatorsResult struct {
	Validators []Validator `json:"validators"`
	Pagination Pagination  `json:"pagination"`
}

func (v Validator) String() string {
	return fmt.Sprintf(`
		OperatorAddress :%v
		ConsensusPubkey :%v
		Jailed          :%v
		Status          :%v
		Tokens          :%v
		DelegatorShares :%v
		Description     :%v
		UnbondingHeight :%v
		UnbondingTime   :%v
		Commission      :%v
		`, v.OperatorAddress, v.ConsensusPubkey, v.Jailed, v.Status, v.Tokens, v.DelegatorShares, v.Description, v.UnbondingHeight, v.UnbondingTime,
		v.Commission)
}

type Description struct {
	ImageUrl        string `json:"image_url"`
	Moniker         string `json:"moniker"`
	Identity        string `json:"identity"`
	Website         string `json:"website"`
	Details         string `json:"details"`
	SecurityContact string `json:"security_contact"`
}
type Commission struct {
	CommissionRate struct {
		Rate          string `json:"rate"`
		MaxRate       string `json:"max_rate"`
		MaxChangeRate string `json:"max_change_rate"`
	} `json:"commission_rates"`
	UpdateTime string `json:"update_time"`
}

type NodeInfoVo struct {
	ProtocolVersion struct {
		P2P   string `json:"p2p"`
		Block string `json:"block"`
		App   string `json:"app"`
	} `json:"protocol_version"`
	ID         string `json:"id"`
	ListenAddr string `json:"listen_addr"`
	Network    string `json:"network"`
	Version    string `json:"version"`
	Channels   string `json:"channels"`
	Moniker    string `json:"moniker"`
	Other      struct {
		TxIndex    string `json:"tx_index"`
		RPCAddress string `json:"rpc_address"`
	} `json:"other"`
}
type GenesisVo struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		Genesis struct {
			GenesisTime     time.Time `json:"genesis_time"`
			ChainID         string    `json:"chain_id"`
			ConsensusParams struct {
				BlockSize struct {
					MaxBytes string `json:"max_bytes"`
					MaxGas   string `json:"max_gas"`
				} `json:"block_size"`
				Evidence struct {
					MaxAge string `json:"max_age"`
				} `json:"evidence"`
				Validator struct {
					PubKeyTypes []string `json:"pub_key_types"`
				} `json:"validator"`
			} `json:"consensus_params"`
			AppHash  string `json:"app_hash"`
			AppState struct {
				Accounts []struct {
					Address        string   `json:"address"`
					Coins          []string `json:"coins"`
					SequenceNumber string   `json:"sequence_number"`
					AccountNumber  string   `json:"account_number"`
				} `json:"accounts"`
				Auth struct {
					CollectedFee interface{} `json:"collected_fee"`
					Data         struct {
						NativeFeeDenom string `json:"native_fee_denom"`
					} `json:"data"`
					Params struct {
						GasPriceThreshold string `json:"gas_price_threshold"`
						TxSize            string `json:"tx_size"`
					} `json:"params"`
				} `json:"auth"`
				Stake struct {
					Pool struct {
						BondedTokens string `json:"bonded_tokens"`
					} `json:"pool"`
					Params struct {
						UnbondingTime string `json:"unbonding_time"`
						MaxValidators int    `json:"max_validators"`
					} `json:"params"`
					LastTotalPower       string      `json:"last_total_power"`
					LastValidatorPowers  interface{} `json:"last_validator_powers"`
					Validators           interface{} `json:"validators"`
					Bonds                interface{} `json:"bonds"`
					UnbondingDelegations interface{} `json:"unbonding_delegations"`
					Redelegations        interface{} `json:"redelegations"`
					Exported             bool        `json:"exported"`
				} `json:"stake"`
				Mint struct {
					Minter struct {
						LastUpdate        time.Time `json:"last_update"`
						MintDenom         string    `json:"mint_denom"`
						InflationBasement string    `json:"inflation_basement"`
					} `json:"minter"`
					Params struct {
						Inflation string `json:"inflation"`
					} `json:"params"`
				} `json:"mint"`
				Distr struct {
					Params struct {
						CommunityTax        string `json:"community_tax"`
						BaseProposerReward  string `json:"base_proposer_reward"`
						BonusProposerReward string `json:"bonus_proposer_reward"`
					} `json:"params"`
					FeePool struct {
						ValAccum struct {
							UpdateHeight string `json:"update_height"`
							Accum        string `json:"accum"`
						} `json:"val_accum"`
						ValPool       interface{} `json:"val_pool"`
						CommunityPool interface{} `json:"community_pool"`
					} `json:"fee_pool"`
					ValidatorDistInfos     interface{} `json:"validator_dist_infos"`
					DelegatorDistInfos     interface{} `json:"delegator_dist_infos"`
					DelegatorWithdrawInfos interface{} `json:"delegator_withdraw_infos"`
					PreviousProposer       string      `json:"previous_proposer"`
				} `json:"distr"`
				Gov struct {
					Params struct {
						CriticalDepositPeriod string `json:"critical_deposit_period"`
						CriticalMinDeposit    []struct {
							Denom  string `json:"denom"`
							Amount string `json:"amount"`
						} `json:"critical_min_deposit"`
						CriticalVotingPeriod   string `json:"critical_voting_period"`
						CriticalMaxNum         string `json:"critical_max_num"`
						CriticalThreshold      string `json:"critical_threshold"`
						CriticalVeto           string `json:"critical_veto"`
						CriticalParticipation  string `json:"critical_participation"`
						CriticalPenalty        string `json:"critical_penalty"`
						ImportantDepositPeriod string `json:"important_deposit_period"`
						ImportantMinDeposit    []struct {
							Denom  string `json:"denom"`
							Amount string `json:"amount"`
						} `json:"important_min_deposit"`
						ImportantVotingPeriod  string `json:"important_voting_period"`
						ImportantMaxNum        string `json:"important_max_num"`
						ImportantThreshold     string `json:"important_threshold"`
						ImportantVeto          string `json:"important_veto"`
						ImportantParticipation string `json:"important_participation"`
						ImportantPenalty       string `json:"important_penalty"`
						NormalDepositPeriod    string `json:"normal_deposit_period"`
						NormalMinDeposit       []struct {
							Denom  string `json:"denom"`
							Amount string `json:"amount"`
						} `json:"normal_min_deposit"`
						NormalVotingPeriod  string `json:"normal_voting_period"`
						NormalMaxNum        string `json:"normal_max_num"`
						NormalThreshold     string `json:"normal_threshold"`
						NormalVeto          string `json:"normal_veto"`
						NormalParticipation string `json:"normal_participation"`
						NormalPenalty       string `json:"normal_penalty"`
						SystemHaltPeriod    string `json:"system_halt_period"`
					} `json:"params"`
				} `json:"gov"`
				Upgrade struct {
					GenesisVersion struct {
						UpgradeInfo struct {
							ProposalID string `json:"ProposalID"`
							Protocol   struct {
								Version   string `json:"version"`
								Software  string `json:"software"`
								Height    string `json:"height"`
								Threshold string `json:"threshold"`
							} `json:"Protocol"`
						} `json:"UpgradeInfo"`
						Success bool `json:"Success"`
					} `json:"GenesisVersion"`
				} `json:"upgrade"`
				Slashing struct {
					Params struct {
						MaxEvidenceAge          string `json:"max_evidence_age"`
						SignedBlocksWindow      string `json:"signed_blocks_window"`
						MinSignedPerWindow      string `json:"min_signed_per_window"`
						DoubleSignJailDuration  string `json:"double_sign_jail_duration"`
						DowntimeJailDuration    string `json:"downtime_jail_duration"`
						CensorshipJailDuration  string `json:"censorship_jail_duration"`
						SlashFractionDoubleSign string `json:"slash_fraction_double_sign"`
						SlashFractionDowntime   string `json:"slash_fraction_downtime"`
						SlashFractionCensorship string `json:"slash_fraction_censorship"`
					} `json:"params"`
					SigningInfos struct {
					} `json:"signing_infos"`
					MissedBlocks struct {
					} `json:"missed_blocks"`
					SlashingPeriods interface{} `json:"slashing_periods"`
				} `json:"slashing"`
				Service struct {
					Params struct {
						MaxRequestTimeout    string `json:"max_request_timeout"`
						MinDepositMultiple   string `json:"min_deposit_multiple"`
						ServiceFeeTax        string `json:"service_fee_tax"`
						SlashFraction        string `json:"slash_fraction"`
						ComplaintRetrospect  string `json:"complaint_retrospect"`
						ArbitrationTimeLimit string `json:"arbitration_time_limit"`
						TxSizeLimit          string `json:"tx_size_limit"`
					} `json:"params"`
				} `json:"service"`
				Guardian struct {
					Profilers []struct {
						Description string `json:"description"`
						Type        string `json:"type"`
						Address     string `json:"address"`
						AddedBy     string `json:"added_by"`
					} `json:"profilers"`
					Trustees []struct {
						Description string `json:"description"`
						Type        string `json:"type"`
						Address     string `json:"address"`
						AddedBy     string `json:"added_by"`
					} `json:"trustees"`
				} `json:"guardian"`
				Gentxs []struct {
					Type  string `json:"type"`
					Value struct {
						Msg []struct {
							Type  string `json:"type"`
							Value struct {
								Description struct {
									Moniker  string `json:"moniker"`
									Identity string `json:"identity"`
									Website  string `json:"website"`
									Details  string `json:"details"`
								} `json:"Description"`
								Commission struct {
									Rate          string `json:"rate"`
									MaxRate       string `json:"max_rate"`
									MaxChangeRate string `json:"max_change_rate"`
								} `json:"Commission"`
								DelegatorAddress string `json:"delegator_address"`
								ValidatorAddress string `json:"validator_address"`
								Pubkey           struct {
									Type  string `json:"type"`
									Value string `json:"value"`
								} `json:"pubkey"`
								Delegation struct {
									Denom  string `json:"denom"`
									Amount string `json:"amount"`
								} `json:"delegation"`
							} `json:"value"`
						} `json:"msg"`
						Fee struct {
							Amount interface{} `json:"amount"`
							Gas    string      `json:"gas"`
						} `json:"fee"`
						Signatures []struct {
							PubKey struct {
								Type  string `json:"type"`
								Value string `json:"value"`
							} `json:"pub_key"`
							Signature     string `json:"signature"`
							AccountNumber string `json:"account_number"`
							Sequence      string `json:"sequence"`
						} `json:"signatures"`
						Memo string `json:"memo"`
					} `json:"value"`
				} `json:"gentxs"`
			} `json:"app_state"`
		} `json:"genesis"`
	} `json:"result"`
}

type BlockResult struct {
	BlockId BlockId `json:"block_id"`
	Block   struct {
		Header struct {
			Version struct {
				Block string `json:"block"`
				App   string `json:"app"`
			} `json:"version"`
			ChainID            string    `json:"chain_id"`
			Height             string    `json:"height"`
			Time               time.Time `json:"time"`
			NumTxs             string    `json:"num_txs"`
			TotalTxs           string    `json:"total_txs"`
			LastBlockID        BlockId   `json:"last_block_id"`
			LastCommitHash     string    `json:"last_commit_hash"`
			DataHash           string    `json:"data_hash"`
			ValidatorsHash     string    `json:"validators_hash"`
			NextValidatorsHash string    `json:"next_validators_hash"`
			ConsensusHash      string    `json:"consensus_hash"`
			AppHash            string    `json:"app_hash"`
			LastResultsHash    string    `json:"last_results_hash"`
			EvidenceHash       string    `json:"evidence_hash"`
			ProposerAddress    string    `json:"proposer_address"`
		} `json:"header"`
		Data struct {
			Txs []string `json:"txs"`
		} `json:"data`
		LastCommit struct {
			Height     string  `json:"height"`
			Round      int64   `json:"round"`
			BlockID    BlockId `json:"block_id"`
			Signatures []struct {
				Timestamp        time.Time `json:"timestamp"`
				ValidatorAddress string    `json:"validator_address"`
				Signature        string    `json:"signature"`
			} `json:"signatures"`
		} `json:"last_commit"`
	} `json:"block"`
}

type BlockId struct {
	Hash  string `json:"hash"`
	Parts struct {
		Total string `json:"total"`
		Hash  string `json:"hash"`
	} `json:"parts"`
}

type ValidatorSet struct {
	BlockHeight string                    `json:"block_height"`
	Validators  []ValidatorOfValidatorSet `json:"validators"`
}

type ValidatorOfValidatorSet struct {
	ConsensusAddr string `json:"address"`
	PubKey        struct {
		Type string `json:"@type"`
		Key  string `json:"key"`
	} `json:"pub_key"`
	ProposerPriority string `json:"proposer_priority"`
	VotingPower      string `json:"voting_power"`
}

type StakePoolVo struct {
	LooseTokens  string `json:"loose_tokens"`
	BondedTokens string `json:"bonded_tokens"`
	TotalSupply  string `json:"total_supply"`
	BondedRatio  string `json:"bonded_ratio"`
}

type DelegationVo struct {
	DelegatorAddr string `json:"delegator_addr"`
	ValidatorAddr string `json:"validator_addr"`
	Shares        string `json:"shares"`
	Height        int64  `json:"height,string"`
}

type DelegationFromVal struct {
	Tokens          string `json:"tokens"`
	DelegatorShares string `json:"delegator_shares"`
	OperatorAddress string `json:"operator_address"`
	BondHeight      int64  `json:"bond_height,string"`
}

type ValidatorDelegations []DelegationVo

func (sort ValidatorDelegations) Len() int {
	return len(sort)
}
func (sort ValidatorDelegations) Swap(i, j int) {
	sort[i], sort[j] = sort[j], sort[i]
}
func (sort ValidatorDelegations) Less(i, j int) bool {
	return sort[i].Height > sort[j].Height
}

type DistributionRewards struct {
	Total       utils.CoinsAsStr         `json:"total"`
	Delegations []RewardsFromDelegations `json:"delegations"`
	Commission  utils.CoinsAsStr         `json:"commission"`
}

type RewardsFromDelegations struct {
	Validator string           `json:"validator"`
	Reward    utils.CoinsAsStr `json:"reward"`
}

type ValidatorSigningInfo struct {
	StartHeight       string `json:"start_height"`
	IndexOffset       string `json:"index_offset"`
	JailedUntil       string `json:"jailed_until"`
	MissedBlocksCount string `json:"missed_blocks_counter"`
}

type ReDelegations struct {
	DelegatorAddr    string `json:"delegator_addr"`
	ValidatorSrcAddr string `json:"validator_src_addr"`
	ValidatorDstAddr string `json:"validator_dst_addr"`
	CreationHeight   string `json:"creation_height"`
	MinTime          int64  `json:"min_time"`
	InitialBalance   string `json:"initial_balance"`
	Balance          string `json:"balance"`
	SharesSrc        string `json:"shares_src"`
	SharesDst        string `json:"shares_dst"`
}

type UnbondingDelegations struct {
	DelegatorAddr  string `json:"delegator_addr"`
	ValidatorAddr  string `json:"validator_addr"`
	InitialBalance string `json:"initial_balance"`
	Balance        string `json:"balance"`
	CreationHeight int64  `json:"creation_height,string"`
	MinTime        string `json:"min_time"`
}

func (un UnbondingDelegations) String() string {
	return fmt.Sprintf(`
		DelegatorAddr  :%v
		ValidatorAddr  :%v
		InitialBalance :%v
		Balance        :%v
		CreationHeight :%v
		MinTime        :%v
		`, un.DelegatorAddr, un.ValidatorAddr, un.InitialBalance, un.Balance, un.CreationHeight, un.MinTime)

}

func (d DelegationVo) String() string {
	return fmt.Sprintf(`
		DelegatorAddr :%v
		ValidatorAddr :%v
		Shares        :%v
		Height        :%v
		`, d.DelegatorAddr, d.ValidatorAddr, d.Shares, d.Height)
}

type SignInfoVo struct {
	StartHeight         string    `json:"start_height"`
	IndexOffset         string    `json:"index_offset"`
	JailedUntil         time.Time `json:"jailed_until"`
	MissedBlocksCounter string    `json:"missed_blocks_counter"`
}

type BlockResultVo struct {
	Height  string `json:"height"`
	Results struct {
		DeliverTx []struct {
			Code      int         `json:"code"`
			Data      interface{} `json:"data"`
			Log       string      `json:"log"`
			Info      string      `json:"info"`
			GasWanted string      `json:"gas_wanted"`
			GasUsed   string      `json:"gas_used"`
			Tags      []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"tags"`
		} `json:"deliver_tx"`
		EndBlock struct {
			ValidatorUpdates []struct {
				PubKey struct {
					Type string `json:"type"`
					Data string `json:"data"`
				} `json:"pub_key"`
				Power string `json:"power"`
			} `json:"validator_updates"`
			ConsensusParamUpdates interface{} `json:"consensus_param_updates"`
			Tags                  []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"tags"`
		} `json:"end_block"`
		BeginBlock struct {
			Tags []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"tags"`
		} `json:"begin_block"`
	} `json:"results"`
}

type BlockCoinFlowVo struct {
	Height   string   `json:"height"`
	CoinFlow []string `json:"coin_flow"`
	Tx       Tx       `json:"tx"`
}

type Tx struct {
	Type string `json:"@type"`
	Body struct {
		Messages interface{} `json:"messages"`
		Memo     string      `json:"memo"`
	} `json:"body"`
	AuthInfo   TxAuthInfo `json:"auth_info"`
	Signatures []string   `json:"signatures"`
}

type TxMessages struct {
	Type        string `json:"@type"`
	FromAddress string `json:"from_address"`
	ToAddress   string `json:"to_address"`
	Amount      []struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	} `json:"amount"`
}

type TxAuthInfo struct {
	SignerInfos []TxSignerInfo `json:"signer_infos"`
	FeeInfo     Fee            `json:"fee"`
}

type Fee struct {
	Amount []struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	} `json:"amount"`
	GasLimit string `json:"gas_limit"`
	Granter  string `json:"granter"`
	Payer    string `json:"payer"`
}

type TxSignerInfo struct {
	TxSignerPubKey struct {
		Type string `json:"@type"`
		Key  string `json:"key"`
	} `json:"public_key"`
	ModeInfo struct {
		Single struct {
			Mode string `json:"mode"`
		} `json:"single"`
	} `json:"mode_info"`
	Sequence string `json:"sequence"`
}
type AssetTokens struct {
	BaseToken struct {
		Id              string `json:"id"`
		Family          string `json:"family"`
		Source          string `json:"source"`
		Gateway         string `json:"gateway"`
		Symbol          string `json:"symbol"`
		Name            string `json:"name"`
		Decimal         int    `json:"decimal"`
		CanonicalSymbol string `json:"canonical_symbol"`
		MinUnitAlias    string `json:"min_unit_alias"`
		InitialSupply   string `json:"initial_supply"`
		MaxSupply       string `json:"max_supply"`
		Mintable        bool   `json:"mintable"`
		Owner           string `json:"owner"`
	} `json:"base_token"`
}

type AssetGateways struct {
	Owner    string `json:"owner"`
	Moniker  string `json:"moniker"`
	Identity string `json:"identity"`
	Details  string `json:"details"`
	Website  string `json:"website"`
}

type TxResult struct {
	TxResponse []struct {
		Height    string `json:"height"`
		TxHash    string `json:"txhash"`
		Code      uint32 `json:"code"`
		RawLog    string `json:"raw_log"`
		Logs      []Log  `json:"logs"`
		Info      string `json:"info"`
		GasUsed   string `json:"gas_used"`
		GasWanted string `json:"gas_wanted"`
		Tx        Tx     `json:"tx"`
		Time      string `json:"timestamp"`
	} `json:"tx_responses"`
	Txs []struct {
		AuthInfo TxAuthInfo `json:"auth_info"`
		Body     struct {
			BodyMessage []interface{} `json:"messages"`
			Memo        string        `json:"memo"`
		} `json:"body"`
		Signatures []string `json:"signatures"`
	} `json:"txs"`
}

type BodyMessage struct {
	Type        string `json:"@type"`
	FromAddress string `bson:"from_address" json:"from_address"`
	ToAddress   string `bson:"to_address" json:"to_address"`
	Amount      []struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	} `bson:"amount" json:"amount"`
}

type Log struct {
	MsgIndex int     `json:"msg_index"`
	Log      string  `json:"log"`
	Events   []Event `json:"events"`
}

type Event struct {
	Type       string           `json:"type"`
	Attributes []EventAttribute `json:"attributes"`
}

type EventAttribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type DistributionParam struct {
	Params struct {
		CommunityTax        string `json:"community_tax"`
		BaseProposerReward  string `json:"base_proposer_reward"`
		BonusProposerReward string `json:"bonus_proposer_reward"`
		WithdrawAddrEnabled bool   `json:"withdraw_addr_enabled"`
	} `json:"params"`
}

type BankParam struct {
	Params struct {
		SendEnabled []struct {
			Denom   string `json:"denom"`
			Enabled bool   `json:"enabled"`
		} `json:"send_enabled"`
		DefaultSendEnabled bool `json:"default_send_enabled"`
	} `json:"params"`
}

type AuthParam struct {
	Params struct {
		MaxMemoCharacters      string `json:"max_memo_characters"`
		TxSigLimit             string `json:"tx_sig_limit"`
		TxSizeCostPerByte      string `json:"tx_size_cost_per_byte"`
		SigVerifyCostEd25591   string `json:"sig_verify_cost_ed25519"`
		SigVerifyCostSecp256k1 string `json:"sig_verify_cost_secp256k1"`
	} `json:"params"`
}

type GovParam struct {
	VotingParams struct {
		VotingPeriod string `json:"voting_period"`
	} `json:"voting_params"`
	DepositParams struct {
		MinDeposit []struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"min_deposit"`
		MaxDeposiyPeriod string `json:"max_deposit_period"`
	}
	TallyParams struct {
		Quorum        string `json:"quorum"`
		Threshold     string `json:"threshold"`
		VetoThreshold string `json:"veto_threshold"`
	} `json:"tally_params"`
}

type ProposalResult struct {
	Proposals []Proposal `json:"proposals"`
}

type Proposal struct {
	ProposalId       string                   `json:"proposal_id"`
	Content          ProposalContent          `json:"content"`
	Status           string                   `json:"status"`
	FinalTallyResult ProposalFinalTallyResult `json:"final_tally_result"`
	SubmitTime       time.Time                `json:"submit_time"`
	DepositEndTime   time.Time                `json:"deposit_end_time"`
	TotalDeposit     []struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	} `json:"total_deposit"`
	VotingStartTime time.Time `json:"voting_start_time"`
	VotingEndTime   time.Time `json:"voting_end_time"`
}

type ProposalContent struct {
	Type        string `bson:"type" json:"@type"`
	Title       string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	Amount      []struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	} `bson:"amount" json:"amount"`
	Changes []struct {
		Key      string `json:"key"`
		Value    string `json:"value"`
		Subspace string `json:"subspace"`
	} `json:"changes"`
	Plan struct {
		Name                string    `json:"name"`
		Time                time.Time `json:"time"`
		Height              string    `json:"height"`
		Info                string    `json:"info"`
		UpgradedClientState string    `json:"upgraded_client_state"`
	} `json:"plan"`
}

type ProposalFinalTallyResult struct {
	Yes        string `bson:"yes" json:"yes"`
	Abstain    string `bson:"abstain" json:"abstain"`
	No         string `bson:"no" json:"no"`
	NoWithVeto string `bson:"no_with_veto" json:"no_with_veto"`
}

type ProposalDepositResult struct {
	Deposits []ProposalDeposit `json:"deposits"`
}

type ProposalDeposit struct {
	ProposalId string                  `bson:"proposal_id" json:"proposal_id"`
	Depositor  string                  `json:"depositor"`
	Amount     []ProposalDepositAmount `json:"amount"`
}

type ProposalDepositAmount struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type ProposalVoteResult struct {
	Votes      []ProposalVote `json:"votes"`
	Pagination Pagination     `json:"pagination"`
}

type Pagination struct {
	NextKey string `json:"next_key"`
	Total   string `json:"total"`
}

type ProposalVote struct {
	ProposalId string `json:"proposal_id"`
	Voter      string `json:"voter"`
	Option     string `json:"option"`
}

type ProposalProposerResult struct {
	Result struct {
		ProposalId string `json:"proposal_id"`
		Proposer   string `json:"proposer"`
	} `json:"result"`
}

type ValSigningInfo struct {
	Info struct {
		Address           string `json:"address"`
		StartHeight       string `json:"start_height"`
		IndexOffset       string `json:"index_offset"`
		MissedBlocksCount string `json:"missed_blocks_counter"`
	} `json:"val_signing_info"`
}

type PropsalTally struct {
	Tally ProposalFinalTallyResult `json:"tally"`
}

type ContractState struct {
	Height int `json:"height"`
	Result []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
}

type WasmCode struct {
	Height string `json:"height"`
	Result []struct {
		Id       int    `json:"id"`
		Creator  string `json:"creator"`
		DataHash string `json:"data_hash"`
	}
}

type WasmCodeContracts struct {
	Height          string   `json:"height"`
	ContractAddress []string `json:"result"`
}

type WasmContract struct {
	Height string `json:"height"`
	Result struct {
		CodeId          int    `json:"code_id"`
		ContractAddress string `json:"address"`
		Creator         string `json:"creator"`
		Label           string `json:"label"`
	}
}
