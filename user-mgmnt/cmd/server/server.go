package main

import (
	"net"
	registersvc "user-mgmnt/cmd/register-svc"
	"user-mgmnt/logging"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func init() {
	logging.ConfigureLog()
}

func main() {
	listner, err := net.Listen("tcp", ":50000")
	if err != nil {
		logging.Log.Fatal("failed to listen: ", zap.Error(err))
	}
	server := grpc.NewServer()
	registersvc.RegisterServers(server, listner)
	logging.Log.Info("server started...")
	logging.Log.Error("server started...")
	logging.Log.Warn("server started...")
	logging.Log.Debug("server started...")
	logging.Log.DPanic("server started...")
}
