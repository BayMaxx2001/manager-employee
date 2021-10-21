package scan

import (
	"bufio"
	"fmt"
	"os"
)

func ScannerString() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	return input
}

func ScannerNumber() int {
	idString := ScannerString()
	idNum, err := ValidateId(idString)
	if err != nil {
		fmt.Println(err)
	}
	return idNum
}
