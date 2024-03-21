package requestlog

import (
	"io"
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/dinizgab/buildco-api/utils/ctx"
)

type Handler struct {
	handler http.Handler
	logger  *slog.Logger
}

func NewHandler(h http.HandlerFunc, l *slog.Logger) *Handler {
	return &Handler{
		handler: h,
		logger:  l,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	le := &logEntry{
		RequestID:         ctx.RequestID(r.Context()),
		ReceivedTime:      start,
		RequestMethod:     r.Method,
		RequestURL:        r.URL.String(),
		RequestHeaderSize: headerSize(r.Header),
		UserAgent:         r.UserAgent(),
		Referer:           r.Referer(),
		Proto:             r.Proto,
		RemoteIP:          ipFromHostPort(r.RemoteAddr),
	}

	if addr, ok := r.Context().Value(http.LocalAddrContextKey).(net.Addr); ok {
		le.ServerIP = ipFromHostPort(addr.String())
	}
	r2 := new(http.Request)
	*r2 = *r
	rcc := &readCounterCloser{r: r.Body}
	r2.Body = rcc
	w2 := &responseStats{w: w}

	h.handler.ServeHTTP(w2, r2)

	le.Latency = time.Since(start)
	if rcc.err == nil && rcc.r != nil {
		// If the handler hasn't encountered an error in the Body (like EOF),
		// then consume the rest of the Body to provide an accurate rcc.n.
		io.Copy(io.Discard, rcc)
	}
	le.RequestBodySize = rcc.n
	le.Status = w2.code
	if le.Status == 0 {
		le.Status = http.StatusOK
	}
	le.ResponseHeaderSize, le.ResponseBodySize = w2.size()
	h.logger.Info(
        "",
		"request_id", le.RequestID,
		"received_time", le.ReceivedTime,
		"method", le.RequestMethod,
		"url", le.RequestURL,
		"header_size", le.RequestHeaderSize,
		"body_size", le.RequestBodySize,
		"agent", le.UserAgent,
		"referer", le.Referer,
		"proto", le.Proto,
		"remote_ip", le.RemoteIP,
		"server_ip", le.ServerIP,
		"status", le.Status,
		"resp_header_size", le.ResponseHeaderSize,
		"resp_body_size", le.ResponseBodySize,
		"latency", le.Latency,
    )
}
