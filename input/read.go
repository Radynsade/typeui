package input

import (
	"bufio"
	"os"
	"strings"
)

func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	// convert CRLF to LF
	value = strings.Replace(value, "\n", "", -1)

	return value
}
