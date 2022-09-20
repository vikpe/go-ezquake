package ezquake

import (
	"time"

	"github.com/vikpe/go-ezquake/internal/pkg/proc"
)

type ClientController struct {
	Process *proc.ProcessController
	pipe    *PipeWriter
}

func NewClientController(username string, binPath string) *ClientController {
	writer := NewPipeWriter(username)
	writer.Clear()

	return &ClientController{
		Process: proc.NewProcessController(binPath),
		pipe:    writer,
	}
}

func (c *ClientController) Command(cmd string) {
	if !c.Process.IsStarted() {
		return
	}

	c.pipe.Write(cmd)
}

func (c *ClientController) CommandWithOptions(cmd string, options CommandOptions) {
	cmdFunc := func() {
		c.Command(cmd)

		if options.Timeout > 0 {
			time.Sleep(options.Timeout)
		}
	}

	if options.Delay > 0 {
		time.AfterFunc(options.Delay, cmdFunc)
	} else {
		cmdFunc()
	}
}
