// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common/v1/errordetails.proto

package commonv1

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

// Validate checks the field values on AttemptInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AttemptInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AttemptInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AttemptInfoMultiError, or
// nil if none found.
func (m *AttemptInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *AttemptInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetAttempts() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AttemptInfoValidationError{
						field:  fmt.Sprintf("Attempts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AttemptInfoValidationError{
						field:  fmt.Sprintf("Attempts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AttemptInfoValidationError{
					field:  fmt.Sprintf("Attempts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return AttemptInfoMultiError(errors)
	}

	return nil
}

// AttemptInfoMultiError is an error wrapping multiple validation errors
// returned by AttemptInfo.ValidateAll() if the designated constraints aren't met.
type AttemptInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttemptInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttemptInfoMultiError) AllErrors() []error { return m }

// AttemptInfoValidationError is the validation error returned by
// AttemptInfo.Validate if the designated constraints aren't met.
type AttemptInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttemptInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttemptInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttemptInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttemptInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttemptInfoValidationError) ErrorName() string { return "AttemptInfoValidationError" }

// Error satisfies the builtin error interface
func (e AttemptInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttemptInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttemptInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttemptInfoValidationError{}

// Validate checks the field values on ExternalError with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ExternalError) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExternalError with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ExternalErrorMultiError, or
// nil if none found.
func (m *ExternalError) ValidateAll() error {
	return m.validate(true)
}

func (m *ExternalError) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	for idx, item := range m.GetDetails() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ExternalErrorValidationError{
						field:  fmt.Sprintf("Details[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ExternalErrorValidationError{
						field:  fmt.Sprintf("Details[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExternalErrorValidationError{
					field:  fmt.Sprintf("Details[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ExternalErrorMultiError(errors)
	}

	return nil
}

// ExternalErrorMultiError is an error wrapping multiple validation errors
// returned by ExternalError.ValidateAll() if the designated constraints
// aren't met.
type ExternalErrorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExternalErrorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExternalErrorMultiError) AllErrors() []error { return m }

// ExternalErrorValidationError is the validation error returned by
// ExternalError.Validate if the designated constraints aren't met.
type ExternalErrorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExternalErrorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExternalErrorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExternalErrorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExternalErrorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExternalErrorValidationError) ErrorName() string { return "ExternalErrorValidationError" }

// Error satisfies the builtin error interface
func (e ExternalErrorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExternalError.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExternalErrorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExternalErrorValidationError{}

// Validate checks the field values on AttemptInfo_Attempt with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AttemptInfo_Attempt) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AttemptInfo_Attempt with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AttemptInfo_AttemptMultiError, or nil if none found.
func (m *AttemptInfo_Attempt) ValidateAll() error {
	return m.validate(true)
}

func (m *AttemptInfo_Attempt) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Number

	if all {
		switch v := interface{}(m.GetStatus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AttemptInfo_AttemptValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AttemptInfo_AttemptValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttemptInfo_AttemptValidationError{
				field:  "Status",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AttemptInfo_AttemptMultiError(errors)
	}

	return nil
}

// AttemptInfo_AttemptMultiError is an error wrapping multiple validation
// errors returned by AttemptInfo_Attempt.ValidateAll() if the designated
// constraints aren't met.
type AttemptInfo_AttemptMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttemptInfo_AttemptMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttemptInfo_AttemptMultiError) AllErrors() []error { return m }

// AttemptInfo_AttemptValidationError is the validation error returned by
// AttemptInfo_Attempt.Validate if the designated constraints aren't met.
type AttemptInfo_AttemptValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttemptInfo_AttemptValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttemptInfo_AttemptValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttemptInfo_AttemptValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttemptInfo_AttemptValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttemptInfo_AttemptValidationError) ErrorName() string {
	return "AttemptInfo_AttemptValidationError"
}

// Error satisfies the builtin error interface
func (e AttemptInfo_AttemptValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttemptInfo_Attempt.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttemptInfo_AttemptValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttemptInfo_AttemptValidationError{}
