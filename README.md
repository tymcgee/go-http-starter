# Go http starter
This is a small project acting as a starter/template for a Go web server.

## Dependencies
These are the dependencies I went with:
- [Chi](https://github.com/go-chi/chi/v5) for http routing and easier middleware setup
- [Godotenv](https://github.com/joho/godotenv) for loading environment files
- [env](https://github.com/caarlos0/env/v10) for binding environment variables to configuration struct
- [Zerolog](https://github.com/rs/zerolog) for structured logging
- [Lumberjack](https://github.com/lumberjack.v2) for builtin log rotation

### Chi
For http routing, I went with [Chi](https://github.com/go-chi/chi/v5). I like that it's fully compatible with native http handlers while providing a nicer api for using multiple middlewares and grouping routes together.

### Godotenv
This application uses environment variables to configure it, so I decided to use [Godotenv](https://github.com/joho/godotenv) to load a set of environment variables at runtime from the `.env` file in the same directory as the application.

### Env
There are loads of configuration helper packages out there -- I decided to use [env](https://github.com/caarlos0/env/v10), but it could easily be swapped out for another if you wanted extra features like defining configuration through `yaml` or `toml`. This does the simple job of binding the environment variables defined in the `.env` file into a defined struct in the code for easier and more type-safe access.

### Zerolog
There are also lots of loggers out there, but I like [Zerolog](https://github.com/rs/zerolog) because it offers some nice http routing middlewares and context-based logging. I'm sure other loggers can do the similar things, though.

### Lumberjack
While log rotation is fairly easy to configure on the OS level, I wanted to include this so that the log file output does not get too big by default with no external configuration. [Lumberjack](https://github.com/lumberjack.v2) is an extremely convenient way to do that.

## Configuration
Runtime configuration is defined in the `.env` file, but this just defines environment variables. The `.env` file is not required, so could just as easily export the environment variables before running.

Accessing configuration is done in the `config` package. The `config` struct is defined in here, along with the exported `Config` variable so you can access the bound config values.
