// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: authority.proto

package pb

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

// Validate checks the field values on Token with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Token) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Token with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TokenMultiError, or nil if none found.
func (m *Token) ValidateAll() error {
	return m.validate(true)
}

func (m *Token) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for IdToken

	// no validation rules for RefreshToken

	if len(errors) > 0 {
		return TokenMultiError(errors)
	}
	return nil
}

// TokenMultiError is an error wrapping multiple validation errors returned by
// Token.ValidateAll() if the designated constraints aren't met.
type TokenMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TokenMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TokenMultiError) AllErrors() []error { return m }

// TokenValidationError is the validation error returned by Token.Validate if
// the designated constraints aren't met.
type TokenValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TokenValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TokenValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TokenValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TokenValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TokenValidationError) ErrorName() string { return "TokenValidationError" }

// Error satisfies the builtin error interface
func (e TokenValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sToken.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TokenValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TokenValidationError{}

// Validate checks the field values on CreateTokenReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CreateTokenReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTokenReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CreateTokenReqMultiError,
// or nil if none found.
func (m *CreateTokenReq) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTokenReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Uid

	if len(errors) > 0 {
		return CreateTokenReqMultiError(errors)
	}
	return nil
}

// CreateTokenReqMultiError is an error wrapping multiple validation errors
// returned by CreateTokenReq.ValidateAll() if the designated constraints
// aren't met.
type CreateTokenReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTokenReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTokenReqMultiError) AllErrors() []error { return m }

// CreateTokenReqValidationError is the validation error returned by
// CreateTokenReq.Validate if the designated constraints aren't met.
type CreateTokenReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTokenReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTokenReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTokenReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTokenReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTokenReqValidationError) ErrorName() string { return "CreateTokenReqValidationError" }

// Error satisfies the builtin error interface
func (e CreateTokenReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTokenReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTokenReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTokenReqValidationError{}

// Validate checks the field values on ListPublicKeysReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListPublicKeysReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListPublicKeysReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListPublicKeysReqMultiError, or nil if none found.
func (m *ListPublicKeysReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ListPublicKeysReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListPublicKeysReqMultiError(errors)
	}
	return nil
}

// ListPublicKeysReqMultiError is an error wrapping multiple validation errors
// returned by ListPublicKeysReq.ValidateAll() if the designated constraints
// aren't met.
type ListPublicKeysReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListPublicKeysReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListPublicKeysReqMultiError) AllErrors() []error { return m }

// ListPublicKeysReqValidationError is the validation error returned by
// ListPublicKeysReq.Validate if the designated constraints aren't met.
type ListPublicKeysReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPublicKeysReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPublicKeysReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPublicKeysReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPublicKeysReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPublicKeysReqValidationError) ErrorName() string {
	return "ListPublicKeysReqValidationError"
}

// Error satisfies the builtin error interface
func (e ListPublicKeysReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPublicKeysReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPublicKeysReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPublicKeysReqValidationError{}

// Validate checks the field values on CreateTokenRes with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CreateTokenRes) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTokenRes with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CreateTokenResMultiError,
// or nil if none found.
func (m *CreateTokenRes) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTokenRes) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetToken()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateTokenResValidationError{
					field:  "Token",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateTokenResValidationError{
					field:  "Token",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetToken()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateTokenResValidationError{
				field:  "Token",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateTokenResMultiError(errors)
	}
	return nil
}

// CreateTokenResMultiError is an error wrapping multiple validation errors
// returned by CreateTokenRes.ValidateAll() if the designated constraints
// aren't met.
type CreateTokenResMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTokenResMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTokenResMultiError) AllErrors() []error { return m }

// CreateTokenResValidationError is the validation error returned by
// CreateTokenRes.Validate if the designated constraints aren't met.
type CreateTokenResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTokenResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTokenResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTokenResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTokenResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTokenResValidationError) ErrorName() string { return "CreateTokenResValidationError" }

// Error satisfies the builtin error interface
func (e CreateTokenResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTokenRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTokenResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTokenResValidationError{}

// Validate checks the field values on ListPublicKeysRes with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListPublicKeysRes) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListPublicKeysRes with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListPublicKeysResMultiError, or nil if none found.
func (m *ListPublicKeysRes) ValidateAll() error {
	return m.validate(true)
}

func (m *ListPublicKeysRes) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Jwks

	if len(errors) > 0 {
		return ListPublicKeysResMultiError(errors)
	}
	return nil
}

// ListPublicKeysResMultiError is an error wrapping multiple validation errors
// returned by ListPublicKeysRes.ValidateAll() if the designated constraints
// aren't met.
type ListPublicKeysResMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListPublicKeysResMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListPublicKeysResMultiError) AllErrors() []error { return m }

// ListPublicKeysResValidationError is the validation error returned by
// ListPublicKeysRes.Validate if the designated constraints aren't met.
type ListPublicKeysResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPublicKeysResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPublicKeysResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPublicKeysResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPublicKeysResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPublicKeysResValidationError) ErrorName() string {
	return "ListPublicKeysResValidationError"
}

// Error satisfies the builtin error interface
func (e ListPublicKeysResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPublicKeysRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPublicKeysResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPublicKeysResValidationError{}
