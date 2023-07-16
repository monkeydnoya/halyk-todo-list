package fiber

import "fmt"

func (s *Server) StartServer() {
	listen := s.ApiConfig.Host + ":" + s.ApiConfig.Port
	fmt.Println(listen)
	if err := s.App.Listen(listen); err != nil {
		panic(err)
	}
}
