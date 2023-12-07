# go-healthcheck
A simple website health checker written in Go


## Installation

1. Make sure you have Go installed on your system.

2. Clone the repository:

    ```bash
    git clone https://github.com/AliSinaDevelo/go-healthcheck.git
    ```

3. Navigate to the project directory:

    ```bash
    cd healthchecker
    ```

4. Build the executable:

    ```bash
    go build
    ```

## Usage

```bash
./healthchecker --domain example.com --port 80
```

### Flags

- `--domain` or `-d`: Domain name to check if it's running (required).
- `--port` or `-p`: Port number (optional, default is 80).

## Example

```bash
./healthchecker --domain example.com --port 80
```

This will output the status of the specified domain and port, indicating whether it is reachable.

## Contributing

If you have suggestions or find issues, please feel free to open an [issue](https://github.com/AliSinaDevelo/go-healthcheck/issues) or submit a [pull request](https://github.com/AliSinaDevelo/go-healthcheck/pulls).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.


