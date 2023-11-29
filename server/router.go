package server

func (s *Server) GameRoutes() {
	r := s.gin.Group("/api")
	r.POST("/games", InitStartGame(s.db).ServeHTTP)
	r.GET("/games/latest/turns/:turn_count", InitLatestTurn(s.db).ServeHTTP)
}
