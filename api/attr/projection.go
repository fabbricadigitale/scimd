package attr

import (
	"github.com/fabbricadigitale/scimd/schemas"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"gopkg.in/fatih/set.v0"
)

// Projection computes a set of attribute paths
//
// The logic it enables is the one described within the RFC at sections:
//
// - https://tools.ietf.org/html/rfc7644#section-3.4.2.5
//
// - https://tools.ietf.org/html/rfc7644#section-3.4.3
//
// - https://tools.ietf.org/html/rfc7644#section-3.9
//
// In short the logic is (D - E - N) ∪ A where:
//
// I: { i | i explicitly included by the user }
//
// E: { e | e explicitly excluded by the user }
//
// N: { n | n.Returned == "never" }
//
// A: { a | a.Returned == "always" }
//
// D: I if I != ∅ || { d | d.Returned == "default" }
//
// Furthermore, attributes with Mutability == "writeOnly" cannot be returned too.
// So, pragmatically, they are treated as Returned == "never", as per:
// - https://tools.ietf.org/html/rfc7643#section-7
//
func Projection(ctx *core.ResourceType, included []*Path, excluded []*Path) []*Path {
	always := set.New()
	for _, a := range Paths(ctx, withReturned(schemas.ReturnedAlways)) {
		always.Add(*a)
	}

	never := set.New()
	for _, n := range Paths(ctx, cannotBeReturned) {
		never.Add(*n)
	}

	defaults := set.New()
	if len(included) > 0 {
		for _, d := range included {
			defaults.Add(*d)
		}
	} else {
		for _, d := range Paths(ctx, withReturned(schemas.ReturnedDefault)) {
			defaults.Add(*d)
		}
	}

	exclusions := set.New()
	for _, e := range excluded {
		exclusions.Add(*e)
	}

	ret := set.Union(set.Difference(defaults, exclusions, never), always)

	return getPathSlice(ret)
}

func getPathSlice(ret set.Interface) []*Path {
	res := make([]*Path, 0)
	for _, item := range ret.List() {
		v, ok := item.(Path)
		if !ok {
			continue
		}

		res = append(res, &v)
	}
	return res
}

func cannotBeReturned(attribute *core.Attribute) bool {
	return attribute.Returned == schemas.ReturnedNever || attribute.Mutability == schemas.MutabilityWriteOnly || len(attribute.SubAttributes) > 0
}

func withReturned(equalTo string) func(attribute *core.Attribute) bool {
	return func(attribute *core.Attribute) bool {
		return attribute.Returned == equalTo
	}
}
