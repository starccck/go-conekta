package conekta

import (
	"encoding/json"
	"testing"

	"github.com/golang/protobuf/proto"
)

func TestEvent(t *testing.T) {
	event := &Event{}
	event.ID = "test"
	event.Object = "event"
	event.Livemode = proto.Bool(false)
	event.CreatedAt = 12321
	event.Data.Description = "test desc"
	m, err := json.Marshal(event)
	t.Logf("event: %v , err: %v \n", event, err)
	t.Logf("marshaled: %v \n", string(m))
}
