package query

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/schemas/core"
)

type attributesTestCase struct {
	attributes *api.Attributes
	expected   map[string]bool
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

var attributesTestCases = []attributesTestCase{
	{
		attributes: &api.Attributes{
			Attributes:         nil,
			ExcludedAttributes: nil,
		},
		expected: allUserAttrs,
	},
	// Attributes that have Returned equal to "never" must be ignored
	// Attributes that have Returned equal to "always" must be always returned even if listed within the ExcludedAttributes
	{
		attributes: &api.Attributes{
			Attributes: nil,
			ExcludedAttributes: []string{
				"password", // never
				"id",       // always
			},
		},
		expected: allUserAttrs,
	},
	{
		attributes: &api.Attributes{
			Attributes: []string{
				"password", // never
			},
			ExcludedAttributes: nil,
		},
		expected: minimalUserAttributes,
	},
	{
		attributes: &api.Attributes{
			Attributes: []string{
				"password", // never
				"id",       // always
			},
			ExcludedAttributes: nil,
		},
		expected: minimalUserAttributes,
	},
	{
		attributes: &api.Attributes{
			Attributes: []string{
				"id", // always
			},
			ExcludedAttributes: []string{
				"password", // never
			},
		},
		expected: minimalUserAttributes,
	},
	// Including attributes
	{
		attributes: &api.Attributes{
			Attributes: []string{
				"displayName",
			},
			ExcludedAttributes: nil,
		},
		expected: a,
	},
	{
		attributes: &api.Attributes{
			Attributes: []string{
				"displayName",
				"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager.displayName",
			},
			ExcludedAttributes: nil,
		},
		expected: b,
	},
	{
		attributes: &api.Attributes{
			Attributes: []string{
				"name.familyName",
				"emails.type",
			},
			ExcludedAttributes: nil,
		},
		expected: d,
	},
	// Excluding attributes
	{
		attributes: &api.Attributes{
			Attributes: nil,
			ExcludedAttributes: []string{
				"meta.resourceType",
				"displayName",
				"urn:ietf:params:scim:schemas:core:2.0:User:emails.value",
				"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager.displayName",
			},
		},
		expected: c,
	},
	// Mixing included and excluded attributes
	{
		attributes: &api.Attributes{
			Attributes: []string{
				"displayName",
			},
			ExcludedAttributes: []string{
				"meta.resourceType",
				"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager.displayName",
			},
		},
		expected: a,
	},
	// Excluded attributes override included attributes
	{
		attributes: &api.Attributes{
			Attributes: []string{
				"displayName",
				"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager.displayName",
			},
			ExcludedAttributes: []string{
				"displayName",
				"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager.displayName",
			},
		},
		expected: minimalUserAttributes,
	},
	// Unknown attributes gets ignored
	{
		attributes: &api.Attributes{
			Attributes: []string{
				"phantom",
			},
			ExcludedAttributes: nil,
		},
		expected: allUserAttrs,
	},
	{
		attributes: &api.Attributes{
			Attributes: nil,
			ExcludedAttributes: []string{
				"phantom",
			},
		},
		expected: allUserAttrs,
	},
	// Unknown attributes listed with included/excluded attributes gets ignored
	{
		attributes: &api.Attributes{
			Attributes: []string{
				"displayName",
				"phantom",
			},
			ExcludedAttributes: nil,
		},
		expected: a,
	},
	{
		attributes: &api.Attributes{
			Attributes: nil,
			ExcludedAttributes: []string{
				"meta.resourceType",
				"displayName",
				"phantom",
				"urn:ietf:params:scim:schemas:core:2.0:User:emails.value",
				"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:manager.displayName",
			},
		},
		expected: c,
	},
}

func TestAttributes(t *testing.T) {

	resType := resTypeRepo.Pull("User")
	resTypeArray := make([]*core.ResourceType, 0)
	resTypeArray = append(resTypeArray, resType)

	for _, testCase := range attributesTestCases {
		result, err := Attributes(resTypeArray, testCase.attributes)
		require.NoError(t, err)

		results := make(map[string]bool)
		for k, v := range result {
			results[k.String()] = v
		}
		require.Equal(t, testCase.expected, results)
	}

}
