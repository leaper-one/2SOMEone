package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func CaculateMD5(path string) (string,error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Open", err)
		return "", err
	}

	defer f.Close()

	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		fmt.Println("Copy", err)
		return "",err
	}

	//fmt.Printf("%x\n", md5hash.Sum(nil))
	return hex.EncodeToString(md5hash.Sum(nil)), nil
}
