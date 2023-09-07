// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: post.proto

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

// Validate checks the field values on ListPostRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListPostRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListPostRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListPostRequestMultiError, or nil if none found.
func (m *ListPostRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListPostRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PageNum

	// no validation rules for PageSize

	// no validation rules for Status

	// no validation rules for PostName

	// no validation rules for PostCode

	if len(errors) > 0 {
		return ListPostRequestMultiError(errors)
	}

	return nil
}

// ListPostRequestMultiError is an error wrapping multiple validation errors
// returned by ListPostRequest.ValidateAll() if the designated constraints
// aren't met.
type ListPostRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListPostRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListPostRequestMultiError) AllErrors() []error { return m }

// ListPostRequestValidationError is the validation error returned by
// ListPostRequest.Validate if the designated constraints aren't met.
type ListPostRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPostRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPostRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPostRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPostRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPostRequestValidationError) ErrorName() string { return "ListPostRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListPostRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPostRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPostRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPostRequestValidationError{}

// Validate checks the field values on ListPostReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListPostReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListPostReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListPostReplyMultiError, or
// nil if none found.
func (m *ListPostReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListPostReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	// no validation rules for PageNum

	// no validation rules for PageSize

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListPostReplyValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListPostReplyValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListPostReplyValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListPostReplyMultiError(errors)
	}

	return nil
}

// ListPostReplyMultiError is an error wrapping multiple validation errors
// returned by ListPostReply.ValidateAll() if the designated constraints
// aren't met.
type ListPostReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListPostReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListPostReplyMultiError) AllErrors() []error { return m }

// ListPostReplyValidationError is the validation error returned by
// ListPostReply.Validate if the designated constraints aren't met.
type ListPostReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPostReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPostReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPostReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPostReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPostReplyValidationError) ErrorName() string { return "ListPostReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListPostReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPostReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPostReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPostReplyValidationError{}

// Validate checks the field values on CreatePostRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreatePostRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreatePostRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreatePostRequestMultiError, or nil if none found.
func (m *CreatePostRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreatePostRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PostName

	// no validation rules for PostCode

	// no validation rules for Sort

	// no validation rules for Status

	// no validation rules for Remark

	if len(errors) > 0 {
		return CreatePostRequestMultiError(errors)
	}

	return nil
}

// CreatePostRequestMultiError is an error wrapping multiple validation errors
// returned by CreatePostRequest.ValidateAll() if the designated constraints
// aren't met.
type CreatePostRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreatePostRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreatePostRequestMultiError) AllErrors() []error { return m }

// CreatePostRequestValidationError is the validation error returned by
// CreatePostRequest.Validate if the designated constraints aren't met.
type CreatePostRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePostRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePostRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePostRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePostRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePostRequestValidationError) ErrorName() string {
	return "CreatePostRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreatePostRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePostRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePostRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePostRequestValidationError{}

// Validate checks the field values on CreatePostReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreatePostReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreatePostReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreatePostReplyMultiError, or nil if none found.
func (m *CreatePostReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreatePostReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreatePostReplyMultiError(errors)
	}

	return nil
}

// CreatePostReplyMultiError is an error wrapping multiple validation errors
// returned by CreatePostReply.ValidateAll() if the designated constraints
// aren't met.
type CreatePostReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreatePostReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreatePostReplyMultiError) AllErrors() []error { return m }

// CreatePostReplyValidationError is the validation error returned by
// CreatePostReply.Validate if the designated constraints aren't met.
type CreatePostReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePostReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePostReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePostReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePostReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePostReplyValidationError) ErrorName() string { return "CreatePostReplyValidationError" }

// Error satisfies the builtin error interface
func (e CreatePostReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePostReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePostReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePostReplyValidationError{}

// Validate checks the field values on UpdatePostRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdatePostRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdatePostRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdatePostRequestMultiError, or nil if none found.
func (m *UpdatePostRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdatePostRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PostId

	// no validation rules for PostName

	// no validation rules for PostCode

	// no validation rules for Sort

	// no validation rules for Status

	// no validation rules for Remark

	if len(errors) > 0 {
		return UpdatePostRequestMultiError(errors)
	}

	return nil
}

// UpdatePostRequestMultiError is an error wrapping multiple validation errors
// returned by UpdatePostRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdatePostRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdatePostRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdatePostRequestMultiError) AllErrors() []error { return m }

// UpdatePostRequestValidationError is the validation error returned by
// UpdatePostRequest.Validate if the designated constraints aren't met.
type UpdatePostRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatePostRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatePostRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatePostRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatePostRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatePostRequestValidationError) ErrorName() string {
	return "UpdatePostRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdatePostRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatePostRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatePostRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatePostRequestValidationError{}

// Validate checks the field values on UpdatePostReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdatePostReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdatePostReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdatePostReplyMultiError, or nil if none found.
func (m *UpdatePostReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdatePostReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdatePostReplyMultiError(errors)
	}

	return nil
}

// UpdatePostReplyMultiError is an error wrapping multiple validation errors
// returned by UpdatePostReply.ValidateAll() if the designated constraints
// aren't met.
type UpdatePostReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdatePostReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdatePostReplyMultiError) AllErrors() []error { return m }

// UpdatePostReplyValidationError is the validation error returned by
// UpdatePostReply.Validate if the designated constraints aren't met.
type UpdatePostReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatePostReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatePostReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatePostReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatePostReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatePostReplyValidationError) ErrorName() string { return "UpdatePostReplyValidationError" }

// Error satisfies the builtin error interface
func (e UpdatePostReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatePostReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatePostReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatePostReplyValidationError{}

// Validate checks the field values on DeletePostRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeletePostRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeletePostRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeletePostRequestMultiError, or nil if none found.
func (m *DeletePostRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeletePostRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeletePostRequestMultiError(errors)
	}

	return nil
}

// DeletePostRequestMultiError is an error wrapping multiple validation errors
// returned by DeletePostRequest.ValidateAll() if the designated constraints
// aren't met.
type DeletePostRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeletePostRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeletePostRequestMultiError) AllErrors() []error { return m }

// DeletePostRequestValidationError is the validation error returned by
// DeletePostRequest.Validate if the designated constraints aren't met.
type DeletePostRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeletePostRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeletePostRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeletePostRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeletePostRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeletePostRequestValidationError) ErrorName() string {
	return "DeletePostRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeletePostRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeletePostRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeletePostRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeletePostRequestValidationError{}

// Validate checks the field values on DeletePostReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeletePostReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeletePostReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeletePostReplyMultiError, or nil if none found.
func (m *DeletePostReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeletePostReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeletePostReplyMultiError(errors)
	}

	return nil
}

// DeletePostReplyMultiError is an error wrapping multiple validation errors
// returned by DeletePostReply.ValidateAll() if the designated constraints
// aren't met.
type DeletePostReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeletePostReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeletePostReplyMultiError) AllErrors() []error { return m }

// DeletePostReplyValidationError is the validation error returned by
// DeletePostReply.Validate if the designated constraints aren't met.
type DeletePostReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeletePostReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeletePostReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeletePostReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeletePostReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeletePostReplyValidationError) ErrorName() string { return "DeletePostReplyValidationError" }

// Error satisfies the builtin error interface
func (e DeletePostReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeletePostReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeletePostReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeletePostReplyValidationError{}
