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
	return &ClientController{
		Process: proc.NewProcessController(binPath),
		pipe:    NewPipeWriter(username),
	}
}

func (c *ClientController) Command(cmd string) {
	if !c.Process.IsStarted() {
		return
	}

	c.pipe.Write(cmd)
}

func (c *ClientController) CommandWithTimeout(cmd string, duration time.Duration) {
	c.Command(cmd)
	time.Sleep(duration)
}

func (c *ClientController) CommandAfterDelay(cmd string, delay time.Duration) {
	time.AfterFunc(delay, func() {
		c.Command(cmd)
	})
}
