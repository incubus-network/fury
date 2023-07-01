package types

const (
	// ModuleName name that will be used throughout the module
	ModuleName = "furydist"

	// StoreKey Top level store key where all module items will be stored
	StoreKey = ModuleName

	// RouterKey Top level router key
	RouterKey = ModuleName

	// QuerierRoute Top level query string
	QuerierRoute = ModuleName

	// DefaultParamspace default name for parameter store
	DefaultParamspace = ModuleName

	// FuryDistMacc module account for furydist
	FuryDistMacc = ModuleName

	// Treasury
	FundModuleAccount = "fury-fund"
)

var (
	CurrentDistPeriodKey = []byte{0x00}
	PreviousBlockTimeKey = []byte{0x01}
)
