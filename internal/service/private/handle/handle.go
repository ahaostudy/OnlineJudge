package handle

import (
	rpcPrivate "main/api/private"
)

type PrivateServer struct {
	rpcPrivate.UnimplementedPrivateServiceServer
}
