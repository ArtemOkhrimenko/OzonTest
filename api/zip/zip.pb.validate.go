// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/zip/zip.proto

package start_proto

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

// Validate checks the field values on GenerateShortLinkRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GenerateShortLinkRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GenerateShortLinkRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GenerateShortLinkRequestMultiError, or nil if none found.
func (m *GenerateShortLinkRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GenerateShortLinkRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if uri, err := url.Parse(m.GetLongLink()); err != nil {
		err = GenerateShortLinkRequestValidationError{
			field:  "LongLink",
			reason: "value must be a valid URI",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	} else if !uri.IsAbs() {
		err := GenerateShortLinkRequestValidationError{
			field:  "LongLink",
			reason: "value must be absolute",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GenerateShortLinkRequestMultiError(errors)
	}

	return nil
}

// GenerateShortLinkRequestMultiError is an error wrapping multiple validation
// errors returned by GenerateShortLinkRequest.ValidateAll() if the designated
// constraints aren't met.
type GenerateShortLinkRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GenerateShortLinkRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GenerateShortLinkRequestMultiError) AllErrors() []error { return m }

// GenerateShortLinkRequestValidationError is the validation error returned by
// GenerateShortLinkRequest.Validate if the designated constraints aren't met.
type GenerateShortLinkRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GenerateShortLinkRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GenerateShortLinkRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GenerateShortLinkRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GenerateShortLinkRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GenerateShortLinkRequestValidationError) ErrorName() string {
	return "GenerateShortLinkRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GenerateShortLinkRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGenerateShortLinkRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GenerateShortLinkRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GenerateShortLinkRequestValidationError{}

// Validate checks the field values on GenerateShortLinkResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GenerateShortLinkResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GenerateShortLinkResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GenerateShortLinkResponseMultiError, or nil if none found.
func (m *GenerateShortLinkResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GenerateShortLinkResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ShortLink

	if len(errors) > 0 {
		return GenerateShortLinkResponseMultiError(errors)
	}

	return nil
}

// GenerateShortLinkResponseMultiError is an error wrapping multiple validation
// errors returned by GenerateShortLinkResponse.ValidateAll() if the
// designated constraints aren't met.
type GenerateShortLinkResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GenerateShortLinkResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GenerateShortLinkResponseMultiError) AllErrors() []error { return m }

// GenerateShortLinkResponseValidationError is the validation error returned by
// GenerateShortLinkResponse.Validate if the designated constraints aren't met.
type GenerateShortLinkResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GenerateShortLinkResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GenerateShortLinkResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GenerateShortLinkResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GenerateShortLinkResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GenerateShortLinkResponseValidationError) ErrorName() string {
	return "GenerateShortLinkResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GenerateShortLinkResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGenerateShortLinkResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GenerateShortLinkResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GenerateShortLinkResponseValidationError{}

// Validate checks the field values on GetLongLinkRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetLongLinkRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetLongLinkRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetLongLinkRequestMultiError, or nil if none found.
func (m *GetLongLinkRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetLongLinkRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if uri, err := url.Parse(m.GetShortLink()); err != nil {
		err = GetLongLinkRequestValidationError{
			field:  "ShortLink",
			reason: "value must be a valid URI",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	} else if !uri.IsAbs() {
		err := GetLongLinkRequestValidationError{
			field:  "ShortLink",
			reason: "value must be absolute",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetLongLinkRequestMultiError(errors)
	}

	return nil
}

// GetLongLinkRequestMultiError is an error wrapping multiple validation errors
// returned by GetLongLinkRequest.ValidateAll() if the designated constraints
// aren't met.
type GetLongLinkRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetLongLinkRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetLongLinkRequestMultiError) AllErrors() []error { return m }

// GetLongLinkRequestValidationError is the validation error returned by
// GetLongLinkRequest.Validate if the designated constraints aren't met.
type GetLongLinkRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetLongLinkRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetLongLinkRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetLongLinkRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetLongLinkRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetLongLinkRequestValidationError) ErrorName() string {
	return "GetLongLinkRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetLongLinkRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetLongLinkRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetLongLinkRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetLongLinkRequestValidationError{}

// Validate checks the field values on GetLongLinkResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetLongLinkResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetLongLinkResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetLongLinkResponseMultiError, or nil if none found.
func (m *GetLongLinkResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetLongLinkResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for LongLink

	if len(errors) > 0 {
		return GetLongLinkResponseMultiError(errors)
	}

	return nil
}

// GetLongLinkResponseMultiError is an error wrapping multiple validation
// errors returned by GetLongLinkResponse.ValidateAll() if the designated
// constraints aren't met.
type GetLongLinkResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetLongLinkResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetLongLinkResponseMultiError) AllErrors() []error { return m }

// GetLongLinkResponseValidationError is the validation error returned by
// GetLongLinkResponse.Validate if the designated constraints aren't met.
type GetLongLinkResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetLongLinkResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetLongLinkResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetLongLinkResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetLongLinkResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetLongLinkResponseValidationError) ErrorName() string {
	return "GetLongLinkResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetLongLinkResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetLongLinkResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetLongLinkResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetLongLinkResponseValidationError{}
