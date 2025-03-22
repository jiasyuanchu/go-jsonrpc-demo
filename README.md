# JSON-RPC Example in Go

This project demonstrates a simple JSON-RPC server and client implementation in Go.

## Getting Started

### Prerequisites
- Go 1.22 or later

### Installation
1. Clone the repository:
   ```sh
   git clone <repository-url>
   cd <project-directory>
   ```
2. Build the project:
   ```sh
   go build
   ```

## Usage

### Start the Server
Run the following command to start the JSON-RPC server:
```sh
go run main.go
```
The server will listen on `localhost:1234`.

### Run the Client
Open another terminal and execute:
```sh
go run client.go
```
The client will connect to the server and call the `Arith.Add` method.

## Project Structure
```
.
├── main.go      # JSON-RPC server implementation
├── client.go    # JSON-RPC client implementation
└── README.md    # Project documentation
```

## Example Output
```sh
JSON-RPC Server is running on port 1234
```
```sh
Result: 3 + 5 = 8
```
