package mig



import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/hashicorp/go-multierror"

	"github.com/golang-migrate/migrate/v4/database"
	iurl "github.com/golang-migrate/migrate/v4/internal/url"
	"github.com/golang-migrate/migrate/v4/source"
)

// DefaultPrefetchMigrations sets the number of migrations to pre-read
// from the source. This is helpful if the source is remote, but has little
// effect for a local source (i.e. file system).
// Please note that this setting has a major impact on the memory usage,
// since each pre-read migration is buffered in memory. See DefaultBufferSize.
var DefaultPrefetchMigrations = uint(10)

// DefaultLockTimeout sets the max time a database driver has to acquire a lock.
var DefaultLockTimeout = 15 * time.Second

var (
	ErrNoChange       = errors.New("no change")
	ErrNilVersion     = errors.New("no migration")
	ErrInvalidVersion = errors.New("version must be >= -1")
	ErrLocked         = errors.New("database locked")
	ErrLockTimeout    = errors.New("timeout: can't acquire database lock")
)

// ErrShortLimit is an error returned when not enough migrations
// can be returned by a source for a given limit.
type ErrShortLimit struct {
	Short uint
}

// Error implements the error interface.
func (e ErrShortLimit) Error() string {
	return fmt.Sprintf("limit %v short", e.Short)
}

type ErrDirty struct {
	Version int
}

func (e ErrDirty) Error() string {
	return fmt.Sprintf("Dirty database version %v. Fix and force version.", e.Version)
}

type Migrate struct {
	sourceName   string
	sourceDrv    source.Driver
	databaseName string
	databaseDrv  database.Driver

	// Log accepts a Logger interface
	Log Logger

	// GracefulStop accepts `true` and will stop executing migrations
	// as soon as possible at a safe break point, so that the database
	// is not corrupted.
	GracefulStop chan bool
	isLockedMu   *sync.Mutex

	isGracefulStop bool
	isLocked       bool

	// PrefetchMigrations defaults to DefaultPrefetchMigrations,
	// but can be set per Migrate instance.
	PrefetchMigrations uint

	// LockTimeout defaults to DefaultLockTimeout,
	// but can be set per Migrate instance.
	LockTimeout time.Duration
}


func (m *Migrate) Up() error {
	if err := m.lock(); err != nil {
		return err
	}

	curVersion, dirty, err := m.databaseDrv.Version()
	if err != nil {
		return m.unlockErr(err)
	}

	if dirty {
		return m.unlockErr(ErrDirty{curVersion})
	}

	ret := make(chan interface{}, m.PrefetchMigrations)

	go m.readUp(curVersion, -1, ret)
	return m.unlockErr(m.runMigrations(ret))
}