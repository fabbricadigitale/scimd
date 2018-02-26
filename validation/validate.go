package validation

import (
	"fmt"
	"log"
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

var translations = []struct {
	tag             string
	translation     string
	override        bool
	customRegisFunc validator.RegisterTranslationsFunc
	customTransFunc validator.TranslationFunc
}{
	{
		tag:         "pathexists",
		translation: "{0} must be an already existing path",
		override:    false,
	},
	{
		tag:         "isdir",
		translation: "{0} must be a directory",
		override:    false,
	},
}

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
	Validator.RegisterValidation("hassubdir", hasSubDirectory)
	Validator.RegisterValidation("isfile", isFile)

	// Registration of translations
	var err error
	for _, t := range translations {
		if t.customTransFunc != nil && t.customRegisFunc != nil {
			err = Validator.RegisterTranslation(t.tag, translator, t.customRegisFunc, t.customTransFunc)
		} else if t.customTransFunc != nil && t.customRegisFunc == nil {
			err = Validator.RegisterTranslation(t.tag, translator, registrationFunc(t.tag, t.translation, t.override), t.customTransFunc)
		} else if t.customTransFunc == nil && t.customRegisFunc != nil {
			err = Validator.RegisterTranslation(t.tag, translator, t.customRegisFunc, translateFunc)
		} else {
			err = Validator.RegisterTranslation(t.tag, translator, registrationFunc(t.tag, t.translation, t.override), translateFunc)
		}

		if err != nil {
			panic(err)
		}
	}
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

		if len(parts) > 0 {
			withStructName := strings.Join(parts, " ")
			withoutFieldName := strings.TrimLeft(te, strings.Title(parts[len(parts)-1]))

			res = append(res, fmt.Sprintf("%s%s.", withStructName, withoutFieldName))
		} else {
			res = append(res, te)
		}
	}
	return strings.Join(res, "\n")
}

// Translation registration helper
func registrationFunc(tag string, translation string, override bool) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) (err error) {
		if err = ut.Add(tag, translation, override); err != nil {
			return
		}

		return
	}
}

// Translation helper
func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		log.Printf("warning: error translating FieldError: %#v", fe)
		return fe.(error).Error()
	}

	return t
}
