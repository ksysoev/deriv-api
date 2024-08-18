package deriv

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"sync"

	"github.com/coder/websocket"
)

func newMockWSServer(cb func(ws *websocket.Conn)) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
			OriginPatterns: []string{"*"},
		})
		if err != nil {
			return
		}

		defer func() { _ = c.CloseNow() }()

		cb(c)

		c.Close(websocket.StatusNormalClosure, "")
	})

	return httptest.NewServer(handler)
}

func echoHandler(ws *websocket.Conn) {
	for {
		msgType, reader, err := ws.Reader(context.TODO())
		if err != nil {
			return
		}

		if msgType != websocket.MessageText {
			continue
		}

		buffer := bytes.NewBuffer(make([]byte, 0))

		if _, err = buffer.ReadFrom(reader); err != nil {
			return
		}

		if err = ws.Write(context.Background(), msgType, buffer.Bytes()); err != nil {
			return
		}
	}
}

func onMessageHanler(cb func(ws *websocket.Conn, msgType websocket.MessageType, msg []byte)) func(ws *websocket.Conn) {
	return func(ws *websocket.Conn) {
		var wg sync.WaitGroup
		defer wg.Wait()

		for {
			msgType, reader, err := ws.Reader(context.TODO())
			if err != nil {
				return
			}

			buffer := bytes.NewBuffer(make([]byte, 0))

			if _, err = buffer.ReadFrom(reader); err != nil {
				return
			}

			msg := buffer.Bytes()

			wg.Add(1)

			go func() {
				defer wg.Done()

				cb(ws, msgType, msg)
			}()
		}
	}
}
