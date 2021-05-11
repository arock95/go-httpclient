package gohttp

import "sync"

var (
	mockupServer = mockServer{}
)

type mockServer struct {
	enabled     bool
	serverMutex sync.Mutex
}

func StartMockServer() {
	// allow safe concurrence (1 thread at a time, locks and unlocks)
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	mockupServer.enabled = true
}

func StopMockServer() {
	mockupServer.enabled = false
}