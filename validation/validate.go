package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

// Validator is a singleton instance providing validation functionalities
var Validator *validator.Validate

var (
	uni             *ut.UniversalTranslator
	translator      ut.Translator
	translatorFound bool
)

func init() {
	eng := en.New()
	uni := ut.New(eng, eng)
	translator, translatorFound = uni.GetTranslator("en")
	if !translatorFound {
		panic("translator not found")
	}

	Validator = validator.New()
	en_translations.RegisterDefaultTranslations(Validator, translator)

	Validator.RegisterValidation("startswith", startsWith)
	Validator.RegisterValidation("nstartswith", notStartsWith)
	Validator.RegisterValidation("endswith", endsWith)
	Validator.RegisterValidation("urn", urn)
	Validator.RegisterValidation("attrname", attrName)
	Validator.RegisterValidation("attrpath", attrPath)
	Validator.RegisterValidation("isfield", isField)
	Validator.RegisterValidation("isset", depsOn)
	Validator.RegisterValidation("uniqueattr", uniqueAttr)
	Validator.RegisterValidation("pathexists", pathExists)
	Validator.RegisterValidation("isdir", isDirectory)
	Validator.RegisterValidation("isfile", isFile)
}

// Errors translates and returns all the error messages by namespace
//
// Notice this means that when multiple errors exists only the last one will be selected.
// Order between different namespaces is not guaranteed.
func Errors(err error) string {
	trans := make(validator.ValidationErrorsTranslations)
	if err != nil {
		// Translate all error at once by namespace
		errs := err.(validator.ValidationErrors)
		trans = errs.Translate(translator)
	}

	res := []string{}
	for ns, te := range trans {
		parts := []string{}
		for i, p := range strings.Split(ns, ".") {
			if i > 1 {
				parts = append(parts, strings.ToLower(p))
			} else if i > 0 {
				parts = append(parts, p)
			}
		}

		withStructName := strings.Join(parts, " ")
		withoutFieldName := strings.TrimLeft(te, strings.Title(parts[len(parts)-1]))

		res = append(res, fmt.Sprintf("%s%s.", withStructName, withoutFieldName))
	}
	return strings.Join(res, "\n")
}
