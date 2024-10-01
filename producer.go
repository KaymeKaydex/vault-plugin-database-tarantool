package tarantool

import (
	"github.com/tarantool/go-tarantool/v2/pool"
)

type producer struct {
	pool pool.ConnectionPool
}

func (c *producer) secretValues() map[string]string {
	return map[string]string{
		c.Password: "[password]",
		c.Username: "[username]",
	}
}
