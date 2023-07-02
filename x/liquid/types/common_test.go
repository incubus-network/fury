package types_test

import (
	"os"
	"testing"

	"github.com/incubus-network/fury/app"
)

func TestMain(m *testing.M) {
	app.SetSDKConfig()
	os.Exit(m.Run())
}
