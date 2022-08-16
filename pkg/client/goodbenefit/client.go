//nolint:dupl
package goodbenefit

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/goodbenefit"

	constant "github.com/NpoolPlatform/account-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get goodbenefit connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateGoodBenefit(ctx context.Context, in *npool.AccountReq) (*npool.Account, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateAccount(ctx, &npool.CreateAccountRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create goodbenefit: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create goodbenefit: %v", err)
	}
	return info.(*npool.Account), nil
}

func CreateGoodBenefits(ctx context.Context, in []*npool.AccountReq) ([]*npool.Account, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateAccounts(ctx, &npool.CreateAccountsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create goodbenefits: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create goodbenefits: %v", err)
	}
	return infos.([]*npool.Account), nil
}

func GetGoodBenefit(ctx context.Context, id string) (*npool.Account, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetAccount(ctx, &npool.GetAccountRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get goodbenefit: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get goodbenefit: %v", err)
	}
	return info.(*npool.Account), nil
}

func GetGoodBenefitOnly(ctx context.Context, conds *npool.Conds) (*npool.Account, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetAccountOnly(ctx, &npool.GetAccountOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get goodbenefit: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get goodbenefit: %v", err)
	}
	return info.(*npool.Account), nil
}

func GetGoodBenefits(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.Account, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetAccounts(ctx, &npool.GetAccountsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get goodbenefits: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get goodbenefits: %v", err)
	}
	return infos.([]*npool.Account), total, nil
}

func ExistGoodBenefit(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistAccount(ctx, &npool.ExistAccountRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get goodbenefit: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get goodbenefit: %v", err)
	}
	return infos.(bool), nil
}

func ExistAccountConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistAccountConds(ctx, &npool.ExistAccountCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get goodbenefit: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get goodbenefit: %v", err)
	}
	return infos.(bool), nil
}

func CountGoodBenefits(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountAccounts(ctx, &npool.CountAccountsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count goodbenefit: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count goodbenefit: %v", err)
	}
	return infos.(uint32), nil
}
