package rpc

import (
	"errors"

	"github.com/AnhQuanTrl/proflow/internal/derrors"
	v1proto "github.com/golang/protobuf/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type ErrorDetailHandler func(*derrors.Error) proto.Message

type grpcError struct {
	code    codes.Code
	message string
	detail  proto.Message
}

// SendGRPCError converts error to a gRPC error.
//
// If err is not a domain error, it is converted to an internal error.
// Otherwise, the domain error is converted to a gRPC error with the corresponding code and message.
// Additionally, if detailHandler is not nil, the gRPC error is equipped with the returned proto message as status.Details.
func SendGRPCError(err error, detailHandler ErrorDetailHandler) error {
	if err == nil {
		return nil
	}

	var derr *derrors.Error
	if !errors.As(err, &derr) {
		return &grpcError{
			code:    codes.Internal,
			message: "internal error",
		}
	}

	grpcError := &grpcError{
		code:    codes.Unknown,
		message: derr.Message,
		detail:  detailHandler(derr),
	}
	switch derr.Type {
	case derrors.VALIDATION:
		grpcError.code = codes.InvalidArgument
	case derrors.NOT_FOUND:
		grpcError.code = codes.NotFound
	}
	return grpcError
}

func (e *grpcError) Error() string {
	return e.message
}

func (e *grpcError) GRPCStatus() *status.Status {
	s, _ := status.New(e.code, e.message).WithDetails(v1proto.MessageV1(e.detail))
	return s
}
