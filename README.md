# 🐊 Gator

## CLI RSS Feed Aggregator written in Go

You should build a RSS feed aggregator based on your needs, and you should still use RSS to consume content, btw.

Also, I wanted an opportunity to use [SQLC](https://sqlc.dev/) and [Goose](http://pressly.github.io/goose/).

Why PostgreSQL? Elephants are cute, that's why I programmed in PHP for years.

Lastly, I love building CLI tools, and I love using Go for that.

## 📖 Requirements

- [Go](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)

## 💻 Usage

Install the program with `go install github.com/davidelng/gator`

You'll have to setup a `.gatorconfig.json` file (see `.gatorconfig.json.example`) in your home directory providing a connection string to a PostgreSQL database (remember to disable ssl if you're in localhost).

Now you can use `gator <cmd> [args...]` to run the program.

Ideally, you'll want an instance of the program running and fetching feeds, and another instance to interact with the results.

### Available commands

- `register <name>` register a new user
- `login <name>` log as specified user
- `users` list of available users
- `reset` reset database (DEBUG ONLY)
- `agg <time_between_requests>` fetch all feeds periodically
- `addfeed <name> <url>` add a new feed
- `feeds` list all saved feeds
- `follow <url>` follow a feed added by another user
- `following` show followed feeds
- `unfollow <url>` unfollow a feed
- `browse <limit>` list recent posts from feeds
