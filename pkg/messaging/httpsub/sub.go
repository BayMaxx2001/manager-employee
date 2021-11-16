package httpsub

import (
	"errors"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/go-chi/chi"
)

var _eventSubs map[string][]Subscriber
var _mu sync.RWMutex

type Subscriber struct {
	Name string
	C    chan []byte
}

func NewSubscriber(name string) *Subscriber {
	return &Subscriber{name, make(chan []byte)}
}

func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	event := chi.URLParam(r, "event")

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	dispatchEvent(event, b)

	w.WriteHeader(http.StatusOK)
}

func dispatchEvent(event string, data []byte) {
	_mu.RLock()
	defer _mu.RUnlock()

	subs, ok := _eventSubs[event]
	if !ok || subs == nil || len(subs) == 0 {
		return
	}

	for _, sub := range subs {
		go func(sub Subscriber) {
			sub.C <- data
		}(sub)
	}
}

func ConnectSub(s Subscriber, eventName string) error {
	_mu.Lock()
	defer _mu.Unlock()

	subs, ok := _eventSubs[eventName]
	if !ok {
		subs = make([]Subscriber, 0)
	}

	for _, sub := range subs {
		if sub.Name == s.Name {
			return errors.New("subscriber already registered")
		}
	}

	subs = append(subs, s)
	_eventSubs[eventName] = subs

	return nil
}

func DisconnectSub(p Subscriber, eventName string) error {
	_mu.Lock()
	defer _mu.Unlock()

	subs, ok := _eventSubs[eventName]
	if !ok {
		subs = make([]Subscriber, 0)
	}

	tmp := make([]Subscriber, 0)
	for _, sub := range subs {
		if sub.Name == p.Name {
			continue
		}
		tmp = append(tmp, sub)
	}

	_eventSubs[eventName] = tmp

	return nil
}

func init() {
	_eventSubs = make(map[string][]Subscriber)
}
