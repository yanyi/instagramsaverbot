# instagramsaverbot

Instagram Saver Bot is an application I built to learn Go. It makes use of the
[Telebot](https://github.com/tucnak/telebot) package.

What Instagram Saver Bot does is to provide users with the ease of sending an
Instagram URL and the bot sends back the user the image(s) requested. All
through Telegram.

The application is live, hosted on Heroku, at [@InstagramSaverBot](t.me/InstagramSaverBot).

## Setup

There are two ways the application can be started:

1. [Using Docker](#run-using-docker)
2. [Using Go with `go build` or `go run`](#run-using-go)

## Configuration

### config.yml

Primarily, the application uses a `config.yml` file.

Navigate to [configs/](./configs). Make a copy of `config.sample.yml` to
`config.yml` and fill up the `api_token` field with your
[Telegram Bot token](https://core.telegram.org/bots#6-botfather):

```yaml
configs:
  app_env: development
  telegram_bot:
    api_token: <insert your API token here>
    api_url: https://api.telegram.org
    bot_name: Instagram Saver Bot
```

### Environment Variables

To add-on the flexibility of using environment variables such as running on
Heroku, one could pass the `-env` flag to the binary (built or using `go run`)
to use.

```sh
$ export API_TOKEN=<insert your API token here>
$ export API_URL=https://api.telegram.org
$ export BOT_NAME=your-bot-name
$ ./instagramsaverbot -env
```

## Run Using Docker

The [Dockerfile](./build/Dockerfile) makes use of
[multi-stage building](https://docs.docker.com/develop/develop-images/multistage-build/)
to keep the resulting image size small.

### Build image

You can build the image using:

```sh
$ make docker.build
...
```

Or:

```sh
$ docker build -f ./build/Dockerfile .
...
```

### Running the built image

Run the built image with
[detached mode](https://docs.docker.com/engine/reference/run/#detached--d)
using:

```sh
$ docker run -d <IMAGE ID>
...
```

The binary that was compiled into the image will run automatically. To view
logs ([with -f flag to follow](https://docs.docker.com/engine/reference/commandline/logs/)):

```sh
$ docker logs -f <CONTAINER ID>
{"event":"Loaded config file","timestamp":"2019-03-07T11:39:40.624153Z"}
{"event":"Unmarshalled config file","timestamp":"2019-03-07T11:39:40.624405Z"}
{"event":"Loaded Telegram message handlers","timestamp":"2019-03-07T11:39:42.208533Z"}
```

## Run Using Go

Build the binary using:

```sh
$ make go.build.generic
...
```

Or compile it with:

```sh
$ go build ./cmd/instagramsaverbot
...
```

Then, run the binary like:

```sh
$ ./instagramsaverbot
{"event":"Loaded config file","timestamp":"2019-03-07T11:39:40.624153Z"}
{"event":"Unmarshalled config file","timestamp":"2019-03-07T11:39:40.624405Z"}
{"event":"Loaded Telegram message handlers","timestamp":"2019-03-07T11:39:42.208533Z"}
...
```
