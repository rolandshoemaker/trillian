package failoverdriver

import (
	"database/sql"
	"database/sql/driver"
	"errors"

	"github.com/golang/glog"
)

type failoverDriver struct {
	uris       []string
	current    int
	underlying driver.Driver
}

func (fd *failoverDriver) Open(_ string) (driver.Conn, error) {
	for {
		conn, err := fd.underlying.Open(fd.uris[fd.current])
		if err == nil {
			return conn, nil
		}
		glog.Warningf("failed to connect to %q: %s", fd.uris[fd.current], err)
		fd.current = (fd.current + 1) % len(fd.uris)
	}
}

func Register(uris []string, underlying driver.Driver) error {
	if len(uris) == 0 {
		return errors.New("empty slice of database URIs")
	}
	fd := &failoverDriver{uris, 0, underlying}
	sql.Register("failoverDriver", fd)
	return nil
}
