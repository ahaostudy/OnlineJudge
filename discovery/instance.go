package discovery

import "fmt"

type Server struct {
	Name    string
	Addr    string
	Version string
}

// 服务目标路径
func (srv *Server) target() string {
	if srv.Version == "" {
		return fmt.Sprintf("%s", srv.Name)
	}
	return fmt.Sprintf("%s/%s", srv.Name, srv.Version)
}

// 服务唯一key
func (srv *Server) key() string {
	return fmt.Sprintf("%s/%s", srv.target(), srv.Addr)
}
