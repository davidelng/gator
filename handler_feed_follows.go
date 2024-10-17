package main

import (
	"context"
	"fmt"
	"time"

	"github.com/davidelng/gator/internal/database"
	"github.com/google/uuid"
)

func handlerGetFeedFollowsByUser(s *state, cmd command, currentUser database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("No feed follows found.")
		return nil
	}

	fmt.Printf("Found %d feeds follows:\n", len(feedFollows))
	for _, feed := range feedFollows {
		fmt.Printf("* %s\n", feed.FeedName)
	}

	return nil
}

func handlerFollowFeed(s *state, cmd command, currentUser database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>")
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't find feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    feed.ID,
		UserID:    currentUser.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't follow feed: %w", err)
	}

	fmt.Println("Feed followed successfully:")
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)

	return nil
}

func handlerUnfollowFeed(s *state, cmd command, currentUser database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>")
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't find feed: %w", err)
	}

	err = s.db.UnfollowFeed(context.Background(), database.UnfollowFeedParams{
		FeedID: feed.ID,
		UserID: currentUser.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't unfollow feed: %w", err)
	}

	return nil
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("* User:          %s\n", username)
	fmt.Printf("* Feed:          %s\n", feedname)
}
