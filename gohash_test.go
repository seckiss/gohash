package gohash

import "testing"

func TestHash80(t *testing.T) {
	// this is sourced from zocommon zutil.js hash80()
	data := map[string]string{"a": "0cc175b9c0f1b6a831c3", "short": "4f09daa9d95bcb166a30", "longer story": "50706163eb2224926fbe", "this must be even longer": "cd903054bd63eb04a6a2", "and this should exceed at least 30 characters total": "c43b8bbc98e28d51edc5"}
	for source, want := range data {
		got := Hash80(source)
		if want != got {
			t.Errorf("want != got; want=%s, got=%s", want, got)
		}
	}
}

func TestHash63(t *testing.T) {
	// this is sourced from zocommon zutil.js hash63()
	data := map[string]int64{"a": 919145239626757800, "short": 5695323626817702678, "longer story": 5796239802200368274, "this must be even longer": 5589020278079613700, "and this should exceed at least 30 characters total": 4916677060340125009}
	for source, want := range data {
		got := Hash63(source)
		if want != got {
			t.Errorf("want != got; want=%d, got=%d", want, got)
		}
	}
}
