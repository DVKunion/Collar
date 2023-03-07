package terminal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/DVKunion/collar/pkg/config"
	"github.com/DVKunion/collar/pkg/log"
	"github.com/containerd/console"
	"github.com/gorilla/websocket"
)

type Terminal struct {
	ctx    context.Context
	conn   *websocket.Conn
	header http.Header
	host   config.Host
	pool   sync.Pool
}

type Message struct {
	MsgType MessageType     `json:"type"`
	Setting *MessageSetting `json:"setting,omitempty"`
	Share   *MessageShare   `json:"share,omitempty"`
	Value   string          `json:"value,omitempty"`
}

var (
	HandShakeMsg = &Message{
		MsgType: Control,
		Setting: &MessageSetting{Col: "129", Row: "29"},
	}
	PingMsg = &Message{
		MsgType: Ping,
	}
)

func (m *Message) Send() []byte {
	data, err := json.Marshal(&m)
	if err != nil {
		return []byte{}
	}
	return data
}

type MessageSetting struct {
	Col string `json:"col"`
	Row string `json:"row"`
}

type MessageShare struct {
	Status int `json:"status,omitempty"`
}

type MessageType string

const (
	Data    MessageType = "data"
	Control MessageType = "ctrl"
	Ping    MessageType = "ping"
)

func NewTerminal(ctx context.Context) *Terminal {
	return &Terminal{
		ctx: ctx,
		header: map[string][]string{
			"X-Ca-Token": {config.SingleConfig.Token},
		},
		pool: sync.Pool{
			New: func() interface{} {
				return make([]byte, 64*1024+262)
			},
		},
	}
}

func (t *Terminal) Connect(url string) error {
	dialer := websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
	}
	wss, resp, err := dialer.Dial(url, t.header)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusSwitchingProtocols {
		return err
	}
	t.conn = wss
	// first need send control msg
	err = t.conn.WriteMessage(websocket.TextMessage, HandShakeMsg.Send())
	if err != nil {
		log.Error(nil, err)
	}
	log.Infof(t.ctx, "connect success")
	return err
}

func (t *Terminal) Transfer() error {
	term, err := console.ConsoleFromFile(os.Stdout)
	if err != nil {
		return err
	}

	err = term.SetRaw()
	if err != nil {
		return err
	}
	defer term.Reset()

	errC := make(chan error, 1)
	go func() {
		b := t.pool.Get().([]byte)
		defer t.pool.Put(b)

		_, err := io.CopyBuffer(term, t, b)
		errC <- err
	}()

	go func() {
		b := t.pool.Get().([]byte)
		defer t.pool.Put(b)

		_, err := io.CopyBuffer(t, term, b)
		errC <- err
	}()

	if err := <-errC; err != nil && err != io.EOF {
		log.Info(nil, "Connect Close")
	}
	return nil
}

func (t *Terminal) Read(b []byte) (n int, err error) {
	_, message, err := t.conn.ReadMessage()
	if err != nil {
		if wsErr, ok := err.(*websocket.CloseError); ok && wsErr.Code == websocket.CloseNormalClosure {
			return 0, io.EOF
		}
		return 0, err
	}
	msg := &Message{}
	err = json.Unmarshal(message, &msg)
	if err != nil {
		return 0, err
	}
	if msg.MsgType == Ping {
		err = t.conn.WriteMessage(websocket.TextMessage, PingMsg.Send())
		if err != nil {
			return 0, err
		}
	}
	fmt.Print(msg.Value)
	copy(b, msg.Value)
	return 0, nil
}

func (t *Terminal) Write(b []byte) (n int, err error) {
	cmd := &Message{
		MsgType: Data,
		Value:   string(b),
	}
	err = t.conn.WriteMessage(websocket.TextMessage, cmd.Send())
	if err != nil {
		return 0, err
	}
	return len(b), nil
}

func (t *Terminal) Login() error {
	return nil
}
