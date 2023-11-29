package core

import (
	"fmt"
	"os"

	"github.com/shafiqaimanx/shellvenom/src/menu"
)

func SavePayloadToFile(selectedPayload string, fileName string) {
	err := os.WriteFile(fileName, []byte(selectedPayload), 0644)
	if err != nil {
		fmt.Printf("%s[-]%s Error writing to file: %v\n", menu.CRIMSON, menu.RESET, err)
		return
	}
	fmt.Printf("%s[-]%s Saved as: %s\n", menu.CHARTREUSE, menu.RESET, fileName)
}
