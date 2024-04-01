# Go http starter
This is a small project acting as a starter/template for a Go web server.

## Dependencies
These are the dependencies I went with:
- [Chi](https://github.com/go-chi/chi/v5) for http routing and easier middleware setup
- [Godotenv](https://github.com/joho/godotenv) for loading environment files
- [env](https://github.com/caarlos0/env/v10) for binding environment variables to configuration struct
- [Zerolog](https://github.com/rs/zerolog) for structured logging
- [Lumberjack](https://github.com/lumberjack.v2) for builtin log rotation
- [sqlc](https://github.com/sqlc-dev/sqlc) for database query generation
- [golang-migrate](https://github.com/golang-migrate/migrate) for database migrations
- [null](https://github.com/guregu/null) for json-compatible null types

### Chi
For http routing, I went with [Chi](https://github.com/go-chi/chi/v5). I like that it's fully compatible with native http handlers while providing a nicer api for using multiple middlewares and grouping routes together.

### Godotenv
This application uses environment variables for runtime configuration, so I decided to use [Godotenv](https://github.com/joho/godotenv) to load a set of environment variables at runtime from the `.env` file in the same directory as the application.

### Env
There are loads of configuration helper packages out there -- I decided to use [env](https://github.com/caarlos0/env/v10), but it could easily be swapped out for another if you wanted extra features like defining configuration through `yaml` or `toml`. This does the simple job of binding the environment variables defined in the `.env` file into a defined struct in the code for easier and more type-safe access.

### Zerolog
There are also lots of loggers out there, but I like [Zerolog](https://github.com/rs/zerolog) because it offers some nice http routing middlewares and context-based logging. I'm sure other loggers can do the similar things, though.

### Lumberjack
While log rotation is fairly easy to configure on the OS level, I wanted to include this so that the log file output does not get too big by default with no external configuration. [Lumberjack](https://github.com/lumberjack.v2) is an extremely convenient way to do that.

### Sqlc
[sqlc](https://github.com/sqlc-dev/sqlc) lets you define your schema and queries in plain SQL and then it generates type-safe functions and queries so you can interact with your database. I wanted to include some kind of SQL helper because writing raw queries and retrieving results gets old fast in vanilla Go.

### Golang-migrate
While not too necessary of a dependency, I do like having a tool like [golang-migrate](https://github.com/golang-migrate/migrate) to bring the local database up to date with what's defined in the schema. Since this is not a hard-dependency in the code it's pretty easy not to use, if you don't want to.

### Null
I added in [null](https://github.com/guregu/null) and configured `sqlc` to use the types from this package instead of the default `slq.Null` types because these ones marshal nicely into `json`. This is not a necessary dependency if you aren't converting database types into `json`.

## Configuration
Runtime configuration is defined in the `.env` file, but this just defines environment variables. The `.env` file is not required, so could just as easily export the environment variables before running.

Accessing configuration is done in the `config` package. The `config` struct is defined in here, along with the exported `Config` variable so you can access the bound config values.

## Routing
HTTP routing is handled in the `router.go` file. It defines a `*chi.Mux`, attaches some middlewares, and then defines the routes.

### Middleware
Most of the middleware I added is for logging: attaching a logger to the request context, putting a request ID on it, and logging requests once they're done.

I also added a recovery middleware in case there is ever a panic. In local dev, this is handled by `chi`'s recoverer, which prints a pretty version of the stack trace. For other environments, I included my own which logs the stack trace using `zerolog`.

## Logging
Since the logger is setup in the request context using the aforementioned middlewares, the easiest way to get a logger from a request is to use `log := hlog.FromRequest(r)`. The logger it returns will include the request ID and anything else we attached in the middleware. As it says in the `hlog` documentation, this is just a shortcut for `log.Ctx(r.Context())`

To utilize it in other methods which are called by request handlers, you'll have to pass the request context to those methods. Then you can use `log.Ctx(ctx)` instead, assuming you no longer have direct access to the request itself. Alternately, you can just pass the logger through the methods, but somehow this feels worse to me than just passing context through.

## Database
If you don't need a database, you can pretty easily remove `sqlc` and `golang-migrate`. Just remove `sqlc.yml`, the entire `db` package, and the `database.go` file, and then clean up whatever database calls happen to be left over.

For the sake of simplicity I'm just using a `sqlite3` database here, so if you want to switch databases you'll have to get a different database driver and probably add in some configuration (host/port/user/pass/schema...) to connect to it, as well as possibly changing some configuration values in `sqlc.yml`.

## Makefile
I've included a `Makefile` which contains some common commands for dealing with the program's lifecycle. Most of what's in there right now has to do with database stuff.
