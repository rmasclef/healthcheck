package healthcheck

import (
	"database/sql"
	"errors"
	"fmt"
)

/*
	checks :
	- connection
	- close
	- ping
	- select
*/
func PostgresCheck(dbUrl string) Check {
	return func() (err error) {
		db, err := sql.Open("postgres", dbUrl)

		defer func(error) {
			connErr := db.Close()

			if err != nil && connErr != nil {
				err = errors.New(fmt.Sprintf("%s\ndb close error: %s", err.Error(), connErr.Error()))
			} else if connErr != nil {
				err = errors.New("db close error: "+connErr.Error())
			}
		}(err)

		if err != nil {
			return errors.New(fmt.Sprintf("%s: %s", "PostgreSQL health check failed during connect", err.Error()))
		}

		err = db.Ping()
		if err != nil {
			return errors.New(fmt.Sprintf("%s: %s", "PostgreSQL health check failed during ping", err.Error()))
		}

		_, err = db.Query(`SELECT VERSION()`)
		if err != nil {
			return errors.New(fmt.Sprintf("%s: %s", "PostgreSQL health check failed during select", err.Error()))
		}

		return nil
	}
}
