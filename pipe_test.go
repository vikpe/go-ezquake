package ezquake_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vikpe/go-ezquake"
	"github.com/vikpe/go-ezquake/test_files"
)

func TestPipeWriter_Write(t *testing.T) {
	username := "test"
	helpers.ResetPipe(username)

	pipeWriter := ezquake.NewPipeWriter(username)

	pipeWriter.Write("console;;")
	assert.Equal(t, "console;", helpers.ReadPipe(username))

	pipeWriter.Write(" ")
	assert.Equal(t, "console;", helpers.ReadPipe(username))

	pipeWriter.Write("lastscores")
	assert.Equal(t, "console;lastscores;", helpers.ReadPipe(username))
}

func TestPipeWriter_Clear(t *testing.T) {
	username := "test"
	helpers.ResetPipe(username)

	pipeWriter := ezquake.NewPipeWriter(username)
	pipeWriter.Write("console")
	pipeWriter.Write("qtvplay 2@foo.com:28000")
	pipeWriter.Clear()
	assert.Equal(t, "", helpers.ReadPipe(username))
}
