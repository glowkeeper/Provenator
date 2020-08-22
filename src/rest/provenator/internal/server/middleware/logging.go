package middleware

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"

	pkgLog "github.com/glowkeeper/zeusApps/src/zeusServer/src/zeus/pkg/log"
	"github.com/glowkeeper/zeusApps/src/zeusServer/src/zeus/utils"
)

// LogRecord warps a http.ResponseWriter and records the status
type LogRecord struct {
	http.ResponseWriter
	status int
}

func (r *LogRecord) Write(p []byte) (int, error) {
	return r.ResponseWriter.Write(p)
}

// WriteHeader overrides ResponseWriter.WriteHeader to keep track of the response code
func (r *LogRecord) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

// Logging middleware logs messages.
func Logging() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			record := &LogRecord{
				ResponseWriter: w,
				status:         200,
			}
			next.ServeHTTP(record, r)

			start := time.Now()
			defer func() {
				pkgLog.SLogger.Info(
					"status: ", record.status,
					", ip: ", utils.IPAddress(r),
			        ", method: ", r.Method,
			        ", path: ", r.URL.EscapedPath(),
					", uri: ", r.RequestURI,
			        ", responseTime: ", time.Since(start),
					", userAgent: ", r.UserAgent(),
				)
			}()
		})
	}
}
