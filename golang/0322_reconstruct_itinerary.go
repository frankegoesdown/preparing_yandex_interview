package main

import (
	"sort"
	"strings"
)

func main() {

}

func findItinerary(tickets [][]string) []string {
	m := genItineraryMap(tickets)

	itinerary := []string{}
	helper(m, &itinerary, "JFK")

	return reverseStrs(itinerary)
}

func helper(m map[string][]string, route *[]string, curr string) {
	for len(m[curr]) > 0 {
		next := m[curr][0]
		m[curr] = m[curr][1:]

		helper(m, route, next)
	}

	*route = append(*route, curr)
}

func genItineraryMap(tickets [][]string) map[string][]string {
	m := make(map[string][]string)

	for _, ticket := range tickets {
		to, from := ticket[0], ticket[1]
		m[to] = append(m[to], from)
	}

	for key, conns := range m {
		sort.Slice(m[key], func(i, j int) bool {
			return strings.Compare(conns[i], conns[j]) == -1
		})
	}

	return m
}

func reverseStrs(s []string) []string {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]

		i++
		j--
	}

	return s
}
