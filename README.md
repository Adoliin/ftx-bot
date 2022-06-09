# FTX-bot
A bot, written in Go, that periodically retrieves trading volume numbers from FTX database, and
persists them in a postgreSQL database. Additionaly, the bot exposes a rest API
that you can use to interact with bot.

## Getting Started
### Prerequisites
#### Environment variables
- `BOT_MARKETS`<br>
The market names that the bot will be fetching (a comma seperated list)
> eg: AAPL-0624,AAVE-PERP,AGLD-PERP

- `BOT_FREQUENCY`<br>
The bot poll frequency at which it well fetch data from FTX exchange APIs (in seconds)
> eg: 15

- `DB_HOST`<br>
Postgres database hostname

- `DB_USER`<br>
Postgres username


- `DB_PWD`<br>
Postgres password

- `DB_NAME`<br>
Postgres database name

- `DB_PORT`<br>
Postgres port

#### How to expose these environment variables?
Depending on the way you intend to deploy this bot there are many ways you can
expose these variables.
For simplicity sake when developing you just need to set a `.env` file like
this for example:

`./ftx-bot/.env`
```
BOT_MARKETS=AAPL-0624,AAVE-PERP,AGLD-PERP
BOT_FREQUENCY=15

DB_HOST=localhost
DB_USER=postgres
DB_PWD=postgres
DB_NAME=postgres
DB_PORT=5432
```

### Interacting with the bot
Using the REST API you can interact with the bot.

#### Available enpoints
- **GET** `/bot/frequency`<br>
Get the poll frequency at which the bot is running (set by `BOT_FREQUENCY` env variable)
> eg. response
```json
{
    "bot_frequency_seconds": 15
}
```
- **PUT** `/bot/frequency/<new_frequency>`<br>
Change the poll frequency at which the bot is running (set by `BOT_FREQUENCY` env variable)
> Where \<new_frequency> is the new bot frequency (in seconds)

## Running it
To run the bot you need can either use `docker` and `docker-compose` to run it in
a container. Or you can provide your own `postgres` database instance and run
the bot as a binary.

> However the following methods are my _way_ of running the bot.

### Development
To start developing you need to set the environment variables, and start a
postgres instance using docker or your own instance.

1. Navigate to root directory
`cd path/to/ftx-bot`
2. Start postgres
`docker-compose up -d`
3. Start the application
`make watch-dev`

PS: You may need to install `reflex` (used for hotreloading the server)
by typing this command:<br>
`go install github.com/cespare/reflex@latest`

### Production
Using `docker-compose` you must uncomment some lines in docker-compose to let
the app build too so you must change your docker-compose.yml file like this for
example:

`./ftx-bot/docker-compose.yml`
```yml
version: "3.9"
services:
  db:
    image: postgres:14-alpine
    restart: always
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
    ports:
      - '5432:5432'
    volumes: 
      - db:/var/lib/postgresql/data
  ftx-bot:
    depends_on:
    - db
    build: .
    ports:
      - 3000:3000
    environment:
      BOT_MARKETS: AAPL-0624,AAVE-PERP,AGLD-PERP
      BOT_FREQUENCY: 15
      DB_USER: postgres
      DB_PWD: postgres
      DB_NAME: postgres
      DB_PORT: 5432
      DB_HOST: db

volumes:
  db:
```

And just run:
```sh
docker-compose up -d
```
and the bot will just start.<br>
You can also change the environment variable as you wish from the
`docker-compose.yml`

## License
This project is licensed under the MIT License - see the
[LICENSE.md](LICENSE.md) file for details.
