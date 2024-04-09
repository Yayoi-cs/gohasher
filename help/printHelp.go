package help

import "fmt"

func PrintHelp() {
	fmt.Println("gohasher v1.00")
	fmt.Println("Usage: gohasher <option> <filename>")
	fmt.Println("options:")
	fmt.Println("\t-md5\t\tBruteforce md5 hash")
	fmt.Println("\t-sha256\t\tBruteforce sha256 hash")

}
