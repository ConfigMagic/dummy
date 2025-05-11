# dummy

## Setup and Running the Server

### Prerequisites
Before you proceed, make sure you have Go installed and your environment is properly set up to build Go projects.

### Step 1: Install the Binaries

Run the following commands to build and install the `dummy` and `dummy-admin` binaries:

1. Go to the `deploy` directory:
    ```bash
    cd deploy
    ```

2. Install the required tools:
    ```bash
    sudo make install
    ```

    You should now see:
    ```
    ✅ Installed dummy and dummy-admin to /usr/local/bin.
    ```

### Step 2: Run the Server with `dummy-admin`

Once the installation is complete, you can use the `dummy-admin` tool to manage the server. To start the server:

1. Run the `dummy-admin` tool:
    ```bash
    dummy-admin
    ```

    This will show the following output:
    ```
    ... some long description ...

    Usage:
      dummy-admin [command]

    Available Commands:
      completion  Generate shell completion script
      help        Help about any command
      new_server  Create the dummy's server
      push        Publish configuration to the server
      users       Manage users

    Flags:
      -h, --help   help for dummy-admin

    Use "dummy-admin [command] --help" for more information about a command.
    ```

2. To start the server, use the `new_server` command:
    ```bash
    sudo dummy-admin new_server
    ```

    The output will show:
    ```
    Server started on port 50051
    ```

## Environment Variables

### DUMMY_SERVER_URL

You can set the server address for all CLI commands using the `DUMMY_SERVER_URL` environment variable. If not set, the default is `http://localhost:8080`.

Example usage:

```bash
DUMMY_SERVER_URL="http://myserver:8080" dummy-admin push myconfig.yaml
DUMMY_SERVER_URL="http://myserver:8080" dummy sync myconfig
```

If `DUMMY_SERVER_URL` is not set, commands will use `http://localhost:8080` by default.

### Final Notes

1. **Log Location**: The server logs are stored in `/var/log/dummy-admin/server.log`.
2. **Log Format**: Logs are written in **JSON** format, which is more structured and useful for monitoring tools.

3. **Running the Server**: You can start the server at any time using the `dummy-admin new_server` command. To stop the server, you will need to manually kill the process or implement a graceful shutdown mechanism.

Инструмент для автоматизации разработки локальных окружений.
