package sqlike

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

var (
	random      = RandomSelectionHandler()
	roundRobbin = RoundRobbinSelectionHandler()
)

type SlaveSelectionHandler func(slaves []*sql.DB) *sql.DB

func (h *SlaveSelectionHandler) UnmarshalText(bytes []byte) error {
	switch text := string(bytes); strings.ToLower(text) {
	case "random":
		*h = random
		return nil

	case "round_robbin":
		*h = roundRobbin
		return nil

	default:
		return fmt.Errorf("failed to unmarshal SlaveSelectionHandler(%s)", text)
	}
}

func (h *SlaveSelectionHandler) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return fmt.Errorf("failed to unmarshal SlaveSelectionHandler) : %w", err)
	}
	if err := h.UnmarshalText([]byte(s)); err != nil {
		return err
	}
	return nil
}

func RandomSelectionHandler() SlaveSelectionHandler {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return func(slaves []*sql.DB) *sql.DB {
		i := r.Intn(len(slaves))
		return slaves[i]
	}
}

func RoundRobbinSelectionHandler() SlaveSelectionHandler {
	var cnt uint64
	var mu sync.Mutex

	return func(slaves []*sql.DB) *sql.DB {
		mu.Lock()
		defer mu.Unlock()

		i := cnt % uint64(len(slaves))

		cnt++
		return slaves[i]
	}
}
