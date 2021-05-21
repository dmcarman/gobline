package views_test

import (
	"gobline/tui/common/views"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestMultipleNewStatBounds(t *testing.T) {
	testCases := []struct {
		desc     string
		max      int
		value    int
		expected string
	}{
		{
			desc:     "zero value",
			max:      100,
			value:    0,
			expected: "0/100",
		},
		{
			desc:     "negative value",
			max:      100,
			value:    -50,
			expected: "0/100",
		},
		{
			desc:     "over max value",
			max:      100,
			value:    150,
			expected: "100/100",
		},
		{
			desc:     "zero max value",
			max:      0,
			value:    100,
			expected: "0/0",
		},
		{
			desc:     "negative max value",
			max:      -10,
			value:    100,
			expected: "-10/-10",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			stat := views.NewStat(tC.value, tC.max)
			if s := stat.View(); s != tC.expected {
				t.Log(s)
				t.Fatalf("failed stat test %v", tC.desc)
			}
		})
	}
}

func TestStatMsgs(t *testing.T) {
	testCases := []struct {
		desc     string
		testMsg  tea.Msg
		expected string
	}{
		{
			desc:     "update value msg",
			testMsg:  views.StatUpdateValueMsg(5),
			expected: "5/10",
		},
		{
			desc:     "update value above max msg",
			testMsg:  views.StatUpdateValueMsg(15),
			expected: "10/10",
		},
		{
			desc:     "update max msg",
			testMsg:  views.StatUpdateMaxMsg(100),
			expected: "0/100",
		},
		{
			desc:     "update max msg below value",
			testMsg:  views.StatUpdateMaxMsg(-5),
			expected: "-5/-5",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var stat tea.Model
			stat = views.NewStat(0, 10)
			stat, _ = stat.Update(tC.testMsg)
			if s := stat.View(); s != tC.expected {
				t.Log(s)
				t.Fatalf("failed stat msg test %v", tC.desc)
			}
		})
	}
}
