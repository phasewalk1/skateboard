package trucks

import (
	"os"
	"os/exec"

	"github.com/phasewalk1/skateboard/util"

	"github.com/charmbracelet/log"
)

func BootstrapTrucks(home string, noDeps bool) error {
	err := os.Chdir(home)
	if err != nil {
		return err
	}

	if !noDeps {
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

    scope := "trucks.install.fennel"
	log.Debug(scope, "checking out Fennel source...")
	log.Debug(scope, "remote:", fnlSrcRemote)

	cloneFnlSrc := exec.Command("git", "clone", fnlSrcRemote)
	util.ExecWithFatal(cloneFnlSrc, scope, "Error cloning Fennel source:")

	err := os.Chdir("fennel")
	if err != nil {
		return err
	}

	makeObjects := exec.Command("make")
	util.ExecWithFatal(makeObjects, scope, "Error building Fennel:")
	return nil
}

func mvFennelLibToHome(home string) error {
	err := os.Chdir(home)
	if err != nil {
		return err
	}

    scope := "trucks.install.fennel"

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
	util.ExecWithFatal(cpLuaLibrary, scope, "Error copying Fennel library:")

	rmUnwanted := exec.Command("rm", "-rf", "fennel")
	util.ExecWithFatal(rmUnwanted, scope, "Error removing Fennel source:")

	return nil
}

func installTrucks() error {
	sk8 := "https://github.com/phasewalk1/skateboard"

    scope := "trucks.install.trucks"

	cloneSkateboard := exec.Command("git", "clone", "--branch", "master", sk8)
	util.ExecWithFatal(cloneSkateboard, scope, "Error cloning skateboard:")

	log.Debug("copying trucks into $HOME/.skateboard")

	err := os.Chdir("skateboard")
	if err != nil {
		return err
	}
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	log.Debug("pwd:", pwd)
	runMake := exec.Command("make")
	util.ExecWithFatal(runMake, scope, "Error building trucks:")

	cpObjects := exec.Command("sh", "-c", "cp -r include/*.lua ../include")
	util.ExecWithFatal(cpObjects, scope, "Error copying trucks objects:")

	getTrucks := exec.Command("cp", "-r", "trucks", "..")
	util.ExecWithFatal(getTrucks, scope ,"Error copying trucks:")

    rmUnwanted := exec.Command("rm", "../trucks/installer.go")
    util.ExecWithFatal(rmUnwanted, scope, "Error pruning trucks:")

	err = os.Chdir("..")
	if err != nil {
		return err
	}
	log.Debug("removing skateboard directory")

	rmUnwanted = exec.Command("rm", "-rf", "skateboard")
	util.ExecWithFatal(rmUnwanted, scope, "Error removing skateboard directory:")

	return nil
}
