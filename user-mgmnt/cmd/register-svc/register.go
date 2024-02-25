package registersvc

import (
	"net"
	"user-mgmnt/logging"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func RegisterServers(server *grpc.Server, listner net.Listener) {
	logging.Log.Info("server started...")
	if err := server.Serve(listner); err != nil {
		logging.Log.Fatal("failed to start server", zap.Error(err))
	}

}
