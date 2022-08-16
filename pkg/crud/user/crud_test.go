package user

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/account-manager/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	testinit "github.com/NpoolPlatform/account-manager/pkg/testinit"
	valuedef "github.com/NpoolPlatform/message/npool"
	account "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/user"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var entity = ent.User{
	ID:         uuid.New(),
	AppID:      uuid.New(),
	UserID:     uuid.New(),
	CoinTypeID: uuid.New(),
	AccountID:  uuid.New(),
	UsedFor:    account.AccountUsedFor_GoodPayment.String(),
	Labels:     []string{"AAAAAAAAAAAAAAAAAA"},
}

var (
	id         = entity.ID.String()
	appID      = entity.AppID.String()
	userID     = entity.UserID.String()
	coinTypeID = entity.CoinTypeID.String()
	accountID  = entity.AccountID.String()
	usedFor    = account.AccountUsedFor_GoodPayment
	labels     = entity.Labels

	req = npool.AccountReq{
		ID:         &id,
		AppID:      &appID,
		UserID:     &userID,
		CoinTypeID: &coinTypeID,
		AccountID:  &accountID,
		UsedFor:    &usedFor,
		Labels:     labels,
	}
)

var info *ent.User

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &req)
	if assert.Nil(t, err) {
		entity.UpdatedAt = info.UpdatedAt
		entity.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), entity.String())
	}
}

func createBulk(t *testing.T) {
	entities := []*ent.User{
		{
			ID:         uuid.New(),
			AppID:      uuid.New(),
			UserID:     uuid.New(),
			CoinTypeID: uuid.New(),
			AccountID:  uuid.New(),
			UsedFor:    account.AccountUsedFor_GoodPayment.String(),
			Labels:     labels,
		},
		{
			ID:         uuid.New(),
			AppID:      uuid.New(),
			UserID:     uuid.New(),
			CoinTypeID: uuid.New(),
			AccountID:  uuid.New(),
			UsedFor:    account.AccountUsedFor_GoodPayment.String(),
			Labels:     labels,
		},
	}

	reqs := []*npool.AccountReq{}
	for _, _entity := range entities {
		_id := _entity.ID.String()
		_appID := _entity.AppID.String()
		_userID := _entity.UserID.String()
		_coinTypeID := _entity.CoinTypeID.String()
		_accountID := _entity.AccountID.String()
		_usedFor := account.AccountUsedFor_GoodPayment
		_labels := _entity.Labels

		reqs = append(reqs, &npool.AccountReq{
			ID:         &_id,
			AppID:      &_appID,
			UserID:     &_userID,
			CoinTypeID: &_coinTypeID,
			AccountID:  &_accountID,
			UsedFor:    &_usedFor,
			Labels:     _labels,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func update(t *testing.T) {
	req.Labels = []string{"BBBBBBBBBBBBBBBBBBBB"}
	entity.Labels = req.Labels

	info, err := Update(context.Background(), &req)
	if assert.Nil(t, err) {
		entity.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), entity.String())
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), entity.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), entity.String())
	}
}

func rows(t *testing.T) {
	infos, total, err := Rows(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		if assert.Equal(t, total, 1) {
			assert.Equal(t, infos[0].String(), entity.String())
		}
	}
}

func rowOnly(t *testing.T) {
	var err error
	info, err = RowOnly(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), entity.String())
	}
}

func count(t *testing.T) {
	count, err := Count(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, uint32(1))
	}
}

func exist(t *testing.T) {
	exist, err := Exist(context.Background(), entity.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existConds(t *testing.T) {
	exist, err := ExistConds(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteA(t *testing.T) {
	info, err := Delete(context.Background(), entity.ID)
	if assert.Nil(t, err) {
		entity.DeletedAt = info.DeletedAt
		assert.Equal(t, info.String(), entity.String())
	}
}

func TestAccount(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("createBulk", createBulk)
	t.Run("update", update)
	t.Run("row", row)
	t.Run("rows", rows)
	t.Run("rowOnly", rowOnly)
	t.Run("exist", exist)
	t.Run("existConds", existConds)
	t.Run("count", count)
	t.Run("delete", deleteA)
}
