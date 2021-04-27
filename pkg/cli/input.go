/*
	Input Function.
*/

package cli

import (
	"bufio"
	"fmt"
	"os"
)

// Input Returns input from command-line.
// message Input message.
func Input(message string) string {
	fmt.Print(message)

	//! Don't use fmt.Scanln
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
