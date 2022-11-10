package tanda

import (
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

func IsSnapAndNotHomeDir() bool {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	snapRealHome := os.Getenv("SNAP_REAL_HOME")

	log.Debug().Msgf("%v %v", "currentDir", currentDir)
	log.Debug().Msgf("%v %v", "homeDir", homeDir)
	log.Debug().Msgf("%v %v", "snapRealHome", snapRealHome)

	isSnapPackage := strings.Contains(homeDir, "snap/tanda")
	isNotHomeDirForSnap := !strings.Contains(currentDir, snapRealHome)

	log.Debug().Msgf("is snap %v", isSnapPackage)
	log.Debug().Msgf("is not home dir %v", isNotHomeDirForSnap)

	// Snap package confirm
	if isSnapPackage && isNotHomeDirForSnap {
		return true
	}

	return false
}
