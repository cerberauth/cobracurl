package cobracurl

import (
	"context"
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/time/rate"
)

func TestParseRate(t *testing.T) {
	tests := []struct {
		input        string
		wantLimit    rate.Limit
		wantNil      bool
		wantErr      bool
	}{
		{"", 0, true, false},
		{"10/s", rate.Every(100 * time.Millisecond), false, false},
		{"2/m", rate.Every(30 * time.Second), false, false},
		{"1/h", rate.Every(time.Hour), false, false},
		{"2/d", rate.Every(12 * time.Hour), false, false},
		{"60", rate.Every(time.Minute), false, false}, // default unit /h: 60/h = 1/min
		{"invalid", 0, false, true},
		{"0/s", 0, false, true},
		{"-1/s", 0, false, true},
		{"10/x", 0, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			rl, err := ParseRate(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			if tt.wantNil {
				assert.Nil(t, rl)
				return
			}
			require.NotNil(t, rl)
			assert.Equal(t, tt.wantLimit, rl.Limit())
			assert.Equal(t, 1, rl.Burst())
		})
	}
}

func TestParseRateWait(t *testing.T) {
	rl, err := ParseRate("2/s")
	require.NoError(t, err)

	ctx := context.Background()

	start := time.Now()
	require.NoError(t, rl.Wait(ctx))
	require.NoError(t, rl.Wait(ctx))
	elapsed := time.Since(start)

	// Two calls at 2/s: first is immediate, second waits ~500ms
	assert.GreaterOrEqual(t, elapsed, 400*time.Millisecond)
}

func TestParseRateWaitCancelled(t *testing.T) {
	rl, err := ParseRate("1/h") // very slow
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	require.NoError(t, rl.Wait(ctx)) // first call is immediate
	err = rl.Wait(ctx)               // second call must wait ~1h, context cancels first
	assert.Error(t, err)
}

func TestBuildRateLimiter(t *testing.T) {
	t.Run("no flag returns nil", func(t *testing.T) {
		cmd := &cobra.Command{}
		rl, err := BuildRateLimiter(cmd)
		require.NoError(t, err)
		assert.Nil(t, rl)
	})

	t.Run("rate flag is parsed", func(t *testing.T) {
		cmd := &cobra.Command{}
		cmd.Flags().String("rate", "10/s", "")
		rl, err := BuildRateLimiter(cmd)
		require.NoError(t, err)
		require.NotNil(t, rl)
		assert.Equal(t, rate.Every(100*time.Millisecond), rl.Limit())
	})

	t.Run("empty rate flag returns nil", func(t *testing.T) {
		cmd := &cobra.Command{}
		cmd.Flags().String("rate", "", "")
		rl, err := BuildRateLimiter(cmd)
		require.NoError(t, err)
		assert.Nil(t, rl)
	})
}
