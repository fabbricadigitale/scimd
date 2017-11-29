package filter

import (
	"bufio"
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
		f, err := CompileString(input)
		if err != nil {
			t.Log(err)
			t.Fail()
		} else {
			output := f.String()
			require.Equal(t, input, output)
		}
	}
}

func TestParseInvalid(t *testing.T) {

	inFile, _ := os.Open("testdata/wrong.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		input := scanner.Text()
		f, err := CompileString(input)
		require.NotNil(t, err)
		require.Nil(t, f)
	}
}
