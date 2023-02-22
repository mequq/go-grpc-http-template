package server

import "github.com/google/wire"

var ServerProviderSet = wire.NewSet(NewGrpcServer, NewHttpServer)
