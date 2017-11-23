package filter

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {

	inFile, _ := os.Open("testdata/ok.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		input := scanner.Text()
		f, _ := CompileString(input)
		output := f.String()
		require.Equal(t, input, output)
		fmt.Println(output)
	}

	// TODO
}
