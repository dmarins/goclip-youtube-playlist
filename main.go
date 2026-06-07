package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

const videoURLPrefix = "https://www.youtube.com/watch?v="

func fetchPlaylistURLs(ctx context.Context, svc *youtube.Service, playlistID string) ([]string, error) {
	var urls []string
	pageToken := ""

	for {
		call := svc.PlaylistItems.List([]string{"contentDetails"}).
			PlaylistId(playlistID).
			MaxResults(50)

		if pageToken != "" {
			call = call.PageToken(pageToken)
		}

		resp, err := call.Context(ctx).Do()
		if err != nil {
			return nil, fmt.Errorf("fetching playlist items: %w", err)
		}

		for _, item := range resp.Items {
			videoID := item.ContentDetails.VideoId
			urls = append(urls, videoURLPrefix+videoID)
		}

		pageToken = resp.NextPageToken
		if pageToken == "" {
			break
		}
	}

	return urls, nil
}

func main() {
	playlistID := flag.String("playlist", "", "YouTube playlist ID (required)")
	envFile := flag.String("env", ".env", "path to .env file")
	flag.Parse()

	if *playlistID == "" {
		fmt.Fprintln(os.Stderr, "error: -playlist flag is required")
		flag.Usage()
		os.Exit(1)
	}

	if err := godotenv.Load(*envFile); err != nil {
		log.Fatalf("loading env file %q: %v", *envFile, err)
	}

	apiKey := os.Getenv("YOUTUBE_API_KEY")
	if apiKey == "" {
		log.Fatal("YOUTUBE_API_KEY is not set in the env file")
	}

	ctx := context.Background()

	svc, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("creating YouTube service: %v", err)
	}

	urls, err := fetchPlaylistURLs(ctx, svc, *playlistID)
	if err != nil {
		log.Fatalf("fetching playlist URLs: %v", err)
	}

	if len(urls) == 0 {
		fmt.Fprintln(os.Stderr, "no videos found in playlist")
		os.Exit(0)
	}

	output := strings.Join(urls, "\n")

	fmt.Println(output)

	if err := clipboard.WriteAll(output); err != nil {
		log.Fatalf("copying to clipboard: %v", err)
	}

	fmt.Fprintf(os.Stderr, "\n%d URL(s) copied to clipboard.\n", len(urls))
}
