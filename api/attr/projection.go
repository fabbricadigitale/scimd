package attr

import (
	"github.com/fabbricadigitale/scimd/schemas"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"gopkg.in/fatih/set.v0"
)

// Projection computes the attribute paths to be ...
func Projection(ctx *core.ResourceType, included []*Path, excluded []*Path) []*Path {
	always := set.New()
	for _, a := range Paths(ctx, withReturned(schemas.ReturnedAlways)) {
		always.Add(a)
	}

	defaults := set.New()
	if len(included) > 0 {
		for _, d := range included {
			defaults.Add(d)
		}
	} else {
		for _, d := range Paths(ctx, withReturned(schemas.ReturnedDefault)) {
			defaults.Add(d)
		}
	}

	exclusions := set.New()
	for _, e := range excluded {
		exclusions.Add(e)
	}

	ret := set.Union(set.Difference(defaults, exclusions), always)

	return getPathSlice(ret)
}

func getPathSlice(ret set.Interface) []*Path {
	res := make([]*Path, 0)
	for _, item := range ret.List() {
		v, ok := item.(*Path)
		if !ok {
			continue
		}

		res = append(res, v)
	}
	return res
}

func withReturned(equalTo string) func(attribute *core.Attribute) bool {
	return func(attribute *core.Attribute) bool {
		return attribute.Returned == equalTo
	}
}
