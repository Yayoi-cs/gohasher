package attack

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func genStringSha256(length int, prefix string) string {
	if length == 0 {
		sus = genSha256(prefix)
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
func genSha256(str string) bool {
	data := []byte(str)
	//hash := md5.Sum(data)
	sha := sha256.New()
	hash := sha.Sum(data)
	hashString := hex.EncodeToString(hash[:])
	if line != "" && hashString == line {
		fmt.Println("BruteForce Success!!")
		fmt.Println("Hash\t", line)
		fmt.Println("Plain\t", str)
		return true
	}
	return false
}

func BruteforceSha256(arg string) {
	line = arg
	for charLen := 1; ; charLen++ {
		if sus {
			return
		}
		fmt.Println("Length:", charLen)
		genStringSha256(charLen, "")
	}
}
