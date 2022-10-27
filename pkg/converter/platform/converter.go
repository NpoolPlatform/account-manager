package platform

import (
	"github.com/NpoolPlatform/account-manager/pkg/db/ent"
	uuid1 "github.com/NpoolPlatform/go-service-framework/pkg/const/uuid"
	account "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/platform"
)

func Ent2Grpc(row *ent.Platform) *npool.Account {
	if row == nil {
		return nil
	}

	var goodID *string
	if row.GoodID.String() != uuid1.InvalidUUIDStr {
		_str := row.GoodID.String()
		goodID = &_str
	}

	return &npool.Account{
		ID:        row.ID.String(),
		AccountID: row.AccountID.String(),
		UsedFor:   account.AccountUsedFor(account.AccountUsedFor_value[row.UsedFor]),
		GoodID:    goodID,
		Backup:    row.Backup,
		CreatedAt: row.CreatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.Platform) []*npool.Account {
	infos := []*npool.Account{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
