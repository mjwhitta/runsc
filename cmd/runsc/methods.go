package main

import (
	"slices"
	"strings"
)

type method struct {
	desc string
	val  uint
}

func stringify(m map[string]method) string {
	var keys []string
	var out []string

	for k := range m {
		keys = append(keys, k)
	}

	slices.SortStableFunc(
		keys,
		func(a string, b string) int {
			var l string = strings.ToLower(a)
			var r string = strings.ToLower(b)

			if l < r {
				return -1
			} else if l > r {
				return 1
			}

			return 0
		},
	)

	for _, key := range keys {
		out = append(out, key+"|"+m[key].desc)
	}

	return strings.Join(out, "\n")
}
