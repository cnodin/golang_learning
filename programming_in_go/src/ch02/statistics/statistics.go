package main

import (
	"encoding/json"
	"sort"
	"net/http"
	"fmt"
)

type statistics struct {
	numbers []float64
	mean float64
	median float64
}

func sum(numbers []float64) (total float64) {
	for _, x := range numbers {
		total += x
	}
	return total
}

func median(numbers []float64) float64 {
	middle := len(numbers) / 2
	result := numbers[middle]
	//如果是偶数，取2个中间值的平均值
	if len(numbers) % 2 == 0 {
		result = (result + numbers[middle - 1]) / 2
	}
	return result
}

func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	sort.Float64s(stats.numbers)
	stats.mean = sum(numbers) / float64(len(numbers))
	stats.median = median(numbers)

	return stats
}

const form = `
	<form action="/" method="POST">
		<label for="numbers">Numbers (comma or space-separated):</label><br />
		<input type="text" name="numbers" size="30" /><br />
		<input type="submit" value="Calculate" />
	</form>
`

func homePage(w http.ResponseWriter, r *http.Request)  {
	err := r.ParseForm()
	fmt.Fprint(w, pageTop, form)
	
}

func main() {
	http.HandleFunc("/", homePage)
}
