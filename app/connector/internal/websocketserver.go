package internal

import (
	"discord/app/connector/internal/hub"
	"discord/pkg/jwtutil"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocketServer struct {
	upgrader  *websocket.Upgrader
	server    *http.Server
	jwtConfig *jwtutil.Config
	clientHub *hub.Hub
}

func NewWebSocketServer(clientHub *hub.Hub, jwtConfig *jwtutil.Config) *WebSocketServer {
	ws := &WebSocketServer{
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		jwtConfig: jwtConfig,
		clientHub: clientHub,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", ws.handler)
	ws.server = &http.Server{
		Handler: mux,
	}

	return ws
}

func (ws *WebSocketServer) Start(address string) {
	ws.server.Addr = address

	go ws.server.ListenAndServe()
}

func (ws *WebSocketServer) Stop() {
	ws.server.Close()
	ws.clientHub.Close()
}

func (ws *WebSocketServer) handler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "无权限", http.StatusUnauthorized)
		return
	}

	claims, err := jwtutil.ValidateToken(token, jwtutil.AccessToken, ws.jwtConfig)
	if err != nil {
		http.Error(w, "无权限", http.StatusUnauthorized)
		return
	}

	conn, err := ws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade to WebSocket", http.StatusInternalServerError)
		return
	}

	hub.ServeClient(claims.UserId, conn, ws.clientHub)

}
