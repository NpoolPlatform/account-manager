package goodbenefit

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	val "github.com/NpoolPlatform/message/npool"

	testinit "github.com/NpoolPlatform/account-manager/pkg/testinit"
	npool "github.com/NpoolPlatform/message/npool/account/mgr/v1/goodbenefit"
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

var accountData = npool.Account{
	ID:            uuid.NewString(),
	GoodID:        uuid.NewString(),
	AccountID:     uuid.NewString(),
	Backup:        true,
	TransactionID: uuid.UUID{}.String(),
}

var (
	accountInfo = npool.AccountReq{
		ID:        &accountData.ID,
		GoodID:    &accountData.GoodID,
		AccountID: &accountData.AccountID,
		Backup:    &accountData.Backup,
	}
)

var info *npool.Account

func createAccount(t *testing.T) {
	var err error
	info, err = CreateAccount(context.Background(), &accountInfo)
	if assert.Nil(t, err) {
		accountData.CreatedAt = info.CreatedAt
		accountData.IntervalHours = info.IntervalHours
		assert.Equal(t, info, &accountData)
	}
}

func updateAccount(t *testing.T) {
	var err error

	backup := true
	tid := uuid.NewString()

	accountInfo.Backup = &backup
	accountInfo.TransactionID = &tid

	accountData.Backup = backup
	accountData.TransactionID = tid

	info, err = UpdateAccount(context.Background(), &accountInfo)
	if assert.Nil(t, err) {
		accountData.CreatedAt = info.CreatedAt
		accountData.IntervalHours = info.IntervalHours
		assert.Equal(t, info, &accountData)
	}
}

func createAccounts(t *testing.T) {
	accountDatas := []npool.Account{
		{
			ID:        uuid.NewString(),
			GoodID:    uuid.NewString(),
			AccountID: uuid.NewString(),
			Backup:    true,
		},
		{
			ID:        uuid.NewString(),
			GoodID:    uuid.NewString(),
			AccountID: uuid.NewString(),
			Backup:    true,
		},
	}

	Accounts := []*npool.AccountReq{}
	for key := range accountDatas {
		Accounts = append(Accounts, &npool.AccountReq{
			ID:            &accountDatas[key].ID,
			GoodID:        &accountDatas[key].GoodID,
			AccountID:     &accountDatas[key].AccountID,
			Backup:        &accountDatas[key].Backup,
			IntervalHours: &accountDatas[key].IntervalHours,
		})
	}

	infos, err := CreateAccounts(context.Background(), Accounts)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func getAccount(t *testing.T) {
	var err error
	info, err = GetAccount(context.Background(), info.ID)
	if assert.Nil(t, err) {
		accountData.IntervalHours = info.IntervalHours
		assert.Equal(t, info, &accountData)
	}
}

func getAccounts(t *testing.T) {
	infos, total, err := GetAccounts(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &accountData)
	}
}

func getAccountOnly(t *testing.T) {
	var err error
	info, err = GetAccountOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &accountData)
	}
}

func countAccounts(t *testing.T) {
	count, err := CountAccounts(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, uint32(1))
	}
}

func existAccount(t *testing.T) {
	exist, err := ExistAccount(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existAccountConds(t *testing.T) {
	exist, err := ExistAccountConds(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteAccount(t *testing.T) {
	info, err := DeleteAccount(context.Background(), info.ID)
	if assert.Nil(t, err) {
		accountData.IntervalHours = info.IntervalHours
		accountData.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &accountData)
	}
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createAccount", createAccount)
	t.Run("updateAccount", updateAccount)
	t.Run("createAccounts", createAccounts)
	t.Run("getAccount", getAccount)
	t.Run("getAccounts", getAccounts)
	t.Run("getAccountOnly", getAccountOnly)
	t.Run("existAccount", existAccount)
	t.Run("existAccountConds", existAccountConds)
	t.Run("count", countAccounts)
	t.Run("delete", deleteAccount)
}
