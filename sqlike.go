package sqlike

import (
	"database/sql"
	"fmt"
	"strings"
)

type ConnectionInfo struct {
	Username string            `json:"username" yaml:"username"`
	Password string            `json:"password" yaml:"password"`
	Host     string            `json:"host"     yaml:"host"`
	Port     uint16            `json:"port"     yaml:"port"`
	Options  map[string]string `json:"options"  yaml:"options"`
}

func (c *ConnectionInfo) dataSourceName(database string) string {
	opts := make([]string, 0)
	for k, v := range c.Options {
		opts = append(opts, fmt.Sprintf("%s=%s", k, v))
	}
	opt := strings.Join(opts, "&")

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?%s",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		database,
		opt)
}

type EngineOption struct {
	Driver                string                `json:"driver"          yaml:"driver"`
	Database              string                `json:"database"        yaml:"database"`
	Master                ConnectionInfo        `json:"master"          yaml:"address"`
	Slaves                []ConnectionInfo      `json:"slaves"          yaml:"slaves"`
	SlaveSelectionHandler SlaveSelectionHandler `json:"slave_selection" yaml:"slave_selection"`
}

type Option func(o *EngineOption)

func FromHostAndPort(driver, host string, port uint16, username, password, database string) Option {
	return func(o *EngineOption) {
		o.Driver = driver
		o.Database = database
		o.Master = ConnectionInfo{Username: username, Password: password, Host: host, Port: port}
	}
}

func WithSlaveByHostAndPort(host string, port uint16, username, password string) Option {
	return func(o *EngineOption) {
		o.Slaves = append(o.Slaves, ConnectionInfo{Username: username, Password: password, Host: host, Port: port})
	}
}

func WithSlaveSelectionHandler(handler SlaveSelectionHandler) Option {
	return func(o *EngineOption) {
		o.SlaveSelectionHandler = handler
	}
}

func NewEngine(opts ...Option) (Engine, error) {
	o := EngineOption{
		Slaves:                make([]ConnectionInfo, 0),
		SlaveSelectionHandler: RoundRobbinSelectionHandler(),
	}
	for _, opt := range opts {
		opt(&o)
	}

	return NewEngineFromOption(&o)
}

func NewEngineFromOption(o *EngineOption) (Engine, error) {
	db, err :=
		sql.Open(o.Driver, o.Master.dataSourceName(o.Database))
	if err != nil {
		return nil, fmt.Errorf("failed to open connection : %w", err)
	}

	dbs := make([]*sql.DB, 0)
	for _, slave := range o.Slaves {
		db, err := sql.Open(o.Driver, slave.dataSourceName(o.Database))
		if err != nil {
			return nil, fmt.Errorf("failed to open slave connection : %w", err)
		}
		dbs = append(dbs, db)
	}
	return &basicEngine{
		dialect:      strings.ToLower(o.Driver),
		master:       db,
		slaves:       dbs,
		slaveHandler: o.SlaveSelectionHandler,
	}, nil
}
