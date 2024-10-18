package main

import (
	"context"
	"fmt"
	"github.com/davidelng/gator/internal/database"
	"strconv"
)

func handlerBrowsePosts(s *state, cmd command, currentUser database.User) error {
	var limit int32 = 2
	if len(cmd.Args) == 1 {
		limitToInt, err := strconv.ParseInt(cmd.Args[0], 10, 32)
		if err == nil {
			limit = int32(limitToInt)
		}
	}

	posts, err := s.db.GetPostsByUserID(context.Background(), database.GetPostsByUserIDParams{
		UserID: currentUser.ID,
		Limit:  limit,
	})

	if err != nil {
		return fmt.Errorf("couldn't browse posts: %w", err)
	}

	if len(posts) == 0 {
		return fmt.Errorf("no posts found")
	}

	fmt.Printf("Found %d posts for user %s:\n", len(posts), currentUser.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil
}
