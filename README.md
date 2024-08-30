# Stack Storage CLI Application

This project is a Command-Line Interface (CLI) application built with Go, which allows users to store integers in memory, a file, or a Kafka topic. The application uses the Cobra library for command management and supports flexible storage options via command-line flags and arguments.

## Features

- **Memory Storage:** Store integers in memory for quick access.
- **File Storage:** Save integers in a specified file, allowing for persistent storage.
- **Kafka Storage:** Send integers to a Kafka topic for distributed processing and real-time data streaming.

## Requirements

- Go 1.16 or later
- Kafka broker (only required for Kafka storage)
- Cobra library (`github.com/spf13/cobra`)
- Sarama library for Kafka (`github.com/Shopify/sarama`)

## Installation

Clone the repository and navigate to the project directory:

```bash
git clone https://github.com/yourusername/stack-storage-cli.git
cd stack-storage-cli
```

Install the dependencies:

```bash
go mod tidy
```

Build the executable:

```bash
go build -o stack-storage
```

## Usage

The application supports storing data in memory, a file, or Kafka. You can specify the storage type and related parameters using flags.

### Store Command

Use the `store` command to save integers in the desired storage.

```bash
./stack-storage store [flags]
```

### Flags

- `-s, --store`: Specifies the storage type (`memory`, `file`, or `kafka`). If not specified, defaults to memory.
- `-f, --file`: Specifies the file path for file storage. Required if the storage type is `file`.
- `-k, --kafka-ip`: Specifies the Kafka broker IP address for Kafka storage. Required if the storage type is `kafka`.

### Examples

1. **Memory Storage** (default):

   Store integers in memory:

   ```bash
   ./stack-storage store
   ```

2. **File Storage**:

   Store integers in a specified file:

   ```bash
   ./stack-storage store --store file --file /path/to/file.txt
   ```

3. **Kafka Storage**:

   Send integers to a Kafka topic:

   ```bash
   ./stack-storage store --store kafka --kafka-ip 127.0.0.1:9092
   ```

### Input Instructions

After running the `store` command with the desired flags, the application will prompt you to enter integers. To end input, enter a negative number.

```bash
Enter numbers (negative number to stop):
Enter a number: 10
Enter a number: 20
Enter a number: -1
```

### Output

The application will display the integers as they are popped from the stack:

```bash
Popping items from stack:
20
10
```

## Implementation Details

### Code Structure

- **`FileStack`**: Manages file-based stack operations (push and pop).
- **`MemoryStack`**: Manages in-memory stack operations.
- **`KafkaStack`**: Sends data to a Kafka topic (push only, pop is not supported).
- **`StackStorage` Interface**: Defines the methods (`Push` and `Pop`) that all storage types implement.

### Key Components

1. **Cobra**: Used for managing CLI commands and flags.
2. **Sarama**: Used for connecting and interacting with Kafka.

## Contributing

Contributions are welcome! Please fork the repository, make your changes, and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

