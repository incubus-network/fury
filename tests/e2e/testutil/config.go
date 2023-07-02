package testutil

import (
	"fmt"
	"os"
	"strconv"

	"github.com/subosito/gotenv"
)

func init() {
	// read the .env file, if present
	gotenv.Load()
}

// SuiteConfig wraps configuration details for running the end-to-end test suite.
type SuiteConfig struct {
	// A funded account used to fnd all other accounts.
	FundedAccountMnemonic string

	// A config for using futool local networks for the test run
	Kvtool *KvtoolConfig
	// A config for connecting to a running network
	LiveNetwork *LiveNetworkConfig

	// Whether or not to start an IBC chain. Use `suite.SkipIfIbcDisabled()` in IBC tests in IBC tests.
	IncludeIbcTests bool

	// The contract address of a deployed ERC-20 token
	FuryErc20Address string

	// When true, the chains will remain running after tests complete (pass or fail)
	SkipShutdown bool
}

// KvtoolConfig wraps configuration options for running the end-to-end test suite against
// a locally running chain. This config must be defined if E2E_RUN_KVTOOL_NETWORKS is true.
type KvtoolConfig struct {
	// The fury.configTemplate flag to be passed to futool, usually "master".
	// This allows one to change the base genesis used to start the chain.
	FuryConfigTemplate string

	// Whether or not to run a chain upgrade & run post-upgrade tests. Use `suite.SkipIfUpgradeDisabled()` in post-upgrade tests.
	IncludeAutomatedUpgrade bool
	// Name of the upgrade, if upgrade is enabled.
	FuryUpgradeName string
	// Height upgrade will be applied to the test chain, if upgrade is enabled.
	FuryUpgradeHeight int64
	// Tag of fury docker image that will be upgraded to the current image before tests are run, if upgrade is enabled.
	FuryUpgradeBaseImageTag string
}

// LiveNetworkConfig wraps configuration options for running the end-to-end test suite
// against a live network. It must be defined if E2E_RUN_KVTOOL_NETWORKS is false.
type LiveNetworkConfig struct {
	FuryRpcUrl    string
	FuryGrpcUrl   string
	FuryEvmRpcUrl string
}

// ParseSuiteConfig builds a SuiteConfig from environment variables.
func ParseSuiteConfig() SuiteConfig {
	config := SuiteConfig{
		// this mnemonic is expected to be a funded account that can seed the funds for all
		// new accounts created during tests. it will be available under Accounts["whale"]
		FundedAccountMnemonic: nonemptyStringEnv("E2E_FURY_FUNDED_ACCOUNT_MNEMONIC"),
		FuryErc20Address:      nonemptyStringEnv("E2E_FURY_ERC20_ADDRESS"),
		IncludeIbcTests:       mustParseBool("E2E_INCLUDE_IBC_TESTS"),
	}

	skipShutdownEnv := os.Getenv("E2E_SKIP_SHUTDOWN")
	if skipShutdownEnv != "" {
		config.SkipShutdown = mustParseBool("E2E_SKIP_SHUTDOWN")
	}

	useKvtoolNetworks := mustParseBool("E2E_RUN_KVTOOL_NETWORKS")
	if useKvtoolNetworks {
		futoolConfig := ParseKvtoolConfig()
		config.Kvtool = &futoolConfig
	} else {
		liveNetworkConfig := ParseLiveNetworkConfig()
		config.LiveNetwork = &liveNetworkConfig
	}

	return config
}

// ParseKvtoolConfig builds a KvtoolConfig from environment variables.
func ParseKvtoolConfig() KvtoolConfig {
	config := KvtoolConfig{
		FuryConfigTemplate:      nonemptyStringEnv("E2E_KVTOOL_FURY_CONFIG_TEMPLATE"),
		IncludeAutomatedUpgrade: mustParseBool("E2E_INCLUDE_AUTOMATED_UPGRADE"),
	}

	if config.IncludeAutomatedUpgrade {
		config.FuryUpgradeName = nonemptyStringEnv("E2E_FURY_UPGRADE_NAME")
		config.FuryUpgradeBaseImageTag = nonemptyStringEnv("E2E_FURY_UPGRADE_BASE_IMAGE_TAG")
		upgradeHeight, err := strconv.ParseInt(nonemptyStringEnv("E2E_FURY_UPGRADE_HEIGHT"), 10, 64)
		if err != nil {
			panic(fmt.Sprintf("E2E_FURY_UPGRADE_HEIGHT must be a number: %s", err))
		}
		config.FuryUpgradeHeight = upgradeHeight
	}

	return config
}

// ParseLiveNetworkConfig builds a LiveNetworkConfig from environment variables.
func ParseLiveNetworkConfig() LiveNetworkConfig {
	return LiveNetworkConfig{
		FuryRpcUrl:    nonemptyStringEnv("E2E_FURY_RPC_URL"),
		FuryGrpcUrl:   nonemptyStringEnv("E2E_FURY_GRPC_URL"),
		FuryEvmRpcUrl: nonemptyStringEnv("E2E_FURY_EVM_RPC_URL"),
	}
}

// mustParseBool is a helper method that panics if the env variable `name`
// cannot be parsed to a boolean
func mustParseBool(name string) bool {
	envValue := os.Getenv(name)
	if envValue == "" {
		panic(fmt.Sprintf("%s is unset but expected a bool", name))
	}
	value, err := strconv.ParseBool(envValue)
	if err != nil {
		panic(fmt.Sprintf("%s (%s) cannot be parsed to a bool: %s", name, envValue, err))
	}
	return value
}

// nonemptyStringEnv is a helper method that panics if the env variable `name`
// is empty or undefined.
func nonemptyStringEnv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		panic(fmt.Sprintf("no %s env variable provided", name))
	}
	return value
}
