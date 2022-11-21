package dns

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
)

const (
	fileGeneratedResolvConfContentHeader      = "# Generated by NetBird"
	fileGeneratedResolvConfSearchBeginContent = "search "
	fileGeneratedResolvConfContentFormat      = fileGeneratedResolvConfContentHeader +
		"\n# If needed you can restore the original file by copying back %s\n\nnameserver %s\n" +
		fileGeneratedResolvConfSearchBeginContent + "%s"
)
const (
	fileDefaultResolvConfBackupLocation = defaultResolvConfPath + ".original.netbird"
	fileMaxLineCharsLimit               = 256
	fileMaxNumberOfSearchDomains        = 6
)

var fileSearchLineBeginCharCount = len(fileGeneratedResolvConfSearchBeginContent)

type fileConfigurator struct {
	originalPerms os.FileMode
}

func newFileConfigurator() hostManager {
	return &fileConfigurator{}
}

func (f *fileConfigurator) applyDNSConfig(config hostDNSConfig) error {
	if !config.routeAll {
		return fmt.Errorf("unable to configure DNS for this peer using file manager without a Primary nameserver group")
	}
	backupFileExist := false
	_, err := os.Stat(fileDefaultResolvConfBackupLocation)
	if err == nil {
		backupFileExist = true
	}

	switch getOSDNSManagerType() {
	case fileManager, netbirdManager:
		if !backupFileExist {
			err = f.backup()
			if err != nil {
				return fmt.Errorf("unable to backup the resolv.conf file")
			}
			backupFileExist = true
		}
	default:
		// todo improve this and maybe restart DNS manager from scratch
		return fmt.Errorf("something happened and file manager is not your prefered host dns configurator, restart the agent")
	}

	var searchDomains string
	appendedDomains := 0
	for _, dConf := range config.domains {
		if dConf.matchOnly {
			continue
		}
		if appendedDomains >= fileMaxNumberOfSearchDomains {
			// lets log all skipped domains
			log.Infof("already appended %d domains to search list. Skipping append of %s domain", fileMaxNumberOfSearchDomains, dConf.domain)
			continue
		}
		if fileSearchLineBeginCharCount+len(searchDomains) > fileMaxLineCharsLimit {
			// lets log all skipped domains
			log.Infof("search list line is larger than %d characters. Skipping append of %s domain", fileMaxLineCharsLimit, dConf.domain)
			continue
		}

		searchDomains += " " + dConf.domain
		appendedDomains++
	}
	content := fmt.Sprintf(fileGeneratedResolvConfContentFormat, fileDefaultResolvConfBackupLocation, config.serverIP, searchDomains)
	err = writeDNSConfig(content, defaultResolvConfPath, f.originalPerms)
	if err != nil {
		f.restore()
		return err
	}
	log.Infof("created a NetBird managed %s file with your DNS settings", defaultResolvConfPath)
	return nil
}

func (f *fileConfigurator) restoreHostDNS() error {
	return f.restore()
}

func (f *fileConfigurator) backup() error {
	stats, err := os.Stat(defaultResolvConfPath)
	if err != nil {
		return fmt.Errorf("got an error while checking stats for %s file. Error: %s", defaultResolvConfPath, err)
	}

	f.originalPerms = stats.Mode()

	err = copyFile(defaultResolvConfPath, fileDefaultResolvConfBackupLocation)
	if err != nil {
		return fmt.Errorf("got error while backing up the %s file. Error: %s", defaultResolvConfPath, err)
	}
	return nil
}

func (f *fileConfigurator) restore() error {
	err := copyFile(fileDefaultResolvConfBackupLocation, defaultResolvConfPath)
	if err != nil {
		return fmt.Errorf("got error while restoring the %s file from %s. Error: %s", defaultResolvConfPath, fileDefaultResolvConfBackupLocation, err)
	}
	return nil
}

func writeDNSConfig(content, fileName string, permissions os.FileMode) error {
	log.Debugf("creating managed file %s", fileName)
	var buf bytes.Buffer
	buf.WriteString(content)
	err := os.WriteFile(fileName, buf.Bytes(), permissions)
	if err != nil {
		return fmt.Errorf("got an creating resolver file %s err: %s", fileName, err)
	}
	return nil
}

func copyFile(src, dest string) error {
	_, err := exec.Command("cp", src, dest).CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}
