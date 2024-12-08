package exception

// to avoid conflict go errors package
// errno naming convention: MODULE_SUBMODULE_ACTION_HTTP-STATUS-CODE

import (
	"context"
	"errors"
	"fmt"
	"grpc-boot-starter/core/correlation"
	"strings"

	"github.com/bufbuild/protovalidate-go"
)

// ServiceError define
type ServiceError struct {
	Errno         string
	CorrelationId string
	Message       string
}

// NewServiceError creation
func NewServiceError(ctx context.Context, errno string, message string) *ServiceError {
	return &ServiceError{
		Errno:         errno,
		CorrelationId: correlation.CorrelationId(ctx),
		Message:       message,
	}
}

// Error returns the error message for the ServiceError type
func (e *ServiceError) Error() string {
	return fmt.Sprintf("errno: %s, correlationId: %s, %s", e.Errno, e.CorrelationId, e.Message)
}

func (s *ServiceError) Is(err error) bool {
	_, ok := err.(*ServiceError)
	return ok
}

// EntityNotFoundError define
type EntityNotFoundError struct {
	ID            any
	CorrelationId string
	Message       string
}

// NewEntityNotFoundError creation
func NewEntityNotFoundError(ctx context.Context, id any, message string) *EntityNotFoundError {
	return &EntityNotFoundError{
		ID:            id,
		CorrelationId: correlation.CorrelationId(ctx),
		Message:       message,
	}
}

// Error returns the error message for the EntityNotFoundError type
func (e *EntityNotFoundError) Error() string {
	return fmt.Sprintf("correlationId: %s, id: %v, %s", e.CorrelationId, e.ID, e.Message)
}

func (s *EntityNotFoundError) Is(err error) bool {
	_, ok := err.(*EntityNotFoundError)
	return ok
}

// RepositoryError
type RepositoryError struct {
	Errno         string
	CorrelationId string
	Message       string
}

// NewRepositoryError creation
func NewRepositoryError(ctx context.Context, errno string, message string) *RepositoryError {
	return &RepositoryError{
		Errno:         errno,
		CorrelationId: correlation.CorrelationId(ctx),
		Message:       message,
	}
}

// Error returns the error message for the ServiceError type
func (e *RepositoryError) Error() string {
	return fmt.Sprintf("errno: %s, correlationId: %s, %s", e.Errno, e.CorrelationId, e.Message)
}

func (s *RepositoryError) Is(err error) bool {
	_, ok := err.(*RepositoryError)
	return ok
}

// ACLError
type ACLError struct {
	Errno         string
	CorrelationId string
	Message       string
}

// NewACLError creation
func NewACLError(ctx context.Context, errno string, message string) *ACLError {
	return &ACLError{
		Errno:         errno,
		CorrelationId: correlation.CorrelationId(ctx),
		Message:       message,
	}
}

// Error returns the error message for the ServiceError type
func (e *ACLError) Error() string {
	return fmt.Sprintf("errno: %s, correlationId: %s, %s", e.Errno, e.CorrelationId, e.Message)
}

func (s *ACLError) Is(err error) bool {
	_, ok := err.(*ACLError)
	return ok
}

// AuthError
type AuthError struct {
	Errno         string
	CorrelationId string
	Message       string
}

// NewAuthError creation
func NewAuthError(ctx context.Context, errno string, message string) *AuthError {
	return &AuthError{
		Errno:         errno,
		CorrelationId: correlation.CorrelationId(ctx),
		Message:       message,
	}
}

// Error returns the error message for the ServiceError type
func (e *AuthError) Error() string {
	return fmt.Sprintf("errno: %s, correlationId: %s, %s", e.Errno, e.CorrelationId, e.Message)
}

func (s *AuthError) Is(err error) bool {
	_, ok := err.(*AuthError)
	return ok
}

// ValidationError
type ValidationError struct {
	Errno         string
	CorrelationId string
	Message       string
}

// NewValidationError creation
func NewValidationErrorOnFailed(ctx context.Context, errno string, err0 error) *ValidationError {
	message := fmt.Sprintf("input validation error on fields: %s", ParseFailedFields(err0))
	return &ValidationError{
		Errno:         errno,
		CorrelationId: correlation.CorrelationId(ctx),
		Message:       message,
	}
}

func NewValidationError(ctx context.Context, errno string, message string) *ValidationError {
	// message := fmt.Sprintf("input validation error on fields: %s", ParseFailedFields(err0))
	return &ValidationError{
		Errno:         errno,
		CorrelationId: correlation.CorrelationId(ctx),
		Message:       message,
	}
}

// Error returns the error message for the ServiceError type
func (e *ValidationError) Error() string {
	return fmt.Sprintf("errno: %s, correlationId: %s, %s", e.Errno, e.CorrelationId, e.Message)
}

func (s *ValidationError) Is(err error) bool {
	_, ok := err.(*ValidationError)
	return ok
}

// ParseFailedFields try to parse failed fields from error messages.
func ParseFailedFields(err error) string {
	if err == nil {
		return ""
	}
	var valErr *protovalidate.ValidationError
	if ok := errors.As(err, &valErr); ok {
		//pb := valErr.ToProto()
		// fmt.Printf("%T", pb)
		var fields = make([]string, 0, 10)
		for _, violation := range valErr.Violations {
			fmt.Println(*violation.FieldPath)
			fields = append(fields, *violation.FieldPath)
		}
		return strings.Join(fields, ", ")
	}
	return ""
}
