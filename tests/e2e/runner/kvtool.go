package runner

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type KvtoolRunnerConfig struct {
	FuryConfigTemplate string

	ImageTag   string
	IncludeIBC bool

	EnableAutomatedUpgrade  bool
	FuryUpgradeName         string
	FuryUpgradeHeight       int64
	FuryUpgradeBaseImageTag string

	SkipShutdown bool
}

// KvtoolRunner implements a NodeRunner that spins up local chains with kvtool.
// It has support for the following:
// - running a Fury node
// - optionally, running an IBC node with a channel opened to the Fury node
// - optionally, start the Fury node on one version and upgrade to another
type KvtoolRunner struct {
	config KvtoolRunnerConfig
}

var _ NodeRunner = &KvtoolRunner{}

// NewKvtoolRunner creates a new KvtoolRunner.
func NewKvtoolRunner(config KvtoolRunnerConfig) *KvtoolRunner {
	return &KvtoolRunner{
		config: config,
	}
}

// StartChains implements NodeRunner.
// For KvtoolRunner, it sets up, runs, and connects to a local chain via kvtool.
func (k *KvtoolRunner) StartChains() Chains {
	// install kvtool if not already installed
	installKvtoolCmd := exec.Command("./scripts/install-kvtool.sh")
	installKvtoolCmd.Stdout = os.Stdout
	installKvtoolCmd.Stderr = os.Stderr
	if err := installKvtoolCmd.Run(); err != nil {
		panic(fmt.Sprintf("failed to install kvtool: %s", err.Error()))
	}

	// start local test network with kvtool
	log.Println("starting fury node")
	kvtoolArgs := []string{"testnet", "bootstrap", "--fury.configTemplate", k.config.FuryConfigTemplate}
	// include an ibc chain if desired
	if k.config.IncludeIBC {
		kvtoolArgs = append(kvtoolArgs, "--ibc")
	}
	// handle automated upgrade functionality, if defined
	if k.config.EnableAutomatedUpgrade {
		kvtoolArgs = append(kvtoolArgs,
			"--upgrade-name", k.config.FuryUpgradeName,
			"--upgrade-height", fmt.Sprint(k.config.FuryUpgradeHeight),
			"--upgrade-base-image-tag", k.config.FuryUpgradeBaseImageTag,
		)
	}
	// start the chain
	startFuryCmd := exec.Command("kvtool", kvtoolArgs...)
	startFuryCmd.Env = os.Environ()
	startFuryCmd.Env = append(startFuryCmd.Env, fmt.Sprintf("FURY_TAG=%s", k.config.ImageTag))
	startFuryCmd.Stdout = os.Stdout
	startFuryCmd.Stderr = os.Stderr
	log.Println(startFuryCmd.String())
	if err := startFuryCmd.Run(); err != nil {
		panic(fmt.Sprintf("failed to start fury: %s", err.Error()))
	}

	// wait for chain to be live.
	// if an upgrade is defined, this waits for the upgrade to be completed.
	if err := waitForChainStart(kvtoolFuryChain); err != nil {
		k.Shutdown()
		panic(err)
	}
	log.Println("fury is started!")

	chains := NewChains()
	chains.Register("fury", &kvtoolFuryChain)
	if k.config.IncludeIBC {
		chains.Register("ibc", &kvtoolIbcChain)
	}
	return chains
}

// Shutdown implements NodeRunner.
// For KvtoolRunner, it shuts down the local kvtool network.
// To prevent shutting down the chain (eg. to preserve logs or examine post-test state)
// use the `SkipShutdown` option on the config.
func (k *KvtoolRunner) Shutdown() {
	if k.config.SkipShutdown {
		log.Printf("would shut down but SkipShutdown is true")
		return
	}
	log.Println("shutting down fury node")
	shutdownFuryCmd := exec.Command("kvtool", "testnet", "down")
	shutdownFuryCmd.Stdout = os.Stdout
	shutdownFuryCmd.Stderr = os.Stderr
	if err := shutdownFuryCmd.Run(); err != nil {
		panic(fmt.Sprintf("failed to shutdown kvtool: %s", err.Error()))
	}
}
