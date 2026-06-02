package internal

type StdWriter struct {
	fn func(format string, a ...any)
}

func (w *StdWriter) Write(p []byte) (n int, err error) {
	w.fn("%s", string(p))
	return len(p), nil
}
