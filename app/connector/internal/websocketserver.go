package internal

import (
	"discord/app/connector/internal/hub"
	"discord/pkg/jwtutil"
	"net/http"

	"github.com/gobwas/ws"
)

type WebSocketServer struct {
	server    *http.Server
	jwtConfig *jwtutil.Config
	clientHub *hub.Hub
}

func NewWebSocketServer(clientHub *hub.Hub, jwtConfig *jwtutil.Config) *WebSocketServer {
	s := &WebSocketServer{
		jwtConfig: jwtConfig,
		clientHub: clientHub,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handler)
	s.server = &http.Server{
		Handler: mux,
	}

	return s
}

func (s *WebSocketServer) Start(address string) {
	s.server.Addr = address

	go s.server.ListenAndServe()
}

func (s *WebSocketServer) Stop() {
	s.server.Close()
	s.clientHub.Close()
}

func (s *WebSocketServer) handler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "无权限", http.StatusUnauthorized)
		return
	}

	claims, err := jwtutil.ValidateToken(token, jwtutil.AccessToken, s.jwtConfig)
	if err != nil {
		http.Error(w, "无权限", http.StatusUnauthorized)
		return
	}

	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		http.Error(w, "Failed to upgrade to WebSocket", http.StatusInternalServerError)
		return
	}

	s.clientHub.Serve(claims.UserId, conn)

}
