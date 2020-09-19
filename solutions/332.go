package solutions

import (
    "sort"
)

func findItinerary(tickets [][]string) []string {
    result, paths := make([]string, 0), make(map[string][]string)

    for _, value := range tickets {
        paths[value[0]] = append(paths[value[0]], value[1])
    }

    for _, value := range paths {
        sort.Strings(value)
    }

    findItineraryRecursive("JFK", &result, paths)

    return reverse(result)
}

func findItineraryRecursive(source string, result *[]string, paths map[string][]string) {
    for len(paths[source]) > 0 {
        current := paths[source][0]
        paths[source] = paths[source][1:]

        findItineraryRecursive(current, result, paths)
    }

    *result = append(*result, source)
}

func reverse(s []string) []string {
    for i, j := 0, len(s) - 1; i < j; {
        s[i], s[j] = s[j], s[i]

        i++
        j--
    }

    return s
}
