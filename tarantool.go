package tarantool

import (
	"context"

	"github.com/hashicorp/vault/sdk/database/dbplugin/v5"
	"github.com/hashicorp/vault/sdk/database/helper/credsutil"
	"github.com/tarantool/go-tarantool/v2/pool"
)

var _ dbplugin.Database = &TarantoolDB{}

// TarantoolDB todo
type TarantoolDB struct {
	credsutil.CredentialsProducer
	pool *pool.ConnectionPool
}

func New() (interface{}, error) {
	db := &TarantoolDB{}
	// Wrap the plugin with middleware to sanitize errors
	dbType := dbplugin.NewDatabaseErrorSanitizerMiddleware(db, db.secretValues)
	return dbType, nil
}

// Initialize the database plugin. This is the equivalent of a constructor for the
// database object itself.
func (d *TarantoolDB) Initialize(ctx context.Context, req dbplugin.InitializeRequest) (dbplugin.InitializeResponse, error) {

}

// NewUser creates a new user within the database. This user is temporary in that it
// will exist until the TTL expires.
func (d *TarantoolDB) NewUser(ctx context.Context, req dbplugin.NewUserRequest) (dbplugin.NewUserResponse, error) {
}

// UpdateUser updates an existing user within the database.
func (d *TarantoolDB) UpdateUser(ctx context.Context, req dbplugin.UpdateUserRequest) (dbplugin.UpdateUserResponse, error) {
}

// DeleteUser from the database. This should not error if the user didn't
// exist prior to this call.
func (d *TarantoolDB) DeleteUser(ctx context.Context, req dbplugin.DeleteUserRequest) (dbplugin.DeleteUserResponse, error) {
}

// Type returns the Name for the particular database backend implementation.
// This type name is usually set as a constant within the database backend
// implementation, e.g. "mysql" for the MySQL database backend. This is used
// for things like metrics and logging. No behavior is switched on this.
func (d *TarantoolDB) Type() (string, error) {}

// Close attempts to close the underlying database connection that was
// established by the backend.
func (d *TarantoolDB) Close() error {}
