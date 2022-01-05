package core

import (
	"context"
	"gh-labeler/pkg/io"
	"gh-labeler/pkg/types"
	"os"

	"github.com/ghodss/yaml"

	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"

	"github.com/spf13/cobra"
)

func ScanRepoForLabels(cmd *cobra.Command, args []string) error {
	repo := cmd.Flag("repo").Value.String()
	org := cmd.Flag("org").Value.String()
	filePath := cmd.Flag("outputFilePath").Value.String()

	accessTokenEnv := os.Getenv("ACCESS_TOKEN")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessTokenEnv},
	)
	c := oauth2.NewClient(ctx, ts)

	client := github.NewClient(c)
	var labels []*types.Label
	pagination := &github.ListOptions{
		PerPage: 50,
		Page:    1,
	}

	for {
		upstreamLabels, resp, err := client.Issues.ListLabels(ctx, org, repo, pagination)
		if err != nil {
			return err
		} else {
			for _, upstreamLabel := range upstreamLabels {
				labels = append(labels, types.From(upstreamLabel))
			}
		}
		if resp.NextPage == 0 {
			break
		} else {
			pagination.Page = resp.NextPage
		}
	}

	yamlData, err := yaml.Marshal(labels)
	if err != nil {
		return err
	}

	return io.WriteFile(filePath, string(yamlData))
}
