package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"google.golang.org/api/googleapi"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

type StakingAPIError struct {
	// Original error
	err error
	// httpCode is the HTTP response status code and will always be populated.
	httpCode int
	message  string
	details  ErrDetails
}

// Code presents the HTTP code of the APIError.
func (s *StakingAPIError) Code() int {
	return s.httpCode
}

// Message presents the high level error message of the APIError.
func (s *StakingAPIError) Message() string {
	return s.message
}

// Details presents the error details of the APIError.
func (s *StakingAPIError) Details() ErrDetails {
	return s.details
}

// Unwrap extracts the original error.
func (s *StakingAPIError) Unwrap() error {
	return s.err
}

// Error returns a readable representation of the StakingAPIError.
func (s *StakingAPIError) Error() string {
	// This happens when error returned wasn't a googleapi.Error in which case we don't know how to parse the error.
	// This should never happen as the auto-generated staking-api client should always return a googleapi.Error.
	// In this case, the best we can do is to print the original error.
	if s.httpCode == 0 && s.message == "" {
		return fmt.Sprintf("stakingapi error: %s", s.Unwrap().Error())
	} else if s.httpCode != 0 && s.message != "" {
		return fmt.Sprintf("stakingapi error: httpCode: %d message: %s", s.httpCode, s.message)
	} else if s.err != nil {
		// Either http code or message is missing. This should be very unlikely. It can happen when an internal
		// service simply returns the response body say as "Unauthorized". In this case, we also print the original error.
		return fmt.Sprintf("stakingapi error: httpCode: %d message: %s err: %s", s.httpCode, s.message, s.Unwrap().Error())
	} else {
		return fmt.Sprintf("stakingapi error: httpCode: %d message: %s", s.httpCode, s.message)
	}
}

func (s *StakingAPIError) Print() error {
	d := struct {
		Code    int        `json:"code"`
		Message string     `json:"message"`
		Details ErrDetails `json:"details"`
	}{
		Code:    s.Code(),
		Message: s.Message(),
		Details: s.Details(),
	}

	jsonBytes, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return err
	}

	log.Printf("\n%s", string(jsonBytes))

	return nil
}

// FromError parses a googleapi.Error and builds a StakingAPIError.
func FromError(err error) *StakingAPIError {
	sae := StakingAPIError{
		err: err,
	}

	var herr *googleapi.Error

	if errors.As(err, &herr) {
		// A proper user facing error message from the staking api ends up being in the googleapi.Error.Body instead of
		// other fields like Code, Message, Details already in it. This requires us to parse the Body field instead to
		// get the actual error details.
		// The body is a json string that can be parsed into a status.Status object.
		s := status.Status{}
		if err := protojson.Unmarshal([]byte(herr.Body), &s); err != nil {
			// If we can't parse the body, just return metadata from the googleapi.Error object.
			return &StakingAPIError{
				err:      herr.Unwrap(),
				httpCode: herr.Code,
				message:  herr.Message,
			}
		}

		sae.httpCode = herr.Code
		sae.message = s.GetMessage()

		// Coerce the Any messages into proto.Message then parse the details.
		details := []interface{}{}

		for _, any := range s.GetDetails() {
			m, err := any.UnmarshalNew()
			if err != nil {
				// Ignore malformed Any values.
				continue
			}

			details = append(details, m)
		}

		sae.details = parseDetails(details)
	}

	return &sae
}

// ErrDetails holds the google/rpc/error_details.proto messages.
type ErrDetails struct {
	ErrorInfo           *errdetails.ErrorInfo           `json:"errorInfo,omitempty"`
	BadRequest          *errdetails.BadRequest          `json:"badRequest,omitempty"`
	PreconditionFailure *errdetails.PreconditionFailure `json:"preconditionFailure,omitempty"`
	QuotaFailure        *errdetails.QuotaFailure        `json:"quotaFailure,omitempty"`
	RetryInfo           *errdetails.RetryInfo           `json:"retryInfo,omitempty"`
	ResourceInfo        *errdetails.ResourceInfo        `json:"resourceInfo,omitempty"`
	RequestInfo         *errdetails.RequestInfo         `json:"requestInfo,omitempty"`
	DebugInfo           *errdetails.DebugInfo           `json:"debugInfo,omitempty"`
	Help                *errdetails.Help                `json:"help,omitempty"`
	LocalizedMessage    *errdetails.LocalizedMessage    `json:"localizedMessage,omitempty"`

	// Unknown stores unidentifiable error details.
	Unknown []interface{} `json:"unknown,omitempty"`
}

// parseDetails accepts a slice of interface{} that should be backed by some
// sort of proto.Message that can be cast to the google/rpc/error_details.proto
// types.
//
// This is for internal use only.
func parseDetails(details []interface{}) ErrDetails {
	var ed ErrDetails

	for _, d := range details {
		switch d := d.(type) {
		case *errdetails.ErrorInfo:
			ed.ErrorInfo = d
		case *errdetails.BadRequest:
			ed.BadRequest = d
		case *errdetails.PreconditionFailure:
			ed.PreconditionFailure = d
		case *errdetails.QuotaFailure:
			ed.QuotaFailure = d
		case *errdetails.RetryInfo:
			ed.RetryInfo = d
		case *errdetails.ResourceInfo:
			ed.ResourceInfo = d
		case *errdetails.RequestInfo:
			ed.RequestInfo = d
		case *errdetails.DebugInfo:
			ed.DebugInfo = d
		case *errdetails.Help:
			ed.Help = d
		case *errdetails.LocalizedMessage:
			ed.LocalizedMessage = d
		default:
			ed.Unknown = append(ed.Unknown, d)
		}
	}

	return ed
}
