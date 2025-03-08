package discovery

import "fmt"

type EtcdConfig struct {
	Address string `yaml:"address"`
}

type Server struct {
	Addr string
	Name string
}

func (srv *Server) GetPath() string {
	return fmt.Sprintf("/%s/%s", srv.Name, srv.Addr)
}
