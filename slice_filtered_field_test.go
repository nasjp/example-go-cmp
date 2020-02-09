package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type Document struct {
	ID      int64
	Name    string
	Version int
}

func TestEqualWithFilteredField(t *testing.T) {
	opt := cmp.Comparer(func(x, y *Document) bool {
		return x.Name == y.Name
	})

	tests := []struct {
		name string
		got  *Document
		want *Document
	}{
		{"EqualName", &Document{ID: 1, Name: "金銭消費貸借契約証書", Version: 1}, &Document{Name: "金銭消費貸借契約証書"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !cmp.Equal(tt.want, tt.got, opt) {
				t.Errorf("got %v, but want %v", tt.got, tt.want)
			}
		})
	}
}

type Documents []*Document

func TestEqualSliceWithFilteredField(t *testing.T) {
	opt := cmp.Comparer(func(x, y *Document) bool {
		return x.Name == y.Name
	})

	tests := []struct {
		name string
		got  Documents
		want Documents
	}{
		{"EqualName",
			Documents{&Document{ID: 1, Name: "金銭消費貸借契約証書", Version: 1}, &Document{ID: 2, Name: "印鑑証明書"}},
			Documents{&Document{Name: "金銭消費貸借契約証書"}, &Document{Name: "印鑑証明書"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !cmp.Equal(tt.want, tt.got, opt) {
				t.Errorf("got %v, but want %v", tt.got, tt.want)
			}
		})
	}
}
