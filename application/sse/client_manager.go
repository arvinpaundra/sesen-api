package sse

import "sync"

var Clients = NewClientManager()

type ClientManager struct {
	mu      sync.RWMutex
	clients map[string]*Client
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		mu:      sync.RWMutex{},
		clients: make(map[string]*Client),
	}
}

func (s *ClientManager) Add(client *Client) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.clients[client.id] = client
}

func (s *ClientManager) Remove(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.clients, id)
}

func (s *ClientManager) Send(id, data string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	client, exist := s.clients[id]
	if !exist {
		return nil
	}

	return client.Send(data)
}
