// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pay-server/v1/pay-service.proto

package v1

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

// Validate checks the field values on PayRequest with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PayRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PayRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PayRequestMultiError, or
// nil if none found.
func (m *PayRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PayRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderId

	// no validation rules for UserId

	// no validation rules for Amount

	if len(errors) > 0 {
		return PayRequestMultiError(errors)
	}

	return nil
}

// PayRequestMultiError is an error wrapping multiple validation errors
// returned by PayRequest.ValidateAll() if the designated constraints aren't met.
type PayRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PayRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PayRequestMultiError) AllErrors() []error { return m }

// PayRequestValidationError is the validation error returned by
// PayRequest.Validate if the designated constraints aren't met.
type PayRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PayRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PayRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PayRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PayRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PayRequestValidationError) ErrorName() string { return "PayRequestValidationError" }

// Error satisfies the builtin error interface
func (e PayRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPayRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PayRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PayRequestValidationError{}

// Validate checks the field values on PayResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PayResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PayResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PayResponseMultiError, or
// nil if none found.
func (m *PayResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PayResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderId

	// no validation rules for UserId

	// no validation rules for Amount

	// no validation rules for Status

	if len(errors) > 0 {
		return PayResponseMultiError(errors)
	}

	return nil
}

// PayResponseMultiError is an error wrapping multiple validation errors
// returned by PayResponse.ValidateAll() if the designated constraints aren't met.
type PayResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PayResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PayResponseMultiError) AllErrors() []error { return m }

// PayResponseValidationError is the validation error returned by
// PayResponse.Validate if the designated constraints aren't met.
type PayResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PayResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PayResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PayResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PayResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PayResponseValidationError) ErrorName() string { return "PayResponseValidationError" }

// Error satisfies the builtin error interface
func (e PayResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPayResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PayResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PayResponseValidationError{}

// Validate checks the field values on DetailRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DetailRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DetailRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DetailRequestMultiError, or
// nil if none found.
func (m *DetailRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DetailRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderId

	if len(errors) > 0 {
		return DetailRequestMultiError(errors)
	}

	return nil
}

// DetailRequestMultiError is an error wrapping multiple validation errors
// returned by DetailRequest.ValidateAll() if the designated constraints
// aren't met.
type DetailRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DetailRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DetailRequestMultiError) AllErrors() []error { return m }

// DetailRequestValidationError is the validation error returned by
// DetailRequest.Validate if the designated constraints aren't met.
type DetailRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DetailRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DetailRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DetailRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DetailRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DetailRequestValidationError) ErrorName() string { return "DetailRequestValidationError" }

// Error satisfies the builtin error interface
func (e DetailRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDetailRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DetailRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DetailRequestValidationError{}

// Validate checks the field values on DetailResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DetailResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DetailResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DetailResponseMultiError,
// or nil if none found.
func (m *DetailResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DetailResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderId

	// no validation rules for UserId

	// no validation rules for Amount

	// no validation rules for Status

	if len(errors) > 0 {
		return DetailResponseMultiError(errors)
	}

	return nil
}

// DetailResponseMultiError is an error wrapping multiple validation errors
// returned by DetailResponse.ValidateAll() if the designated constraints
// aren't met.
type DetailResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DetailResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DetailResponseMultiError) AllErrors() []error { return m }

// DetailResponseValidationError is the validation error returned by
// DetailResponse.Validate if the designated constraints aren't met.
type DetailResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DetailResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DetailResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DetailResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DetailResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DetailResponseValidationError) ErrorName() string { return "DetailResponseValidationError" }

// Error satisfies the builtin error interface
func (e DetailResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDetailResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DetailResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DetailResponseValidationError{}
