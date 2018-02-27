package mongo

import (
	"regexp"
	"strings"
)

const (
	keyAllowedPattern = `[A-Za-z0-9()+,\-.:=@;$_!*']+`
)

var keyRegexp = regexp.MustCompile(`^` + keyAllowedPattern + `$`)

// list of old, new string pairs.
// old(s) MUST match keyAllowedPattern, new(s) MUST not.
// List cannot be odd
var replacerArgs = []string{
	"$", "§",
	".", "°",
}

var unreplacerArgs []string

func init() {
	l := len(replacerArgs)
	if l%2 == 0 {
		unreplacerArgs = make([]string, l)
		for i := 0; i < l; i += 2 {
			unreplacerArgs[i], unreplacerArgs[i+1] = replacerArgs[i+1], replacerArgs[i]
		}
	} else {
		panic("mongo replacerArgs: odd list count")
	}
}

// keyEscape returns and escaped k suitable to be used as BSON Field Name, it panics if k does not match the keyAllowedPattern defined by this package.
//
// Since Mongo 3.6, field names can contain dots (i.e. .) and dollar signs (i.e. $) but querying on such fields is not yet functional (no official solution).
// Thus, escaping is still required, furthermore, to be safe, keyEscape accept only a subset of charset allowed by mongo, so escaped chars can not be present within a key.
// However, the charset's subset is enough for the purpose of SCIM.
//
// References:
// - Mongodb restrictions on Fields Names ( https://docs.mongodb.com/manual/reference/limits/#Restrictions-on-Field-Names )
// - https://jira.mongodb.org/browse/DOCS-9311
func keyEscape(k string) string {
	if k == "" {
		return k
	}
	if !keyRegexp.MatchString(k) {
		panic("mongo: invalid key: " + k)
	}

	r := strings.NewReplacer(replacerArgs...)
	return r.Replace(k)
}

// keyUnescape unscapes
func keyUnescape(k string) string {
	r := strings.NewReplacer(unreplacerArgs...)
	return r.Replace(k)
}

// escapeAttribute transform an attribute with urn, if present, in a mongodb format (with point to individuate sub-attributes)
func escapeAttribute(k string) string {
	index := strings.LastIndex(k, ":")
	if index == -1 {
		return keyEscape(k)
	}
	escaped := []string{keyEscape(k[:index]), k[index+1:]}

	return strings.Join(escaped, ".")
}
