package main

import (
	"sort"
)

// top gives the keys of the top n values in a map[string]int. Tied values will
// be included and counted separately but will not be sorted consistently.
func top(n int, data map[string]int) []string {

	ranked := make(map[int][]string)
	for k, v := range data {
		ranked[v] = append(ranked[v], k)
	}

	var keys []int
	for k := range ranked {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	var langs []string
	for _, rank := range keys {
		for _, l := range ranked[rank] {
			langs = append(langs, l)
		}
	}

	if len(langs) < n {
		n = len(langs)
	}
	return langs[:n]
}

func topSlice(n int, data map[string]int) []string {

	type item struct {
		val   string
		score int
	}

	var all []item
	for k, v := range data {
		all = append(all, item{val: k, score: v})
	}

	sort.Slice(all, func(i, j int) bool {
		return all[i].score > all[j].score
	})

	var sorted []string
	for _, item := range all {
		sorted = append(sorted, item.val)
	}

	if len(sorted) < n {
		n = len(sorted)
	}
	return sorted[:n]
}

func topSliceMap(n int, data map[string]int) []string {
	var all []map[string]int
	for k, v := range data {
		all = append(all, map[string]int{k: v})
	}

	sort.Slice(all, func(i, j int) bool {
		return mapVal(all[i]) > mapVal(all[j])
	})

	var sorted []string
	for _, item := range all {
		sorted = append(sorted, mapKey(item))
	}

	if len(sorted) < n {
		n = len(sorted)
	}
	return sorted[:n]
}

func mapVal(m map[string]int) int {
	for _, v := range m {
		return v
	}
	return 0
}

func mapKey(m map[string]int) string {
	for k := range m {
		return k
	}
	return ""
}
