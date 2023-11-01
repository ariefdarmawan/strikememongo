package monitor

import (
	"os"
	"testing"
	"time"

	"github.com/ariefdarmawan/strikememongo/memongohelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func runSleep() *os.Process {
	cmd := memongohelper.RunCommand("sleep", "10")

	err := cmd.Start()
	if err != nil {
		panic(err)
	}

	return cmd.Process
}

func TestMonitor(t *testing.T) {
	parent := runSleep()
	child := runSleep()

	// Start the monitor
	_, err := RunMonitor(parent.Pid, child.Pid)
	require.NoError(t, err)

	// Kill the parent
	require.NoError(t, parent.Kill())

	// Child should die within 3 seconds
	startWait := time.Now()
	_, err = child.Wait()
	require.NoError(t, err)

	assert.True(t, time.Since(startWait).Seconds() < 3)
}
