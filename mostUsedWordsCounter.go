package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type wordCount struct {
	word  string
	count int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run . <filename>")
		return
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("could not read file:", err)
		return
	}

	text := strings.ToLower(string(data))

	counts := make(map[string]int)
	for _, w := range strings.Fields(text) {
		w = strings.Trim(w, ".,!?;:\"'()")
		if w == "" {
			continue
		}
		counts[w]++
	}

	items := make([]wordCount, 0, len(counts))
	for w, c := range counts {
		items = append(items, wordCount{word: w, count: c})
	}

	sort.Slice(items, func(i, j int) bool {
		if items[i].count == items[j].count {
			return items[i].word < items[j].word
		}
		return items[i].count > items[j].count
	})

	limit := 10
	if len(items) < limit {
		limit = len(items)
	}
	for i := 0; i < limit; i++ {
		fmt.Printf("%d. %-10s %d\n", i+1, items[i].word, items[i].count)
	}
}

