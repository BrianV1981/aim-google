package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/BrianV1981/aim-google/internal/googleapi"
	"github.com/BrianV1981/aim-google/internal/outfmt"
	"github.com/BrianV1981/aim-google/internal/ui"
)

var newYouTubeService = googleapi.NewYouTube

type YouTubeCmd struct {
	Stats  YouTubeStatsCmd  `cmd:"" name:"stats" help:"Get channel stats"`
	Videos YouTubeVideosCmd `cmd:"" name:"videos" aliases:"search" help:"Search or list channel videos"`
	Update YouTubeUpdateCmd `cmd:"" name:"update" help:"Update video title, description, or tags"`
}

type YouTubeStatsCmd struct {
	ChannelID string `arg:"" optional:"" help:"Channel ID (leave empty for your own channel)"`
}

func (c *YouTubeStatsCmd) Run(ctx context.Context, f *RootFlags) error {
	u := ui.FromContext(ctx)
	account, err := requireAccount(f)
	if err != nil {
		return err
	}

	svc, err := newYouTubeService(ctx, account)
	if err != nil {
		return err
	}

	call := svc.Channels.List([]string{"snippet", "statistics"})
	if c.ChannelID != "" {
		call = call.Id(c.ChannelID)
	} else {
		call = call.Mine(true)
	}

	resp, err := call.Context(ctx).Do()
	if err != nil {
		return err
	}

	if outfmt.IsJSON(ctx) {
		return outfmt.WriteJSON(ctx, os.Stdout, resp)
	}

	w, flush := tableWriter(ctx)
	defer flush()
	fmt.Fprintln(w, "ID\tTITLE\tSUBSCRIBERS\tVIEWS\tVIDEOS")
	for _, ch := range resp.Items {
		fmt.Fprintf(w, "%s\t%s\t%d\t%d\t%d\n", ch.Id, ch.Snippet.Title, ch.Statistics.SubscriberCount, ch.Statistics.ViewCount, ch.Statistics.VideoCount)
	}

	if len(resp.Items) == 0 {
		u.Err().Println("No channel found")
	}
	return nil
}

type YouTubeVideosCmd struct {
	ChannelID  string `arg:"" optional:"" help:"Channel ID (leave empty for your own channel)"`
	MaxResults int64  `help:"Maximum number of results" default:"10"`
}

func (c *YouTubeVideosCmd) Run(ctx context.Context, f *RootFlags) error {
	u := ui.FromContext(ctx)
	account, err := requireAccount(f)
	if err != nil {
		return err
	}

	svc, err := newYouTubeService(ctx, account)
	if err != nil {
		return err
	}

	call := svc.Search.List([]string{"snippet"}).Type("video").MaxResults(c.MaxResults)
	if c.ChannelID != "" {
		call = call.ChannelId(c.ChannelID)
	} else {
		call = call.ForMine(true)
	}

	resp, err := call.Context(ctx).Do()
	if err != nil {
		return err
	}

	if outfmt.IsJSON(ctx) {
		return outfmt.WriteJSON(ctx, os.Stdout, resp)
	}

	w, flush := tableWriter(ctx)
	defer flush()
	fmt.Fprintln(w, "VIDEO ID\tTITLE\tPUBLISHED AT")
	for _, item := range resp.Items {
		fmt.Fprintf(w, "%s\t%s\t%s\n", item.Id.VideoId, item.Snippet.Title, item.Snippet.PublishedAt)
	}

	if len(resp.Items) == 0 {
		u.Err().Println("No videos found")
	}
	return nil
}

type YouTubeUpdateCmd struct {
	VideoID     string `arg:"" help:"Video ID to update"`
	Title       string `help:"New title for the video"`
	Description string `help:"New description for the video"`
	Tags        string `help:"Comma-separated tags"`
}

func (c *YouTubeUpdateCmd) Run(ctx context.Context, f *RootFlags) error {
	u := ui.FromContext(ctx)
	account, err := requireAccount(f)
	if err != nil {
		return err
	}

	svc, err := newYouTubeService(ctx, account)
	if err != nil {
		return err
	}

	// Fetch existing video snippet first (categoryId is required for Update)
	listCall := svc.Videos.List([]string{"snippet"}).Id(c.VideoID)
	resp, err := listCall.Context(ctx).Do()
	if err != nil {
		return err
	}
	if len(resp.Items) == 0 {
		return fmt.Errorf("video %q not found", c.VideoID)
	}

	video := resp.Items[0]

	updated := false
	if c.Title != "" {
		video.Snippet.Title = c.Title
		updated = true
	}
	if c.Description != "" {
		video.Snippet.Description = c.Description
		updated = true
	}
	if c.Tags != "" {
		video.Snippet.Tags = strings.Split(c.Tags, ",")
		for i := range video.Snippet.Tags {
			video.Snippet.Tags[i] = strings.TrimSpace(video.Snippet.Tags[i])
		}
		updated = true
	}

	if !updated {
		u.Err().Println("No updates provided. Use --title, --description, or --tags")
		return nil
	}

	updateCall := svc.Videos.Update([]string{"snippet"}, video)
	updateResp, err := updateCall.Context(ctx).Do()
	if err != nil {
		return err
	}

	if outfmt.IsJSON(ctx) {
		return outfmt.WriteJSON(ctx, os.Stdout, updateResp)
	}

	u.Out().Printf("Updated video: %s\n", updateResp.Id)
	return nil
}
