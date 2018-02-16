package attr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type projectionTestCase struct {
	included []string
	excluded []string
	expected map[string]bool
}

// NOTE: attributes with children (subattributes) are excluded
var allUserAttrs = map[string]bool{
	"id":         true,
	"externalId": true,
	// "meta":              true,
	"meta.created":      true,
	"meta.lastModified": true,
	"meta.location":     true,
	"meta.resourceType": true,
	"meta.version":      true,
	"schemas":           true,
	"urn:ietf:params:scim:schemas:core:2.0:User:userName": true,
	// "urn:ietf:params:scim:schemas:core:2.0:User:name":                                true,
	"urn:ietf:params:scim:schemas:core:2.0:User:name.formatted":       true,
	"urn:ietf:params:scim:schemas:core:2.0:User:name.familyName":      true,
	"urn:ietf:params:scim:schemas:core:2.0:User:name.givenName":       true,
	"urn:ietf:params:scim:schemas:core:2.0:User:name.middleName":      true,
	"urn:ietf:params:scim:schemas:core:2.0:User:name.honorificPrefix": true,
	"urn:ietf:params:scim:schemas:core:2.0:User:name.honorificSuffix": true,
	"urn:ietf:params:scim:schemas:core:2.0:User:displayName":          true,
	"urn:ietf:params:scim:schemas:core:2.0:User:nickName":             true,
	"urn:ietf:params:scim:schemas:core:2.0:User:profileUrl":           true,
	"urn:ietf:params:scim:schemas:core:2.0:User:title":                true,
	"urn:ietf:params:scim:schemas:core:2.0:User:userType":             true,
	"urn:ietf:params:scim:schemas:core:2.0:User:preferredLanguage":    true,
	"urn:ietf:params:scim:schemas:core:2.0:User:locale":               true,
	"urn:ietf:params:scim:schemas:core:2.0:User:timezone":             true,
	"urn:ietf:params:scim:schemas:core:2.0:User:active":               true,
	// "urn:ietf:params:scim:schemas:core:2.0:User:emails":                              true,
	"urn:ietf:params:scim:schemas:core:2.0:User:emails.value":   true,
	"urn:ietf:params:scim:schemas:core:2.0:User:emails.display": true,
	"urn:ietf:params:scim:schemas:core:2.0:User:emails.type":    true,
	"urn:ietf:params:scim:schemas:core:2.0:User:emails.primary": true,
	// "urn:ietf:params:scim:schemas:core:2.0:User:phoneNumbers":                        true,
	"urn:ietf:params:scim:schemas:core:2.0:User:phoneNumbers.value":   true,
	"urn:ietf:params:scim:schemas:core:2.0:User:phoneNumbers.display": true,
	"urn:ietf:params:scim:schemas:core:2.0:User:phoneNumbers.type":    true,
	"urn:ietf:params:scim:schemas:core:2.0:User:phoneNumbers.primary": true,
	// "urn:ietf:params:scim:schemas:core:2.0:User:ims":                                 true,
	"urn:ietf:params:scim:schemas:core:2.0:User:ims.value":   true,
	"urn:ietf:params:scim:schemas:core:2.0:User:ims.display": true,
	"urn:ietf:params:scim:schemas:core:2.0:User:ims.type":    true,
	"urn:ietf:params:scim:schemas:core:2.0:User:ims.primary": true,
	// "urn:ietf:params:scim:schemas:core:2.0:User:photos":                              true,
	"urn:ietf:params:scim:schemas:core:2.0:User:photos.value":   true,
	"urn:ietf:params:scim:schemas:core:2.0:User:photos.display": true,
	"urn:ietf:params:scim:schemas:core:2.0:User:photos.type":    true,
	"urn:ietf:params:scim:schemas:core:2.0:User:photos.primary": true,
	// "urn:ietf:params:scim:schemas:core:2.0:User:addresses":                           true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.formatted":     true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.streetAddress": true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.locality":      true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.region":        true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.postalCode":    true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.country":       true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.type":          true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.primary":       true,
	// "urn:ietf:params:scim:schemas:core:2.0:User:groups":                              true,
	"urn:ietf:params:scim:schemas:core:2.0:User:groups.value":   true,
	"urn:ietf:params:scim:schemas:core:2.0:User:groups.$ref":    true,
	"urn:ietf:params:scim:schemas:core:2.0:User:groups.display": true,
	"urn:ietf:params:scim:schemas:core:2.0:User:groups.type":    true,
	// "urn:ietf:params:scim:schemas:core:2.0:User:entitlements":                        true,
	"urn:ietf:params:scim:schemas:core:2.0:User:entitlements.value":   true,
	"urn:ietf:params:scim:schemas:core:2.0:User:entitlements.display": true,
	"urn:ietf:params:scim:schemas:core:2.0:User:entitlements.type":    true,
	"urn:ietf:params:scim:schemas:core:2.0:User:entitlements.primary": true,
	// "urn:ietf:params:scim:schemas:core:2.0:User:roles":                               true,
	"urn:ietf:params:scim:schemas:core:2.0:User:roles.value":   true,
	"urn:ietf:params:scim:schemas:core:2.0:User:roles.display": true,
	"urn:ietf:params:scim:schemas:core:2.0:User:roles.type":    true,
	"urn:ietf:params:scim:schemas:core:2.0:User:roles.primary": true,
	// "urn:ietf:params:scim:schemas:core:2.0:User:x509Certificates":                    true,
	"urn:ietf:params:scim:schemas:core:2.0:User:x509Certificates.value":         true,
	"urn:ietf:params:scim:schemas:core:2.0:User:x509Certificates.display":       true,
	"urn:ietf:params:scim:schemas:core:2.0:User:x509Certificates.type":          true,
	"urn:ietf:params:scim:schemas:core:2.0:User:x509Certificates.primary":       true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:costCenter":     true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:department":     true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:division":       true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:employeeNumber": true,
	// "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager":             true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager.$ref":        true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager.displayName": true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager.value":       true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:organization":        true,
}

var minimalUserAttributes = map[string]bool{
	"id":                true,
	"schemas":           true,
	"meta.resourceType": true,
}

var a = map[string]bool{
	"id":                                                     true,
	"schemas":                                                true,
	"meta.resourceType":                                      true,
	"urn:ietf:params:scim:schemas:core:2.0:User:displayName": true,
}

var b = map[string]bool{
	"id":                                                                             true,
	"schemas":                                                                        true,
	"meta.resourceType":                                                              true,
	"urn:ietf:params:scim:schemas:core:2.0:User:displayName":                         true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager.displayName": true,
}

var c = map[string]bool{
	"id":         true,
	"externalId": true,
	// "meta":              true,
	"meta.created":      true,
	"meta.lastModified": true,
	"meta.location":     true,
	"meta.resourceType": true,
	"meta.version":      true,
	"schemas":           true,
	"urn:ietf:params:scim:schemas:core:2.0:User:userName": true,
	// "urn:ietf:params:scim:schemas:core:2.0:User:name":                                true,
	"urn:ietf:params:scim:schemas:core:2.0:User:name.formatted":       true,
	"urn:ietf:params:scim:schemas:core:2.0:User:name.familyName":      true,
	"urn:ietf:params:scim:schemas:core:2.0:User:name.givenName":       true,
	"urn:ietf:params:scim:schemas:core:2.0:User:name.middleName":      true,
	"urn:ietf:params:scim:schemas:core:2.0:User:name.honorificPrefix": true,
	"urn:ietf:params:scim:schemas:core:2.0:User:name.honorificSuffix": true,
	"urn:ietf:params:scim:schemas:core:2.0:User:nickName":             true,
	"urn:ietf:params:scim:schemas:core:2.0:User:profileUrl":           true,
	"urn:ietf:params:scim:schemas:core:2.0:User:title":                true,
	"urn:ietf:params:scim:schemas:core:2.0:User:userType":             true,
	"urn:ietf:params:scim:schemas:core:2.0:User:preferredLanguage":    true,
	"urn:ietf:params:scim:schemas:core:2.0:User:locale":               true,
	"urn:ietf:params:scim:schemas:core:2.0:User:timezone":             true,
	"urn:ietf:params:scim:schemas:core:2.0:User:active":               true,
	// "urn:ietf:params:scim:schemas:core:2.0:User:emails":                              true,
	"urn:ietf:params:scim:schemas:core:2.0:User:emails.display": true,
	"urn:ietf:params:scim:schemas:core:2.0:User:emails.type":    true,
	"urn:ietf:params:scim:schemas:core:2.0:User:emails.primary": true,
	// "urn:ietf:params:scim:schemas:core:2.0:User:phoneNumbers":                        true,
	"urn:ietf:params:scim:schemas:core:2.0:User:phoneNumbers.value":   true,
	"urn:ietf:params:scim:schemas:core:2.0:User:phoneNumbers.display": true,
	"urn:ietf:params:scim:schemas:core:2.0:User:phoneNumbers.type":    true,
	"urn:ietf:params:scim:schemas:core:2.0:User:phoneNumbers.primary": true,
	// "urn:ietf:params:scim:schemas:core:2.0:User:ims":                                 true,
	"urn:ietf:params:scim:schemas:core:2.0:User:ims.value":   true,
	"urn:ietf:params:scim:schemas:core:2.0:User:ims.display": true,
	"urn:ietf:params:scim:schemas:core:2.0:User:ims.type":    true,
	"urn:ietf:params:scim:schemas:core:2.0:User:ims.primary": true,
	// "urn:ietf:params:scim:schemas:core:2.0:User:photos":                              true,
	"urn:ietf:params:scim:schemas:core:2.0:User:photos.value":   true,
	"urn:ietf:params:scim:schemas:core:2.0:User:photos.display": true,
	"urn:ietf:params:scim:schemas:core:2.0:User:photos.type":    true,
	"urn:ietf:params:scim:schemas:core:2.0:User:photos.primary": true,
	// "urn:ietf:params:scim:schemas:core:2.0:User:addresses":                           true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.formatted":     true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.streetAddress": true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.locality":      true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.region":        true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.postalCode":    true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.country":       true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.type":          true,
	"urn:ietf:params:scim:schemas:core:2.0:User:addresses.primary":       true,
	// "urn:ietf:params:scim:schemas:core:2.0:User:groups":                              true,
	"urn:ietf:params:scim:schemas:core:2.0:User:groups.value":   true,
	"urn:ietf:params:scim:schemas:core:2.0:User:groups.$ref":    true,
	"urn:ietf:params:scim:schemas:core:2.0:User:groups.display": true,
	"urn:ietf:params:scim:schemas:core:2.0:User:groups.type":    true,
	// "urn:ietf:params:scim:schemas:core:2.0:User:entitlements":                        true,
	"urn:ietf:params:scim:schemas:core:2.0:User:entitlements.value":   true,
	"urn:ietf:params:scim:schemas:core:2.0:User:entitlements.display": true,
	"urn:ietf:params:scim:schemas:core:2.0:User:entitlements.type":    true,
	"urn:ietf:params:scim:schemas:core:2.0:User:entitlements.primary": true,
	// "urn:ietf:params:scim:schemas:core:2.0:User:roles":                               true,
	"urn:ietf:params:scim:schemas:core:2.0:User:roles.value":   true,
	"urn:ietf:params:scim:schemas:core:2.0:User:roles.display": true,
	"urn:ietf:params:scim:schemas:core:2.0:User:roles.type":    true,
	"urn:ietf:params:scim:schemas:core:2.0:User:roles.primary": true,
	// "urn:ietf:params:scim:schemas:core:2.0:User:x509Certificates":                    true,
	"urn:ietf:params:scim:schemas:core:2.0:User:x509Certificates.value":         true,
	"urn:ietf:params:scim:schemas:core:2.0:User:x509Certificates.display":       true,
	"urn:ietf:params:scim:schemas:core:2.0:User:x509Certificates.type":          true,
	"urn:ietf:params:scim:schemas:core:2.0:User:x509Certificates.primary":       true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:costCenter":     true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:department":     true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:division":       true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:employeeNumber": true,
	// "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager":             true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager.$ref":  true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager.value": true,
	"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:organization":  true,
}

var d = map[string]bool{
	"id":                                                         true,
	"schemas":                                                    true,
	"meta.resourceType":                                          true,
	"urn:ietf:params:scim:schemas:core:2.0:User:name.familyName": true,
	"urn:ietf:params:scim:schemas:core:2.0:User:emails.type":     true,
}

var projectionTestCases = []projectionTestCase{

	{
		nil,
		nil,
		allUserAttrs,
	},
	{
		[]string{},
		[]string{},
		allUserAttrs,
	},
	// "never" attributes must be ignored returing the minimal attribute's set
	{
		[]string{
			"password", // never
		},
		[]string{},
		minimalUserAttributes,
	},
	{
		[]string{
			"password", // never
			"id",       // always
		},
		[]string{},
		minimalUserAttributes,
	},
	{
		[]string{},
		[]string{
			"password", // never
			"id",       // always
		},
		allUserAttrs,
	},
	{
		[]string{},
		[]string{
			"password", // never
		},
		allUserAttrs,
	},
	{
		[]string{
			"id", // always
		},
		[]string{
			"password", // never
		},
		minimalUserAttributes,
	},
	// Including attributes
	{
		[]string{
			"displayName",
		},
		[]string{},
		a,
	},
	{
		[]string{
			"name.familyName",
			"emails.type",
		},
		[]string{},
		d,
	},
	{
		[]string{
			"displayName",
			"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager.displayName",
		},
		[]string{},
		b,
	},
	// Excluding attributes
	{
		[]string{},
		[]string{
			"meta.resourceType",
			"displayName",
			"urn:ietf:params:scim:schemas:core:2.0:User:emails.value",
			"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager.displayName",
		},
		c,
	},
	// Mixing included and excluded attributes
	{
		[]string{
			"displayName",
		},
		[]string{
			"meta.resourceType",
			"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager.displayName",
		},
		a,
	},
	// Excluded attributes override included attributes
	{
		[]string{
			"displayName",
			"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager.displayName",
		},
		[]string{
			"displayName",
			"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager.displayName",
		},
		minimalUserAttributes,
	},
	// Unknown attributes gets ignored
	{
		[]string{
			"phantom",
		},
		[]string{},
		allUserAttrs,
	},
	{
		[]string{},
		[]string{
			"phantom",
		},
		allUserAttrs,
	},
	// Unknown attributes listed with included/excluded attributes gets ignored
	{
		[]string{
			"phantom",
			"displayName",
		},
		[]string{},
		a,
	},
	{
		[]string{},
		[]string{
			"phantom",
			"displayName",
			"urn:ietf:params:scim:schemas:core:2.0:User:emails.value",
			"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager.displayName",
		},
		c,
	},
}

func TestProjection(t *testing.T) {
	rt := resTypeRepo.Pull("User")

	for _, tt := range projectionTestCases {
		// Make included attributes
		included := make([]*Path, len(tt.included))
		for idx, i := range tt.included {
			p := Parse(i)
			if !p.Undefined() {
				included[idx] = p
			} else {
				t.Fatalf("Wrong included attribute: %s", i)
			}
		}

		// Make excluded attributes
		excluded := make([]*Path, len(tt.excluded))
		for idx, e := range tt.excluded {
			p := Parse(e)
			if !p.Undefined() {
				excluded[idx] = p
			} else {
				t.Fatalf("Wrong excluded attribute: %s", e)
			}
		}

		result := Projection(rt, included, excluded)
		results := make(map[string]bool)
		for _, r := range result {
			results[r.String()] = true
		}
		require.Equal(t, tt.expected, results)
	}
}
