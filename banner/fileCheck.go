package Ascii

import (
	"errors"
	"os"
)

func FileCheck(fileName string) (string, error) {
	var fileSize int64

	file_info, err := os.Stat("./bannerFiles/" + fileName + ".txt")
	if err != nil {
		return fileName, err
	}
	fileSize = file_info.Size()

	if (fileName == "standard" && fileSize != 6623) ||
		(fileName == "thinkertoy" && fileSize != 4703) ||
		(fileName == "shadow" && fileSize != 7463) {
		return fileName, errors.New("the Banner file has been altered: ")
	}

	return fileName, nil
}
