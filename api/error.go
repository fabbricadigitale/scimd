package api

import "fmt"

type NotFoundError struct {
	Subject string
}

func (e *NotFoundError) Error() string {
	if e.Subject == "" {
		return "Not found"
	}
	return e.Subject + " not found"
}

// InvalidPathError is thrown when path attribute is invalid or malformed
type InvalidPathError struct {
	Path   string
	Detail string
}

func (e *InvalidPathError) Error() string {
	if len(e.Path) > 0 {
		return fmt.Sprintf("Path [%s] is invalid: %s", e.Path, e.Detail)
	}
	return fmt.Sprintf("Path is invalid: %s", e.Detail)
}

// InvalidFilterError is thrown when the specified filter syntax is invalid or the specified attribute and filter comparison is not supported
type InvalidFilterError struct {
	Filter string
	Detail string
}

func (e *InvalidFilterError) Error() string {
	if len(e.Filter) > 0 {
		return fmt.Sprintf("Filter '%s' is invalid: %s", e.Filter, e.Detail)
	}
	return fmt.Sprintf("Filter is invalid: %s", e.Detail)
}

// InternalServerError is a wrapper of a generic server error
type InternalServerError struct {
	Detail string
}

func (e *InternalServerError) Error() string {
	return fmt.Sprintf("Ops! Internal server error: %s", e.Detail)
}

// ResourceNotFoundError is a wrapper of a mongo.ResourceNotFoundError
type ResourceNotFoundError struct {
	Detail string
}

func (e *ResourceNotFoundError) Error() string {
	return fmt.Sprintf("Resource %s not found", e.Detail)
}

// MissingRequiredPropertyError is ...
type MissingRequiredPropertyError struct {
	Path   string
	Detail string
}

func (e *MissingRequiredPropertyError) Error() string {
	if e.Path != "" {
		return fmt.Sprintf("No attribute defined for '%s'", e.Path)
	}
	return fmt.Sprintf("%s", e.Detail)
}

// MutabilityError is ...
type MutabilityError struct {
	Path   string
	Detail string
}

func (e *MutabilityError) Error() string {
	if e.Path != "" {
		return fmt.Sprintf("Attribute '%s' is immutable", e.Path)
	}
	return fmt.Sprintf("%s", e.Detail)
}
