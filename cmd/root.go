package cmd

import (
	"errors"
	"github.com/tiagoncardoso/fc-pge-stress-test/pkg/fcst"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fc-pge-stress-test",
	Short: "Stress test",
	Long:  `Full Cycle Stress Test for Go Expert MBA.`,
	RunE:  runStressTest(),
}

func runStressTest() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		params, err := getParams(cmd)
		if err != nil {
			return err
		}

		fcst.FcStress(params)

		return nil
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntP("requests", "r", 0, "Number of requests to perform")
	rootCmd.Flags().IntP("concurrency", "c", 0, "Number of concurrent requests")
	rootCmd.Flags().StringP("url", "u", "", "Target URL to test")
}

func getParams(cmd *cobra.Command) (fcst.StressTestParams, error) {
	var params fcst.StressTestParams
	requests, _ := cmd.Flags().GetInt("requests")
	concurrency, _ := cmd.Flags().GetInt("concurrency")
	url, _ := cmd.Flags().GetString("url")

	if requests == 0 {
		return params, errors.New("number of requests must be greater than 0")
	}

	if concurrency == 0 {
		return params, errors.New("number of concurrent requests must be greater than 0")
	}

	if url == "" {
		return params, errors.New("URL must be provided")
	}

	params = fcst.StressTestParams{
		Requests:    requests,
		Concurrency: concurrency,
		Url:         url,
	}

	return params, nil
}
