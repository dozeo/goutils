// +build linux,amd64

package goutils

import (
	"fmt"
	"strings"

	"os"
	"time"

	"github.com/ErikDubbelboer/gspt"
)

func (c *Processes) SetProcTitle(title string) {
	gspt.SetProcTitle(title)
}

func (c *Processes) AppendProcTitle(title string) {
	gspt.SetProcTitle(fmt.Sprintf("%s%s", strings.Join(os.Args, " "), title))
}

func init() {
	go func() {
		time.Sleep(10 * time.Millisecond)
		Process.AppendProcTitle("")
	}()
}
