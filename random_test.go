package strikememongo

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomDatabase(t *testing.T) {
	s := RandomDatabase()

	assert.Len(t, s, DBNameLen)

	dbNameRunes := []rune(DBNameChars)
	for _, c := range s {
		assert.Contains(t, dbNameRunes, c)
	}
}

func TestRandomDatabaseEntropy(t *testing.T) {
	seen := map[string]bool{}

	for i := 0; i < 1000; i++ {
		s := RandomDatabase()
		assert.False(t, seen[s])

		seen[s] = true
	}
}

func TestDbDir(t *testing.T) {
	dbDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Error(err.Error())
	}
	t.Logf("dir:%s", dbDir)
}
