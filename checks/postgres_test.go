package healthcheck

import (
	"os"
	"strings"
	"testing"

	_ "github.com/lib/pq"
)

const dbAddrEnvKey = "HEALTH_GO_POSTGRES_ADDR"

func TestPostgresCheckWithGoodDbAddrReturnsNothing(t *testing.T) {
	if os.Getenv(dbAddrEnvKey) == "" {
		t.SkipNow()
	}

	err := PostgresCheck(os.Getenv(dbAddrEnvKey))()

	if err != nil {
		t.Fatalf("PostgreSQL check failed:\nexpected no error\ngot '%s'.", err.Error())
	}
}

func TestPostgresCheckWithInvalidDbAddrReturnsPingError(t *testing.T) {
	err := PostgresCheck("postgres://fake-addr")()

	expectedErrMessage := "PostgreSQL health check failed during ping: dial tcp: lookup fake-addr"

	if err == nil {
		t.Fatalf("expected: '%s'.\ngot: no error returned.", expectedErrMessage)

		return
	}

	if !strings.Contains(err.Error(), expectedErrMessage) {
		t.Fatalf("expected: '%s' to contain '%s'.", err.Error(), expectedErrMessage)
	}
}
