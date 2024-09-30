// Generated by `wit-bindgen-wrpc-go` 0.8.0. DO NOT EDIT!
package runner

import (
	bytes "bytes"
	context "context"
	binary "encoding/binary"
	errors "errors"
	fmt "fmt"
	io "io"
	slog "log/slog"
	math "math"
	utf8 "unicode/utf8"
	wrpc "wrpc.io/go"
)

type Handler interface {
	Call(ctx__ context.Context, input string) (*wrpc.Result[string, string], error)
}

func ServeInterface(s wrpc.Server, h Handler) (stop func() error, err error) {
	stops := make([]func() error, 0, 1)
	stop = func() error {
		for _, stop := range stops {
			if err := stop(); err != nil {
				return err
			}
		}
		return nil
	}

	stop0, err := s.Serve("lxfontes:unprovider/runner", "call", func(ctx context.Context, w wrpc.IndexWriteCloser, r wrpc.IndexReadCloser) {
		defer func() {
			if err := w.Close(); err != nil {
				slog.DebugContext(ctx, "failed to close writer", "instance", "lxfontes:unprovider/runner", "name", "call", "err", err)
			}
		}()
		slog.DebugContext(ctx, "reading parameter", "i", 0)
		p0, err := func(r interface {
			io.ByteReader
			io.Reader
		}) (string, error) {
			var x uint32
			var s uint8
			for i := 0; i < 5; i++ {
				slog.Debug("reading string length byte", "i", i)
				b, err := r.ReadByte()
				if err != nil {
					if i > 0 && err == io.EOF {
						err = io.ErrUnexpectedEOF
					}
					return "", fmt.Errorf("failed to read string length byte: %w", err)
				}
				if s == 28 && b > 0x0f {
					return "", errors.New("string length overflows a 32-bit integer")
				}
				if b < 0x80 {
					x = x | uint32(b)<<s
					buf := make([]byte, x)
					slog.Debug("reading string bytes", "len", x)
					_, err = r.Read(buf)
					if err != nil {
						return "", fmt.Errorf("failed to read string bytes: %w", err)
					}
					if !utf8.Valid(buf) {
						return string(buf), errors.New("string is not valid UTF-8")
					}
					return string(buf), nil
				}
				x |= uint32(b&0x7f) << s
				s += 7
			}
			return "", errors.New("string length overflows a 32-bit integer")
		}(r)

		if err != nil {
			slog.WarnContext(ctx, "failed to read parameter", "i", 0, "instance", "lxfontes:unprovider/runner", "name", "call", "err", err)
			if err := r.Close(); err != nil {
				slog.ErrorContext(ctx, "failed to close reader", "instance", "lxfontes:unprovider/runner", "name", "call", "err", err)
			}
			return
		}
		slog.DebugContext(ctx, "calling `lxfontes:unprovider/runner.call` handler")
		r0, err := h.Call(ctx, p0)
		if cErr := r.Close(); cErr != nil {
			slog.ErrorContext(ctx, "failed to close reader", "instance", "lxfontes:unprovider/runner", "name", "call", "err", err)
		}
		if err != nil {
			slog.WarnContext(ctx, "failed to handle invocation", "instance", "lxfontes:unprovider/runner", "name", "call", "err", err)
			return
		}

		var buf bytes.Buffer
		writes := make(map[uint32]func(wrpc.IndexWriter) error, 1)

		write0, err := func(v *wrpc.Result[string, string], w interface {
			io.ByteWriter
			io.Writer
		}) (func(wrpc.IndexWriter) error, error) {
			switch {
			case v.Ok == nil && v.Err == nil:
				return nil, errors.New("both result variants cannot be nil")
			case v.Ok != nil && v.Err != nil:
				return nil, errors.New("exactly one result variant must non-nil")

			case v.Ok != nil:
				slog.Debug("writing `result::ok` status byte")
				if err := w.WriteByte(0); err != nil {
					return nil, fmt.Errorf("failed to write `result::ok` status byte: %w", err)
				}
				slog.Debug("writing `result::ok` payload")
				write, err := (func(wrpc.IndexWriter) error)(nil), func(v string, w io.Writer) (err error) {
					n := len(v)
					if n > math.MaxUint32 {
						return fmt.Errorf("string byte length of %d overflows a 32-bit integer", n)
					}
					if err = func(v int, w io.Writer) error {
						b := make([]byte, binary.MaxVarintLen32)
						i := binary.PutUvarint(b, uint64(v))
						slog.Debug("writing string byte length", "len", n)
						_, err = w.Write(b[:i])
						return err
					}(n, w); err != nil {
						return fmt.Errorf("failed to write string byte length of %d: %w", n, err)
					}
					slog.Debug("writing string bytes")
					_, err = w.Write([]byte(v))
					if err != nil {
						return fmt.Errorf("failed to write string bytes: %w", err)
					}
					return nil
				}(*v.Ok, w)
				if err != nil {
					return nil, fmt.Errorf("failed to write `result::ok` payload: %w", err)
				}
				if write != nil {
					return write, nil
				}
				return nil, nil
			default:
				slog.Debug("writing `result::err` status byte")
				if err := w.WriteByte(1); err != nil {
					return nil, fmt.Errorf("failed to write `result::err` status byte: %w", err)
				}
				slog.Debug("writing `result::err` payload")
				write, err := (func(wrpc.IndexWriter) error)(nil), func(v string, w io.Writer) (err error) {
					n := len(v)
					if n > math.MaxUint32 {
						return fmt.Errorf("string byte length of %d overflows a 32-bit integer", n)
					}
					if err = func(v int, w io.Writer) error {
						b := make([]byte, binary.MaxVarintLen32)
						i := binary.PutUvarint(b, uint64(v))
						slog.Debug("writing string byte length", "len", n)
						_, err = w.Write(b[:i])
						return err
					}(n, w); err != nil {
						return fmt.Errorf("failed to write string byte length of %d: %w", n, err)
					}
					slog.Debug("writing string bytes")
					_, err = w.Write([]byte(v))
					if err != nil {
						return fmt.Errorf("failed to write string bytes: %w", err)
					}
					return nil
				}(*v.Err, w)
				if err != nil {
					return nil, fmt.Errorf("failed to write `result::err` payload: %w", err)
				}
				if write != nil {
					return write, nil
				}
				return nil, nil
			}
		}(r0, &buf)
		if err != nil {
			slog.WarnContext(ctx, "failed to write result value", "i", 0, "lxfontes:unprovider/runner", "name", "call", "err", err)
			return
		}
		if write0 != nil {
			writes[0] = write0
		}
		slog.DebugContext(ctx, "transmitting `lxfontes:unprovider/runner.call` result")
		_, err = w.Write(buf.Bytes())
		if err != nil {
			slog.WarnContext(ctx, "failed to write result", "lxfontes:unprovider/runner", "name", "call", "err", err)
			return
		}
		if len(writes) > 0 {
			for index, write := range writes {
				w, err := w.Index(index)
				if err != nil {
					slog.ErrorContext(ctx, "failed to index writer", "index", index, "lxfontes:unprovider/runner", "name", "call", "err", err)
					return
				}
				index := index
				write := write
				go func() {
					if err := write(w); err != nil {
						slog.WarnContext(ctx, "failed to write nested result value", "index", index, "lxfontes:unprovider/runner", "name", "call", "err", err)
					}
				}()
			}
		}
	})
	if err != nil {
		return nil, fmt.Errorf("failed to serve `lxfontes:unprovider/runner.call`: %w", err)
	}
	stops = append(stops, stop0)
	return stop, nil
}
