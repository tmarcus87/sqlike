package sqlike

import (
	"database/sql"
	"fmt"
	"strings"
)

type EngineOption struct {
	Driver                string                `json:"driver"          yaml:"driver"`
	Address               string                `json:"address"         yaml:"address"`
	SlaveAddresses        []string              `json:"slave_addresses" yaml:"slave_addresses"`
	SlaveSelectionHandler SlaveSelectionHandler `json:"slave_selection" yaml:"slave_selection"`
}

type Option func(o *EngineOption)

func FromHostAndPort(driver, host string, port uint16, username, password string) Option {
	return func(o *EngineOption) {
		o.Driver = driver
		o.Address = fmt.Sprintf("%s:%s@tcp(%s:%d)", username, password, host, port)
	}
}

func WithSlaveByHostAndPort(host string, port uint16, username, password string) Option {
	return func(o *EngineOption) {
		o.SlaveAddresses = append(o.SlaveAddresses, fmt.Sprintf("%s:%s@tcp(%s:%d)", username, password, host, port))
	}
}

func WithSlaveSelectionHandler(handler SlaveSelectionHandler) Option {
	return func(o *EngineOption) {
		o.SlaveSelectionHandler = handler
	}
}

func NewEngine(opts ...Option) (Engine, error) {
	o := EngineOption{}
	for _, opt := range opts {
		opt(&o)
	}

	return NewEngineFromOption(&o)
}

func NewEngineFromOption(o *EngineOption) (Engine, error) {
	db, err := sql.Open(o.Driver, o.Address)
	if err != nil {
		return nil, fmt.Errorf("failed to open connection : %w", err)
	}

	dbs := make([]*sql.DB, 0)
	if len(o.SlaveAddresses) == 0 {
		dbs = []*sql.DB{db}
	} else {
		for _, address := range o.SlaveAddresses {
			db, err := sql.Open(o.Driver, address)
			if err != nil {
				return nil, fmt.Errorf("failed to open slave connection : %w", err)
			}
			dbs = append(dbs, db)
		}
	}
	return &basicEngine{
		dialect: strings.ToLower(o.Driver),
		master:  db,
		slaves:  dbs,
	}, nil
}
