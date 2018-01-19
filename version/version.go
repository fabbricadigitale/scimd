package version

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

// GenerateVersion is an utility to create a weak validator to assign to the resource's version attribute.
// Refs [rfc7232](https://tools.ietf.org/html/rfc7232#section-2.3).
// TODO => Generate strong validator
func GenerateVersion(weak bool, args ...string) string {
	hash := sha1.New()
	for _, arg := range args {
		hash.Write([]byte(arg))
	}
	return fmt.Sprintf("W/\"%s\"", base64.StdEncoding.EncodeToString(hash.Sum(nil)))
}
