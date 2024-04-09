package main

import (
	"fmt"
	"gohasher/attack"
	"gohasher/help"
	"gohasher/ioHelper"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		help.PrintHelp()
		os.Exit(1)
	}

	line, err := ioHelper.ReadHashFile(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Hash:", line)

	switch args[1] {
	case "-md5":
		fmt.Println("option: md5")
		attack.BruteforceMd5(line)
	case "-sha256":
		fmt.Println("option: sha256")
		attack.BruteforceSha256(line)
	default:
		help.PrintHelp()
	}
	os.Exit(0)

}
