package main

import (
	"testing"
	"time"
)

func TestSafeCounter_Inc(t *testing.T) {

	tests := []struct {
		name   string
		want   int
		expect int
	}{
		{
			name: "Safe count should count correctly",
			want: 110,
		},
		{
			name: "Safe count should count correctly",
			want: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := SafeCounter{v: make(map[string]int)}
			for i := 0; i < tt.want; i++ {
				go c.Inc("somekey")
			}
			// wait for go routine
			time.Sleep(time.Second)
			if got := c.Value("somekey"); got != tt.want {

				t.Errorf("Same() = %v, want %v", got, tt.want)
			}
		})
	}
}
