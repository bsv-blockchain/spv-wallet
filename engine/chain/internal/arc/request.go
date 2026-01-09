package arc

import (
	"context"
	"errors"
	"net"

	"github.com/go-resty/resty/v2"

	"github.com/bsv-blockchain/spv-wallet/engine/chain/errors"
	chainerrors "github.com/bsv-blockchain/spv-wallet/engine/chain/errors"
	"github.com/bsv-blockchain/spv-wallet/engine/chain/models"
	chainmodels "github.com/bsv-blockchain/spv-wallet/engine/chain/models"
	"github.com/bsv-blockchain/spv-wallet/engine/spverrors"
	"github.com/bsv-blockchain/spv-wallet/models"
)

func (s *Service) prepareARCRequest(ctx context.Context) *resty.Request {
	req := s.httpClient.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json")

	if s.arcCfg.Token != "" {
		req.SetHeader("Authorization", s.arcCfg.Token)
	}

	if s.arcCfg.DeploymentID != "" {
		req.SetHeader("XDeployment-ID", s.arcCfg.DeploymentID)
	}

	return req
}

func (s *Service) wrapRequestError(err error) error {
	var e net.Error
	if errors.As(err, &e) {
		return chainerrors.ErrARCUnreachable.Wrap(e)
	}
	return spverrors.ErrInternal.Wrap(err)
}

func (s *Service) wrapARCError(baseError models.SPVError, errResult *chainmodels.ArcError) error {
	if errResult == nil || errResult.IsEmpty() {
		return baseError
	}
	return baseError.Wrap(errResult)
}
