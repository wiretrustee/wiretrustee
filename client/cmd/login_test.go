package cmd

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/netbirdio/netbird/client/internal"
	"github.com/netbirdio/netbird/iface"
	"github.com/netbirdio/netbird/util"
)

func TestLogin(t *testing.T) {
	mgmAddr := startTestingServices(t)

	tempDir := t.TempDir()
	confPath := tempDir + "/config.json"
	mgmtURL := fmt.Sprintf("http://%s", mgmAddr)
	rootCmd.SetArgs([]string{
		"login",
		"--config",
		confPath,
		"--log-file",
		"console",
		"--setup-key",
		strings.ToUpper("a2c8e62b-38f5-4553-b31e-dd66c696cebb"),
		"--management-url",
		mgmtURL,
	})
	err := rootCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}

	// validate generated config
	actualConf := &internal.Config{}
	_, err = util.ReadJson(confPath, actualConf)
	if err != nil {
		t.Errorf("expected proper config file written, got broken %v", err)
	}

	if actualConf.ManagementURL.String() != mgmtURL {
		t.Errorf("expected management URL %s got %s", mgmtURL, actualConf.ManagementURL.String())
	}

	if actualConf.WgIface != iface.WgInterfaceDefault {
		t.Errorf("expected WgIfaceName %s got %s", iface.WgInterfaceDefault, actualConf.WgIface)
	}

	if len(actualConf.PrivateKey) == 0 {
		t.Errorf("expected non empty Private key, got empty")
	}
}

func TestIsLinuxRunningDesktop(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("skipping test on non-linux platform")
	}

	err := os.Setenv("XDG_FOO", "BAR")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		err := os.Unsetenv("XDG_FOO")
		if err != nil {
			t.Fatal(err)
		}
	})
	isDesktop := isLinuxRunningDesktop()
	if !isDesktop {
		t.Errorf("expected desktop environment, got false")
	}
}
