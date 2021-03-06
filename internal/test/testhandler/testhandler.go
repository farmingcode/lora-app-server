package testhandler

import (
	"github.com/brocaar/lora-app-server/internal/handler"
	"github.com/brocaar/lorawan"
)

// TestHandler implements a Handler for testing.
type TestHandler struct {
	SendDataUpChan            chan handler.DataUpPayload
	SendJoinNotificationChan  chan handler.JoinNotification
	SendACKNotificationChan   chan handler.ACKNotification
	SendErrorNotificationChan chan handler.ErrorNotification
	DataDownPayloadChan       chan handler.DataDownPayload
}

func NewTestHandler() *TestHandler {
	return &TestHandler{
		SendDataUpChan:            make(chan handler.DataUpPayload, 100),
		SendJoinNotificationChan:  make(chan handler.JoinNotification, 100),
		SendACKNotificationChan:   make(chan handler.ACKNotification, 100),
		SendErrorNotificationChan: make(chan handler.ErrorNotification, 100),
		DataDownPayloadChan:       make(chan handler.DataDownPayload, 100),
	}
}

func (t *TestHandler) Close() error {
	return nil
}

func (t *TestHandler) SendDataUp(appEUI lorawan.EUI64, devEUI lorawan.EUI64, payload handler.DataUpPayload) error {
	t.SendDataUpChan <- payload
	return nil
}

func (t *TestHandler) SendJoinNotification(appEUI lorawan.EUI64, devEUI lorawan.EUI64, payload handler.JoinNotification) error {
	t.SendJoinNotificationChan <- payload
	return nil
}

func (t *TestHandler) SendACKNotification(appEUI lorawan.EUI64, devEUI lorawan.EUI64, payload handler.ACKNotification) error {
	t.SendACKNotificationChan <- payload
	return nil
}

func (t *TestHandler) SendErrorNotification(appEUI lorawan.EUI64, devEUI lorawan.EUI64, payload handler.ErrorNotification) error {
	t.SendErrorNotificationChan <- payload
	return nil
}

func (t *TestHandler) DataDownChan() chan handler.DataDownPayload {
	return t.DataDownPayloadChan
}
