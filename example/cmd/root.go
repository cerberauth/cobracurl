package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/cerberauth/cobracurl"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cobracurl-demo",
	Short: "A demo CLI for sending HTTP requests using cobracurl",
	RunE: func(cmd *cobra.Command, args []string) error {
		req, err := cobracurl.BuildRequest(cmd)
		if err != nil {
			return err
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		fmt.Println("Response:", string(body))
		return nil
	},
}

func init() {
	cobracurl.RegisterFlags(rootCmd.Flags())
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
