package ezquake

import (
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
