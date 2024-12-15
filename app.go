package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/creack/pty"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"os/exec"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	var err error
	var ptmx *os.File

	runtime.EventsOn(ctx, "client-ready", func(_ ...interface{}) {
		c := exec.Command("bash")

		ptmx, err = pty.StartWithSize(c, &pty.Winsize{
			Rows: 30,
			Cols: 100,
			X:    0,
			Y:    0,
		})

		if err != nil {
			fmt.Println("Error encountered", err)
		}

		go func() {
			reader := bufio.NewReader(ptmx)

			for {
				b := make([]byte, 8)
				reader.Read(b)
				fmt.Print(string(b))
				runtime.EventsEmit(ctx, "ptty-write", string(b))
			}
		}()
	})

	runtime.EventsOn(ctx, "craftty-write", func(data ...interface{}) {
		byteData := []byte(data[0].(string))
		ptmx.Write(byteData)
	})
}
