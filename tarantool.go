package tarantool

import (
	"context"

	"github.com/hashicorp/vault/sdk/database/dbplugin/v5"
	"github.com/hashicorp/vault/sdk/database/helper/credsutil"
)

const (
	tarantoolTypeName = "tarantool"
)

var _ dbplugin.Database = &Tarantool{}

// Tarantool todo
type Tarantool struct {
	credsutil.CredentialsProducer
	producer producer
}

func New() (interface{}, error) {
	db := &Tarantool{}
	// Wrap the plugin with middleware to sanitize errors
	dbType := dbplugin.NewDatabaseErrorSanitizerMiddleware(db, db.producer.secretValues)

	return dbType, nil
}

// Initialize the database plugin. This is the equivalent of a constructor for the
// database object itself.
func (d *Tarantool) Initialize(ctx context.Context, req dbplugin.InitializeRequest) (dbplugin.InitializeResponse, error) {

}

// NewUser creates a new user within the database. This user is temporary in that it
// will exist until the TTL expires.
func (d *Tarantool) NewUser(ctx context.Context, req dbplugin.NewUserRequest) (dbplugin.NewUserResponse, error) {
}

// UpdateUser updates an existing user within the database.
func (d *Tarantool) UpdateUser(ctx context.Context, req dbplugin.UpdateUserRequest) (dbplugin.UpdateUserResponse, error) {
}

// DeleteUser from the database. This should not error if the user didn't
// exist prior to this call.
func (d *Tarantool) DeleteUser(ctx context.Context, req dbplugin.DeleteUserRequest) (dbplugin.DeleteUserResponse, error) {

}

func (d *Tarantool) Type() (string, error) {
	return tarantoolTypeName, nil
}

func (d *Tarantool) Close() error {
	errs := d.producer.Close()
	if len(errs) > 0 {
		return errs[0] // TODO: we need to fix
	}

	return nil
}
