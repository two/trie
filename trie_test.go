package trie

import (
	"reflect"
	"testing"
)

func TestTrie_Insert(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  Trie
	}{
		"insert nil": {
			input: []string{},
			want:  Trie{},
		},
		"insert one": {
			input: []string{"a"},
			want: Trie{
				children: map[rune]*Trie{
					rune('a'): &Trie{
						end: true,
					},
				},
			},
		},
		"insert two": {
			input: []string{"a", "ab"},
			want: Trie{
				children: map[rune]*Trie{
					rune('a'): &Trie{
						end: true,
						children: map[rune]*Trie{
							rune('b'): &Trie{
								end: true,
							},
						},
					},
				},
			},
		},
		"insert three": {
			input: []string{"a", "ab", "ac"},
			want: Trie{
				children: map[rune]*Trie{
					rune('a'): &Trie{
						end: true,
						children: map[rune]*Trie{
							rune('b'): &Trie{
								end: true,
							},
							rune('c'): &Trie{
								end: true,
							},
						},
					},
				},
			},
		},
		"insert long word": {
			input: []string{"abc"},
			want: Trie{
				children: map[rune]*Trie{
					rune('a'): &Trie{
						children: map[rune]*Trie{
							rune('b'): &Trie{
								children: map[rune]*Trie{
									rune('c'): &Trie{
										end: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tr := NewTrie()
			for _, v := range tc.input {
				tr.Insert(v)
			}
			if !reflect.DeepEqual(tc.want, tr) {
				t.Fatalf("expected: %v, got %v", tc.want, tr)
			}
		})
	}

}

func TestTrie_Search(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"match one":     {input: "a", want: true},
		"no match one":  {input: "b", want: false},
		"match two":     {input: "aa", want: true},
		"no match two":  {input: "ac", want: false},
		"match tree":    {input: "aab", want: true},
		"no match tree": {input: "aac", want: false},
		"shorter":       {input: "ab", want: false},
		"longer":        {input: "abcd", want: false},
	}
	tr := NewTrie()
	for _, v := range []string{"a", "aa", "abc", "aaa", "aab"} {
		tr.Insert(v)
	}

	for name, tc := range tests {
		got := tr.Search(tc.input)
		t.Run(name, func(t *testing.T) {
			if tc.want != got {
				t.Fatalf("expected: %v, got %v", tc.want, got)
			}
		})
	}

}
