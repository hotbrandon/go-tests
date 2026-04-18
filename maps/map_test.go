package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	names := Dictionary{"mike": "agent 001", "jim": "agent 002"}

	t.Run("search existing key", func(t *testing.T) {
		got, err := names.Search("mike")
		want := "agent 001"

		if err != nil {
			t.Fatal(err)
		}
		if got != want {
			t.Errorf("want %s, got %s", want, got)
		}
	})

	t.Run("search non-existent key", func(t *testing.T) {
		got, err := names.Search("jack")
		want := ""

		if err == nil {
			t.Error("expecting errors but got none")
		}
		if got != want {
			t.Errorf("want %s, got %s", want, got)
		}
	})
}

func TestAdd(t *testing.T) {
	names := Dictionary{}
	names.Add("brandon", "agent 003")
	got, _ := names.Search("brandon")
	want := "agent 003"

	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}

}
