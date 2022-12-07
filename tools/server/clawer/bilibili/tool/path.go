package tool

import (
	"os"
	"os/exec"
	"path/filepath"
)

func GetMp4Dir(title string) string {
	curDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fullDirPath := filepath.Join(curDir, "output", title)
	err = os.MkdirAll(fullDirPath, 0777)
	if err != nil {
		panic(err)
	}
	return fullDirPath
}

func FileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}

func CheckFfmegStatus() bool {
	_, err := exec.LookPath("ffmpeg")
	if err != nil {
		return false
	} else {
		return true
	}
}
