package core

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// Projects queries all projects.
func (k Keeper) Projects(ctx context.Context, request *core.QueryProjectsRequest) (*core.QueryProjectsResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, err
	}

	it, err := k.stateStore.ProjectTable().List(ctx, api.ProjectIdIndexKey{}, ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}
	defer it.Close()

	projects := make([]*core.ProjectInfo, 0)
	for it.Next() {
		project, err := it.Value()
		if err != nil {
			return nil, err
		}

		admin := sdk.AccAddress(project.Admin)

		class, err := k.stateStore.ClassTable().Get(ctx, project.ClassKey)
		if err != nil {
			return nil, err
		}

		info := core.ProjectInfo{
			Id:           project.Id,
			Admin:        admin.String(),
			ClassId:      class.Id,
			Jurisdiction: project.Jurisdiction,
			Metadata:     project.Metadata,
			ReferenceId:  project.ReferenceId,
		}

		projects = append(projects, &info)
	}

	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}

	return &core.QueryProjectsResponse{
		Projects:   projects,
		Pagination: pr,
	}, nil
}
