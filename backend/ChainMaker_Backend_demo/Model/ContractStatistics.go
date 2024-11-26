package Model

type ContractStatistics struct {
	Id               int64
	ContractName     string
	ContractVersion  string
	ContractOperator string
	TxNum            int
	Timestamp        int64
	Addr             string
	Sender           string
}
