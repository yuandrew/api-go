// The MIT License (MIT)
//
// Copyright (c) 2020 Temporal Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package serviceerror

import (
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"

	"go.temporal.io/temporal-proto/errordetails"
)

type (
	// ClientVersionNotSupported represents client version is not supported error.
	ClientVersionNotSupported struct {
		Message           string
		FeatureVersion    string
		ClientImpl        string
		SupportedVersions string
		st                *status.Status
	}
)

// NewClientVersionNotSupported returns new ClientVersionNotSupported error.
func NewClientVersionNotSupported(message, featureVersion, clientImpl, supportedVersions string) *ClientVersionNotSupported {
	return &ClientVersionNotSupported{
		Message:           message,
		FeatureVersion:    featureVersion,
		ClientImpl:        clientImpl,
		SupportedVersions: supportedVersions,
	}
}

// Error returns string message.
func (e *ClientVersionNotSupported) Error() string {
	return e.Message
}

// GRPCStatus returns corresponding gRPC status.Status.
func (e *ClientVersionNotSupported) GRPCStatus() *status.Status {
	if e.st != nil {
		return e.st
	}

	st := status.New(codes.FailedPrecondition, e.Message)
	st, _ = st.WithDetails(
		&errordetails.ClientVersionNotSupportedFailure{
			FeatureVersion:    e.FeatureVersion,
			ClientImpl:        e.ClientImpl,
			SupportedVersions: e.SupportedVersions,
		},
	)
	return st
}

func newClientVersionNotSupported(st *status.Status, failure *errordetails.ClientVersionNotSupportedFailure) *ClientVersionNotSupported {
	return &ClientVersionNotSupported{
		Message:           st.Message(),
		FeatureVersion:    failure.FeatureVersion,
		ClientImpl:        failure.ClientImpl,
		SupportedVersions: failure.SupportedVersions,
		st:                st,
	}
}
