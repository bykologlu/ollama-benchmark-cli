package logs

import (
	"fmt"
	"os"
	"time"

	"ollama-benchmark/internal/benchmark"
	"ollama-benchmark/internal/i18n"
	"ollama-benchmark/internal/output"
)

func AppendPerformanceLog(results []benchmark.BenchmarkResult) error {
	f, err := os.OpenFile("benchmark.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf(i18n.T("err_file_log"), err)
	}
	defer f.Close()

	aggregated := output.Aggregate(results)

	for _, r := range aggregated {
		entry := fmt.Sprintf(
			"[%s] %s: %s | %s: %.2fs | %s: %d | %s: %.2f\n",
			time.Now().Format(time.RFC3339),
			i18n.T("header_model"), r.Model,
			i18n.T("header_time"), r.AvgDuration.Seconds(),
			i18n.T("header_tokens"), r.TotalTokens,
			i18n.T("header_tps"), r.TokenPerS,
		)
		f.WriteString(entry)
	}

	return nil
}
