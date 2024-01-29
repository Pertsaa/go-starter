package server

func (s *server) routes() {
	s.router.Get("/", s.handleHello())
}
