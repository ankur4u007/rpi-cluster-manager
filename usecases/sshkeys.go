package usecases

import (
	"fmt"
	"io/ioutil"

	"github.com/ankur4u007/dietpi-image-flasher/entities/domain"
	"github.com/ankur4u007/dietpi-image-flasher/entities/utilities"
)

func ConfigureSshKeys() map[string]bool {
	configs := make(map[string]bool)
	if domain.Config.Boot.SshKeys.Enabled {
		if utilities.IsValueNonEmpty(domain.Config.Boot.SshKeys.PublicKeyPath) {
			sshPubKey, err := ioutil.ReadFile(domain.Config.Boot.SshKeys.PublicKeyPath)
			if err == nil {
				sshPubKeyConfig := fmt.Sprintf("AUTO_SETUP_SSH_PUBKEY=%s", string(sshPubKey))
				configs[sshPubKeyConfig] = true
			}
		}
		if domain.Config.Boot.SshKeys.DisablePasswordLogins {
			configs["SOFTWARE_DISABLE_SSH_PASSWORD_LOGINS=1"] = true
		}
	} else {
		fmt.Println("Skipping ssh keys setup")
	}
	return configs
}
