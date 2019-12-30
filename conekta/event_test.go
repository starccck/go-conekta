package conekta

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/micro/go-micro/util/log"
)

func TestEvent(t *testing.T) {
	event := &Event{}
	event.ID = "test"
	event.Object = "event"
	event.Livemode = proto.Bool(false)
	event.CreatedAt = 12321
	event.Data.Object.Description = "test desc"
	m, err := json.Marshal(event)
	t.Logf("event: %v , err: %v \n", event, err)
	t.Logf("marshaled: %v \n", string(m))
}

func TestEventStruct(t *testing.T) {
	fp, err := os.Open("event.json")
	if err != nil {
		t.Errorf("open file error, err: %v \n", err)
		t.Fail()
	}
	defer fp.Close()

	b, err := ioutil.ReadAll(fp)
	if err != nil {
		t.Errorf("read file error: %v \n", err)
		t.Fail()
	}

	e := Event{}
	if err = json.Unmarshal(b, &e); err != nil {
		log.Errorf("unmarshal error: %v \n", err)
	}

	t.Logf("struct: %v \n", e)

	t.Logf("bytes: %v \n", string(b))
}
