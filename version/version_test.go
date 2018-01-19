package version

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const id = "2819c223-7f76-453a-919d-ab1234567891"
const time = "2011-05-13 04:42:34 +0000 UTC"
const weakETag = `W/"J0yxl786FxJjKExNdOcFPmtneY8="`

func TestGenerateVersion(t *testing.T) {

	etag := GenerateVersion(true, id, time)
	require.Equal(t, weakETag, etag)
}
