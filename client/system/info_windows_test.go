package system

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func Test_sysInfo(t *testing.T) {
	sysInfo()
	serialNum, prodName, manufacturer := sysInfo()
	if serialNum == "" {
		t.Errorf("serialNum is empty")
	}

	if prodName == "" {
		t.Errorf("prodName is empty")
	}

	if manufacturer == "" {
		t.Errorf("manufacturer is empty")
	}
	log.Infof("Windows sys info: %s, %s, %s", serialNum, prodName, manufacturer)
}
