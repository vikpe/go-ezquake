package ezquake_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vikpe/go-ezquake"
	"github.com/vikpe/go-ezquake/test_files"
)

func TestNewClientController(t *testing.T) {
	username := "test0"
	helpers.WriteToPipe(helpers.PipePath(username), "lastscores;")

	// should reset pipe
	ezquake.NewClientController(username, "")
	assert.Equal(t, "", helpers.ReadPipe(username))
}

func TestClientController_Command(t *testing.T) {
	username := "test1"
	controller := ezquake.NewClientController(username, "")

	controller.Command("console")
	assert.Equal(t, "console;", helpers.ReadPipe(username))
}

func TestClientController_CommandWithOptions(t *testing.T) {
	t.Run("block", func(t *testing.T) {
		username := "test2"
		controller := ezquake.NewClientController(username, "")

		blockDuration := 50 * time.Millisecond

		go func() {
			controller.CommandWithOptions("+showscores", ezquake.CommandOptions{Block: blockDuration})
			controller.Command("-showscores")
		}()

		assert.Equal(t, "+showscores;", helpers.ReadPipe(username))

		time.Sleep(blockDuration)
		assert.Equal(t, "+showscores;-showscores;", helpers.ReadPipe(username))
	})

	t.Run("delay", func(t *testing.T) {
		username := "test3"
		controller := ezquake.NewClientController(username, "")

		delayDuration := 50 * time.Millisecond
		controller.CommandWithOptions("console", ezquake.CommandOptions{Delay: delayDuration})

		time.Sleep(45 * time.Millisecond)
		assert.Equal(t, "", helpers.ReadPipe(username))

		time.Sleep(5 * time.Millisecond)
		assert.Equal(t, "console;", helpers.ReadPipe(username))
	})
}
