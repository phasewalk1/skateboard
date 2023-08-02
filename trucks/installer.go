package trucks

import (
	"os"
	"os/exec"

	"github.com/phasewalk1/skateboard/util"

	"github.com/charmbracelet/log"
)

func BootstrapTrucks(home string) error {
	err := os.Chdir(home)
	if err != nil {
		return err
	}

	err = checkoutAndBuildFennel()
	if err != nil {
		log.Fatal("error bootstrapping Fennel:", err)
		return err
	}

	err = mvFennelLibToHome(home)
	if err != nil {
		log.Fatal("error copying sources into home:", err)
		return err
	}

	err = installTrucks()
	if err != nil {
		log.Fatal("error installing trucks:", err)
		return err
	}

	return nil
}

func checkoutAndBuildFennel() error {
	fnlSrcRemote := "https://git.sr.ht/~technomancy/fennel"
	log.Debug("checking out Fennel source...")
	log.Debug("remote:", fnlSrcRemote)

	cloneFnlSrc := exec.Command("git", "clone", fnlSrcRemote)
	util.ExecWithFatal(cloneFnlSrc, "Error cloning Fennel source:")

	err := os.Chdir("fennel")
	if err != nil {
		return err
	}

	makeObjects := exec.Command("make")
	util.ExecWithFatal(makeObjects, "Error building Fennel:")
	return nil
}

func mvFennelLibToHome(home string) error {
	err := os.Chdir(home)
	if err != nil {
		return err
	}

	err = os.Mkdir("include", 0755)
	if err != nil {
		log.Warn("couldn't create include directory:", err)
		log.Warn("creating include directory in current directory")
		cmd := exec.Command("mkdir", "include")
		_, err = cmd.CombinedOutput()
		if err != nil {
			log.Fatal("Error creating include directory:", err)
			return err
		}
	}

	cpLuaLibrary := exec.Command("cp", "fennel/fennel.lua", "include")
	util.ExecWithFatal(cpLuaLibrary, "Error copying Fennel library:")

	rmUnwanted := exec.Command("rm", "-rf", "fennel")
	util.ExecWithFatal(rmUnwanted, "Error removing Fennel source:")

	return nil
}

func installTrucks() error {
	sk8 := "https://github.com/phasewalk1/skateboard"

	cloneSkateboard := exec.Command("git", "clone", "--branch", "master", sk8)
	util.ExecWithFatal(cloneSkateboard, "Error cloning skateboard:")

	log.Debug("copying trucks into $HOME/.skateboard")

	getTrucks := exec.Command("cp", "-r", "skateboard/trucks", ".")
	util.ExecWithFatal(getTrucks, "Error copying trucks:")

	log.Debug("removing skateboard directory")

	rmUnwanted := exec.Command("rm", "-rf", "skateboard")
	util.ExecWithFatal(rmUnwanted, "Error removing skateboard directory:")

	return nil
}
