package api

import (
	"context"

	"github.com/NpoolPlatform/account-manager/api/transfer"

	account "github.com/NpoolPlatform/message/npool/account/mgr/v1"

	account1 "github.com/NpoolPlatform/account-manager/api/account"
	"github.com/NpoolPlatform/account-manager/api/deposit"
	"github.com/NpoolPlatform/account-manager/api/goodbenefit"
	"github.com/NpoolPlatform/account-manager/api/limitation"
	"github.com/NpoolPlatform/account-manager/api/payment"
	"github.com/NpoolPlatform/account-manager/api/platform"
	"github.com/NpoolPlatform/account-manager/api/user"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	account.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	account.RegisterManagerServer(server, &Server{})
	goodbenefit.Register(server)
	account1.Register(server)
	user.Register(server)
	payment.Register(server)
	platform.Register(server)
	limitation.Register(server)
	deposit.Register(server)
	transfer.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := account.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
