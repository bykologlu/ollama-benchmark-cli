package output

import (
	"fmt"
	"os"
	"sort"

	"ollama-benchmark/internal/benchmark"
	"ollama-benchmark/internal/i18n"
)

func WriteComparison(path string, results []benchmark.BenchmarkResult) error {
	agg := Aggregate(results)
	sort.Slice(agg, func(i, j int) bool {
		return agg[i].TokenPerS > agg[j].TokenPerS
	})

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf(i18n.T("err_file_create"), err)
	}
	defer file.Close()

	fmt.Fprintf(file, "%s\t%s\t%s\t%s\t%s\n",
		i18n.T("header_model"),
		i18n.T("header_time"),
		i18n.T("header_tokens"),
		i18n.T("header_tps"),
		i18n.T("header_rank"),
	)

	for i, r := range agg {
		rank := "-"
		if i == 0 {
			rank = "🥇"
		} else if i == 1 {
			rank = "🥈"
		} else if i == 2 {
			rank = "🥉"
		}
		fmt.Fprintf(file, "%s\t%.2f\t%d\t%.1f\t%s\n", r.Model, r.AvgDuration.Seconds(), r.TotalTokens, r.TokenPerS, rank)
	}

	return nil
}
