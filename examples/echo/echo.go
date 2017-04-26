package echo

import (
	"fmt"
	"time"

	"github.com/youngtrips/anet"
)

var (
	events chan anet.Event
)

func init() {
	events = make(chan anet.Event, 65535)
}

func handleEvent(ev anet.Event) {
	sess := ev.Session
	switch ev.Type {
	case anet.EVENT_ACCEPT:
		sess.Start(events)
		break
	case anet.EVENT_CONNECT_SUCCESS:
		sess.Start(events)
		sess.Send("hello")
		break
	case anet.EVENT_MESSAGE:
		fmt.Println("recv: ", time.Now().UnixNano(), ev.Data)
		sess.Send(ev.Data)
		break
	}
}

func ServLoop() {
	proto := NewEchoProtocol()
	srv := anet.NewServer("tcp", ":9092", proto, events)
	srv.ListenAndServe()
	for {
		select {
		case ev, ok := <-events:
			if ok {
				handleEvent(ev)
			}
		}
	}
}

func CltLoop() {
	proto := NewEchoProtocol()
	session := anet.ConnectTo("tcp", "127.0.0.1:9092", proto, events, false)
	fmt.Println(session)

	for {
		select {
		case ev, ok := <-events:
			if ok {
				handleEvent(ev)
			}
			break
		}
	}
}
