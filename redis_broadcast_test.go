package socketio

import (
	"testing"
)

func TestRedisBroadcastPublish(t *testing.T) {
	opts := &RedisAdapterOptions{
		Addr:     "127.0.0.1:6379",
		Password: "",
	}
	opts = getOptions(opts)
	rb, _ := newRedisBroadcast("/", opts)

	rb.Send("test_room", "test_event", "test_message1")
	err := rb.pub.Conn.Close()
	if err != nil {
		t.Error(err)
	}
	rb.Send("test_room", "test_event", "test_message2")
}

func TestRedisBroadcastSubscribe(t *testing.T) {
	opts := &RedisAdapterOptions{
		Addr:     "127.0.0.1:6379",
		Password: "",
	}
	opts = getOptions(opts)
	rb, _ := newRedisBroadcast("/", opts)
	go rb.dispatch()
	err := rb.sub.Conn.Close()
	if err != nil {
		t.Error(err)
	}
	rb.Send("test_room", "test_event", "test_message2")
	//select {}
}
