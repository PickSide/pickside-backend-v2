package ws

import (
	"encoding/json"
	"errors"
	"fmt"
	"pickside/service/data"
	"pickside/service/util"
	"regexp"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
	mu    sync.Mutex
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) HandleWS(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client: ", ws.RemoteAddr())

	s.mu.Lock()
	if _, ok := s.conns[ws]; !ok {
		s.conns = make(map[*websocket.Conn]bool)
	}
	s.conns[ws] = true
	s.mu.Unlock()

	s.readLoop(ws)
}

type Message struct {
	Event   string          `json:"event,required"`
	Content json.RawMessage `json:"content,omitempty"`
}

func (m *Message) ValidateEvent() error {
	pattern := `^[a-zA-Z]+:[a-zA-Z]+$`
	matched, err := regexp.MatchString(pattern, m.Event)
	if err != nil {
		return err
	}
	if !matched {
		return fmt.Errorf("event does not match required format (string:string)")
	}
	return nil
}

func (s *Server) readLoop(ws *websocket.Conn) {
	for {
		messageType, n, err := ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("Error: %v\n", err)
			}
			break
		}

		var msg Message
		if err := json.Unmarshal(n, &msg); err != nil {
			fmt.Printf("Error unmarshalling message: %v\n", err)
			continue
		}

		if err := msg.ValidateEvent(); err != nil {
			fmt.Println("Validation error:", err)
			continue
		}

		r, err := DelegateMessage(&msg)
		if err != nil {
			fmt.Println("Delegation error:", err)
			continue
		}

		resp, err := json.Marshal(r)
		if err != nil {
			fmt.Println("Delegation error:", err)
			continue
		}

		if err := ws.WriteMessage(messageType, []byte(resp)); err != nil {
			fmt.Println("write error:", err)
			break
		}
	}

	s.mu.Lock()
	delete(s.conns, ws)
	s.mu.Unlock()
	ws.Close()
}

var AllowedNamespaces = []string{
	"chatrooms",
	"groups",
	"notifications",
	"users",
}

func DelegateMessage(msg *Message) (any, error) {
	s := strings.Split(msg.Event, ":")
	namespace := s[0]
	action := s[1]

	if ok := util.Contains(AllowedNamespaces, namespace); !ok {
		return nil, errors.New("namespace is not allowed")
	}

	if namespace == "chatrooms" {
	}
	if namespace == "groups" {
	}
	if namespace == "notifications" {
		if action == "user" {
			notifications, err := data.GetUserNotifications(1)
			if err != nil {
				return nil, err
			}
			return notifications, nil
		}
		if action == "global" {
			globalNotifications, err := data.GetGlobalNotifications()
			if err != nil {
				return nil, err
			}
			return globalNotifications, nil
		}
	}
	if namespace == "users" {
	}
	return nil, nil
}
