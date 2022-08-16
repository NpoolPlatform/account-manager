package user

import (
	"github.com/NpoolPlatform/message/npool/account/mgr/v1/user"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	user.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	user.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
