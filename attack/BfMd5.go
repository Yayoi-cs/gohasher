package attack

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func genStringMd5(length int, prefix string) string {
	if length == 0 {
		sus = genMd5(prefix)
		return prefix
	}
	for _, char := range charset {
		if sus {
			break
		}
		newPrefix := prefix + string(char)
		genStringMd5(length-1, newPrefix)
	}
	return ""
}
func genMd5(str string) bool {
	data := []byte(str)
	hash := md5.Sum(data)
	hashString := hex.EncodeToString(hash[:])
	if line != "" && hashString == line {
		fmt.Println("BruteForce Success!!")
		fmt.Println("Hash\t", line)
		fmt.Println("Plain\t", str)
		return true
	}
	return false
}

func BruteforceMd5(arg string) {
	line = arg
	for charLen := 1; ; charLen++ {
		if sus {
			return
		}
		fmt.Println("Length:", charLen)
		genStringMd5(charLen, "")
	}
}
