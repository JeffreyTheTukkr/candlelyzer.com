package databases

import (
	"os"
	"testing"
)

func Test_getConnectionString(t *testing.T) {
	// set variables
	_ = os.Setenv("PG_HOST", "hostname")
	_ = os.Setenv("PG_PORT", "100")
	_ = os.Setenv("PG_USER", "username")
	_ = os.Setenv("PG_PASS", "password")
	_ = os.Setenv("PG_NAME", "db_name")

	t.Run("verify connection string is set and formatted properly", func(t *testing.T) {
		got := getConnectionString()
		want := "postgresql://username:password@hostname:100/db_name"

		if got != want {
			t.Errorf("psql connection string is invalid, got: %s, want: %s", got, want)
		}
	})
}
