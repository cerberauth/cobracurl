package cobracurl

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/time/rate"
)

// ParseRate parses a curl-style rate string (e.g. "10/s", "100/m", "1000/h", "5000/d")
// and returns a rate.Limiter with burst=1 (no bursting, steady rate).
// Returns nil if rateStr is empty.
func ParseRate(rateStr string) (*rate.Limiter, error) {
	if rateStr == "" {
		return nil, nil
	}

	parts := strings.SplitN(rateStr, "/", 2)
	n, err := strconv.ParseFloat(parts[0], 64)
	if err != nil || n <= 0 {
		return nil, fmt.Errorf("invalid rate %q: count must be a positive number", rateStr)
	}

	unit := "h" // curl default is per-hour when no unit given
	if len(parts) == 2 {
		unit = strings.ToLower(strings.TrimSpace(parts[1]))
	}

	var period time.Duration
	switch unit {
	case "s":
		period = time.Second
	case "m":
		period = time.Minute
	case "h":
		period = time.Hour
	case "d":
		period = 24 * time.Hour
	default:
		return nil, fmt.Errorf("invalid rate %q: unit must be s, m, h, or d", rateStr)
	}

	r := rate.Every(time.Duration(float64(period) / n))
	return rate.NewLimiter(r, 1), nil
}

// BuildRateLimiter creates a rate.Limiter from the --rate cobra flag.
// Returns nil if the flag is not set or empty.
func BuildRateLimiter(cmd *cobra.Command) (*rate.Limiter, error) {
	rateStr, _ := cmd.Flags().GetString("rate")
	return ParseRate(rateStr)
}
