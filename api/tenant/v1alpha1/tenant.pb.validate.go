// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: tenant/v1alpha1/tenant.proto

package valpha1_tenant

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Tenant with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Tenant) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Tenant with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TenantMultiError, or nil if none found.
func (m *Tenant) ValidateAll() error {
	return m.validate(true)
}

func (m *Tenant) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetName() != "" {

		if !_Tenant_Name_Pattern.MatchString(m.GetName()) {
			err := TenantValidationError{
				field:  "Name",
				reason: "value does not match regex pattern \"^tenants/[^/]+$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetTitle() != "" {

		if l := utf8.RuneCountInString(m.GetTitle()); l < 1 || l > 100 {
			err := TenantValidationError{
				field:  "Title",
				reason: "value length must be between 1 and 100 runes, inclusive",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetDescription() != "" {

		if l := utf8.RuneCountInString(m.GetDescription()); l < 1 || l > 1000 {
			err := TenantValidationError{
				field:  "Description",
				reason: "value length must be between 1 and 1000 runes, inclusive",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	// no validation rules for ApiKey

	if len(errors) > 0 {
		return TenantMultiError(errors)
	}

	return nil
}

// TenantMultiError is an error wrapping multiple validation errors returned by
// Tenant.ValidateAll() if the designated constraints aren't met.
type TenantMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TenantMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TenantMultiError) AllErrors() []error { return m }

// TenantValidationError is the validation error returned by Tenant.Validate if
// the designated constraints aren't met.
type TenantValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TenantValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TenantValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TenantValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TenantValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TenantValidationError) ErrorName() string { return "TenantValidationError" }

// Error satisfies the builtin error interface
func (e TenantValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTenant.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TenantValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TenantValidationError{}

var _Tenant_Name_Pattern = regexp.MustCompile("^tenants/[^/]+$")

// Validate checks the field values on TenantRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TenantRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TenantRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TenantRequestMultiError, or
// nil if none found.
func (m *TenantRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *TenantRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if !_TenantRequest_Name_Pattern.MatchString(m.GetName()) {
		err := TenantRequestValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^tenants/[^/]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return TenantRequestMultiError(errors)
	}

	return nil
}

// TenantRequestMultiError is an error wrapping multiple validation errors
// returned by TenantRequest.ValidateAll() if the designated constraints
// aren't met.
type TenantRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TenantRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TenantRequestMultiError) AllErrors() []error { return m }

// TenantRequestValidationError is the validation error returned by
// TenantRequest.Validate if the designated constraints aren't met.
type TenantRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TenantRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TenantRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TenantRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TenantRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TenantRequestValidationError) ErrorName() string { return "TenantRequestValidationError" }

// Error satisfies the builtin error interface
func (e TenantRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTenantRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TenantRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TenantRequestValidationError{}

var _TenantRequest_Name_Pattern = regexp.MustCompile("^tenants/[^/]+$")

// Validate checks the field values on TenantsRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TenantsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TenantsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TenantsRequestMultiError,
// or nil if none found.
func (m *TenantsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *TenantsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PageSize

	// no validation rules for PageToken

	if len(errors) > 0 {
		return TenantsRequestMultiError(errors)
	}

	return nil
}

// TenantsRequestMultiError is an error wrapping multiple validation errors
// returned by TenantsRequest.ValidateAll() if the designated constraints
// aren't met.
type TenantsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TenantsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TenantsRequestMultiError) AllErrors() []error { return m }

// TenantsRequestValidationError is the validation error returned by
// TenantsRequest.Validate if the designated constraints aren't met.
type TenantsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TenantsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TenantsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TenantsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TenantsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TenantsRequestValidationError) ErrorName() string { return "TenantsRequestValidationError" }

// Error satisfies the builtin error interface
func (e TenantsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTenantsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TenantsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TenantsRequestValidationError{}

// Validate checks the field values on ListTenantsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListTenantsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListTenantsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListTenantsResponseMultiError, or nil if none found.
func (m *ListTenantsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListTenantsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetTenants() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListTenantsResponseValidationError{
						field:  fmt.Sprintf("Tenants[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListTenantsResponseValidationError{
						field:  fmt.Sprintf("Tenants[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListTenantsResponseValidationError{
					field:  fmt.Sprintf("Tenants[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for NextPageToken

	if len(errors) > 0 {
		return ListTenantsResponseMultiError(errors)
	}

	return nil
}

// ListTenantsResponseMultiError is an error wrapping multiple validation
// errors returned by ListTenantsResponse.ValidateAll() if the designated
// constraints aren't met.
type ListTenantsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListTenantsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListTenantsResponseMultiError) AllErrors() []error { return m }

// ListTenantsResponseValidationError is the validation error returned by
// ListTenantsResponse.Validate if the designated constraints aren't met.
type ListTenantsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListTenantsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListTenantsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListTenantsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListTenantsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListTenantsResponseValidationError) ErrorName() string {
	return "ListTenantsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListTenantsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListTenantsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListTenantsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListTenantsResponseValidationError{}
