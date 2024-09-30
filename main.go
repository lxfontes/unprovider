package main

//go:generate wit-bindgen-wrpc go --out-dir internal/server --package github.com/lxfontes/unprovider/internal/server wit

import (
	"context"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/lxfontes/unprovider/internal/server"
	"github.com/lxfontes/unprovider/internal/server/exports/lxfontes/unprovider/runner"
	"go.wasmcloud.dev/provider"
	wrpc "wrpc.io/go"
)

var (
	_           = (runner.Handler)(&UnproviderHandler{})
	execCommand string
)

type UnproviderHandler struct{}

func (h *UnproviderHandler) Call(ctx context.Context, input string) (*wrpc.Result[string, string], error) {
	slog.Info("Call", "input", input)

	cmd := exec.Command("/bin/sh", "-c", execCommand)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		slog.Error("Failed to create stdin pipe", "err", err)
		return wrpc.Err[string]("Failed to create stdin pipe"), nil
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, input)
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error("Command failed", "err", err)
		return wrpc.Err[string]("Command failed"), nil
	}

	return wrpc.Ok[string, string](string(out)), nil
}

func main() {
	wasmcloudprovider, err := provider.New()
	if err != nil {
		slog.Error("Failed to create provider", "err", err)
		os.Exit(1)
	}

	providerCh := make(chan error, 1)
	signalCh := make(chan os.Signal, 1)

	if rawCommand, ok := wasmcloudprovider.HostData().Config["command"]; !ok {
		slog.Error("No command provided in provider_config")
		os.Exit(1)
	} else {
		execCommand = rawCommand
	}

	handler := &UnproviderHandler{}
	// Handle RPC operations
	stopFunc, err := server.Serve(wasmcloudprovider.RPCClient, handler)
	if err != nil {
		slog.Error("Failed to create handler", "err", err)
		wasmcloudprovider.Shutdown()
		os.Exit(1)
	}

	// Handle control interface operations
	go func() {
		err := wasmcloudprovider.Start()
		providerCh <- err
	}()

	// Shutdown on SIGINT
	signal.Notify(signalCh, syscall.SIGINT)

	slog.Info("Unprovider started")
	select {
	case err = <-providerCh:
		slog.Error("Provider failed", "err", err)
		os.Exit(1)
	case <-signalCh:
		wasmcloudprovider.Shutdown()
		stopFunc()
	}
}
