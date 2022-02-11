package log_v1

import (
	"fmt"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	status "google.golang.org/grpc/status"
)

// ErrorOffsetOutOfRange error.
type ErrorOffsetOutOfRange struct {
	Offset uint64
}

// GRPCStatus builds a grpc status object to give more details about errors
// to the client.
func (e ErrorOffsetOutOfRange) GRPCStatus() *status.Status {
	st := status.New(
		404,
		fmt.Sprintf("offset out of range: %d", e.Offset),
	)
	msg := fmt.Sprintf(
		"The requested offset is outside the log's range: %d",
		e.Offset,
	)
	d := &errdetails.LocalizedMessage{
		Locale:  "en-US",
		Message: msg,
	}

	std, err := st.WithDetails(d)
	if err != nil {
		return st
	}

	return std
}

func (e ErrorOffsetOutOfRange) Error() string {
	return e.GRPCStatus().Err().Error()
}
