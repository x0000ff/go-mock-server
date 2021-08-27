# Go Mock Server

- 👶 Easy to use
- 🔄 Automatic reloading when you change your mock files
- 📦 JSON support out of the box
- 📖 Easy configuration (see [config.toml](config.toml))

## Run

```shell
$ make build & make run 
```

## Auto reload

1. Install [Modd](https://github.com/cortesi/modd)
    - For MacOSX the easier way is to install with [Homebrew](https://formulae.brew.sh/): `$ brew install modd`

2. Just run it:

```shell
$ modd
```

3. Navigate to [http://localhost:8080/me](http://localhost:8080/users)
   or [http://localhost:8080/stats](http://localhost:8080/stats)

4. It'll
    - Rebuild the mock server when you change any `.go` file
    - Relaunch when you change any `.json` file

## How to add mocks?

- Add your mock file to the `mocks/` folder
- Edit `config.toml` like:

```go
[[routes]]
path = "/users"
filename = "users.json"
```

