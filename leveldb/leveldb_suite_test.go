package leveldb

import (
	"testing"

	"github.com/jcarter3/goleveldb/leveldb/testutil"
)

func TestLevelDB(t *testing.T) {
	testutil.RunSuite(t, "LevelDB Suite")
}
