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
// Returns a nil limiter if rateStr is empty.
//
// When the unit is omitted the rate is per hour, as in curl, which is easy to
// misread as per second. The limiter is still returned in that case, along
// with a non-empty warning describing the assumption; warning is empty
// otherwise. Callers should surface it to the user.
func ParseRate(rateStr string) (limiter *rate.Limiter, warning string, err error) {
	if rateStr == "" {
		return nil, "", nil
	}

	parts := strings.SplitN(rateStr, "/", 2)
	n, err := strconv.ParseFloat(parts[0], 64)
	if err != nil || n <= 0 {
		return nil, "", fmt.Errorf("invalid rate %q: count must be a positive number", rateStr)
	}

	unit := "h" // curl default is per-hour when no unit given
	if len(parts) == 2 {
		unit = strings.ToLower(strings.TrimSpace(parts[1]))
	} else {
		warning = fmt.Sprintf("rate %q has no unit and is interpreted as per hour; use %q for per hour or %q for per second", rateStr, rateStr+"/h", rateStr+"/s")
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
		return nil, "", fmt.Errorf("invalid rate %q: unit must be s, m, h, or d", rateStr)
	}

	r := rate.Every(time.Duration(float64(period) / n))
	return rate.NewLimiter(r, 1), warning, nil
}

// BuildRateLimiter creates a rate.Limiter from the --rate cobra flag.
// Returns nil if the flag is not set or empty. When the rate has no unit, it
// is applied per hour and a warning is written to the command's stderr.
func BuildRateLimiter(cmd *cobra.Command) (*rate.Limiter, error) {
	rateStr, _ := cmd.Flags().GetString("rate")
	rl, warning, err := ParseRate(rateStr)
	if err != nil {
		return nil, err
	}
	if warning != "" {
		fmt.Fprintln(cmd.ErrOrStderr(), "Warning:", warning)
	}
	return rl, nil
}
