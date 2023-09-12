package clissh

import (
	"io"

	"golang.org/x/crypto/ssh"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . SecureSession

type SecureSession interface {
	RequestPty(term string, height, width int, termModes ssh.TerminalModes) error
	SendRequest(name string, wantReply bool, payload []byte) (bool, error)
	StdinPipe() (io.WriteCloser, error)
	StdoutPipe() (io.Reader, error)
	StderrPipe() (io.Reader, error)
	Start(command string) error
	Shell() error
	Wait() error
	Close() error
}
