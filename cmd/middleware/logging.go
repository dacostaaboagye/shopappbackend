package middleware

import (
	"context"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/Aboagye-Dacosta/shopBackend/logger"
)

// responseWriter wraps http.ResponseWriter to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode   int
	bytesWritten int64
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK, // Default status code
	}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	n, err := rw.ResponseWriter.Write(b)
	rw.bytesWritten += int64(n)
	return n, err
}

// getClientIP extracts the client's real IP address in IPv4 format
// Checks proxy headers (X-Forwarded-For, X-Real-IP) before falling back to RemoteAddr
func getClientIP(r *http.Request) string {
	// Check X-Forwarded-For header (common in reverse proxies/load balancers)
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		// X-Forwarded-For can contain multiple IPs, the first one is the original client
		ips := strings.Split(xForwardedFor, ",")
		if len(ips) > 0 {
			ip := strings.TrimSpace(ips[0])
			if ipv4 := parseIPv4(ip); ipv4 != "" {
				return ipv4
			}
		}
	}

	// Check X-Real-IP header (alternative proxy header)
	xRealIP := r.Header.Get("X-Real-IP")
	if xRealIP != "" {
		if ipv4 := parseIPv4(xRealIP); ipv4 != "" {
			return ipv4
		}
	}

	// Check CF-Connecting-IP (Cloudflare)
	cfIP := r.Header.Get("CF-Connecting-IP")
	if cfIP != "" {
		if ipv4 := parseIPv4(cfIP); ipv4 != "" {
			return ipv4
		}
	}

	// Fall back to RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		// If SplitHostPort fails, RemoteAddr might not have a port
		ip = r.RemoteAddr
	}

	if ipv4 := parseIPv4(ip); ipv4 != "" {
		return ipv4
	}

	// If no IPv4 found, return the original IP (could be IPv6 or invalid)
	return ip
}

// parseIPv4 parses an IP address and returns it if it's IPv4, otherwise returns empty string
func parseIPv4(ipStr string) string {
	ip := net.ParseIP(strings.TrimSpace(ipStr))
	if ip == nil {
		return ""
	}

	// Check if it's IPv4
	if ipv4 := ip.To4(); ipv4 != nil {
		return ipv4.String()
	}

	return ""
}

// RequestLogger logs every HTTP request with relevant context information
func RequestLogger(log *logger.AppLogger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Wrap response writer to capture status code
			rw := newResponseWriter(w)

			// Extract context values
			ctx := context.WithValue(r.Context(), logger.LOGGER_KEY, log)
			// Process the request
			next.ServeHTTP(rw, r.WithContext(ctx))

			ctx = r.Context()
			// Calculate request duration
			duration := time.Since(start)

			// Get client IP in IPv4 format
			clientIP := getClientIP(r)

			// Build log attributes
			attrs := []interface{}{
				"method", r.Method,
				"path", r.URL.Path,
				"status", rw.statusCode,
				"duration_ms", duration.Milliseconds(),
				"bytes", rw.bytesWritten,
				"client_ip", clientIP,
			}

			// Log error status codes to error log, others to info log
			// Context values are included in attrs, and CustomLogger will also extract
			// them from context if available (for consistency)
			if rw.statusCode >= http.StatusBadRequest {
				log.ErrLogger.ErrorContext(ctx, "HTTP request completed", attrs...)
			} else {
				log.InfoLogger.InfoContext(ctx, "HTTP request completed", attrs...)
			}
		})

	}
}
