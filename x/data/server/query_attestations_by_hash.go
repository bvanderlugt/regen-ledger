package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/data"
)

// AttestationsByHash queries data attestations by the ContentHash of the data.
func (s serverImpl) AttestationsByHash(ctx context.Context, request *data.QueryAttestationsByHashRequest) (*data.QueryAttestationsByHashResponse, error) {
	if request.ContentHash == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("content hash cannot be empty")
	}

	iri, err := request.ContentHash.ToIRI()
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}

	dataId, err := s.stateStore.DataIDTable().GetByIri(ctx, iri)
	if err != nil {
		return nil, sdkerrors.ErrNotFound.Wrap("data record with content hash")
	}

	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}

	it, err := s.stateStore.DataAttestorTable().List(
		ctx,
		api.DataAttestorIdAttestorIndexKey{}.WithId(dataId.Id),
		ormlist.Paginate(pg),
	)
	if err != nil {
		return nil, err
	}
	defer it.Close()

	var attestations []*data.AttestationInfo
	for it.Next() {
		dataAttestor, err := it.Value()
		if err != nil {
			return nil, err
		}

		attestations = append(attestations, &data.AttestationInfo{
			Iri:       iri,
			Attestor:  sdk.AccAddress(dataAttestor.Attestor).String(),
			Timestamp: types.ProtobufToGogoTimestamp(dataAttestor.Timestamp),
		})
	}

	pageRes, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}

	return &data.QueryAttestationsByHashResponse{
		Attestations: attestations,
		Pagination:   pageRes,
	}, nil
}
