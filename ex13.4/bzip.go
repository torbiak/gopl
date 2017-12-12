// ex13.4 provides a bzip2 writer using the system's bzip2 binary.
package bzip

import (
	"io"
	"os/exec"
)

type Writer struct {
	cmd   exec.Cmd
	stdin io.WriteCloser
}

func NewWriter(w io.Writer) (io.WriteCloser, error) {
	cmd := exec.Cmd{Path: "/bin/bzip2", Stdout: w}
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	cmd.Start()
	if err != nil {
		return nil, err
	}
	return &Writer{cmd, stdin}, nil
}

func (w *Writer) Write(data []byte) (int, error) {
	return w.stdin.Write(data)
}

func (w *Writer) Close() error {
	pipeErr := w.stdin.Close()
	cmdErr := w.cmd.Wait()
	if pipeErr != nil {
		return pipeErr
	}
	if cmdErr != nil {
		return cmdErr
	}
	return nil
}
