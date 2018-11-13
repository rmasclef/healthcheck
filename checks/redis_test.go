package healthcheck

import (
	"os"
	"testing"
)

const redisAddrEnvKey = "HEALTH_GO_REDIS_ADDR"
const redisProxyAddrEnvKey = "HEALTH_GO_REDIS_PROXY_ADDR"

func TestGetCheckFunctionOnRedisReturnsNothingOnSuccess(t *testing.T) {
	if os.Getenv(redisAddrEnvKey) == "" {
		t.SkipNow()
	}

	err := RedisCheck(os.Getenv(redisAddrEnvKey))()

	if err != nil {
		t.Fatalf("Redis proxy check failed:\nexpected no error\ngot '%s'.", err.Error())
	}
}

func TestGetCheckFunctionOnTwemProxyReturnsNothingOnSuccess(t *testing.T) {
	if os.Getenv(redisProxyAddrEnvKey) == "" {
		t.SkipNow()
	}

	err := RedisCheck(os.Getenv(redisProxyAddrEnvKey))()

	if err != nil {
		t.Fatalf("Redis proxy check failed:\nexpected no error\ngot '%s'.", err.Error())
	}
}

func TestGetCheckFunctionWithInvalidAddrReturnsError(t *testing.T) {
	err := RedisCheck("fake-addr")()

	expectedErrMessage := "redis ping failed: dial tcp: address fake-addr: missing port in address"

	if err == nil {
		t.Fatalf("expected: '%s'.\ngot: no error returned.", expectedErrMessage)

		return
	}

	if err.Error() != expectedErrMessage {
		t.Fatalf("expected: '%s'\ngot: '%s'.", expectedErrMessage, err.Error())
	}
}
