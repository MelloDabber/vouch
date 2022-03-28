// Copyright © 2022 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mock

import (
	"context"
	"errors"

	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/attestantio/vouch/services/feerecipientprovider"
)

// Service is a mock fee recipient provider.
type Service struct{}

// New creates a new mock fee recipient provider.
func New() feerecipientprovider.Service {
	return &Service{}
}

// FeeRecipients returns the fee recipients for the given validators.
func (s *Service) FeeRecipients(_ context.Context,
	_ []phase0.ValidatorIndex,
) (
	map[phase0.ValidatorIndex]bellatrix.ExecutionAddress,
	error,
) {
	return map[phase0.ValidatorIndex]bellatrix.ExecutionAddress{
		0: {0x00},
		1: {0x01},
		2: {0x02},
	}, nil
}

// ErroringService is a mock fee recipient provider.
type ErroringService struct{}

// NewErroring creates a new mock fee recipient provider.
func NewErroring() feerecipientprovider.Service {
	return &ErroringService{}
}

// FeeRecipients returns the fee recipients for the given validators.
func (s *ErroringService) FeeRecipients(_ context.Context,
	_ []phase0.ValidatorIndex,
) (
	map[phase0.ValidatorIndex]bellatrix.ExecutionAddress,
	error,
) {
	return nil, errors.New("error")
}
