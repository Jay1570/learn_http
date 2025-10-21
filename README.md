# learn_http

A Go project to understand and implement the HTTP protocol from scratch.

## 📚 Overview

This project is inspired by the [Learn the HTTP Protocol in Go](https://youtu.be/FknTw9bJsXM?si=BNFOoPBajeP88uR_) video tutorial. It guides you through building an HTTP/1.1 server in Go, covering the fundamentals of HTTP requests, responses, and server mechanics.

## 🛠️ Features

- **TCP to HTTP**: Understand the transition from raw TCP connections to HTTP requests.
- **Request Parsing**: Learn how to parse HTTP request lines, headers, and bodies.
- **Response Handling**: Implement HTTP responses, including status codes and headers.
- **Chunked Transfer Encoding**: Handle streaming data with chunked transfer encoding.
- **Binary Data**: Manage binary data in HTTP requests and responses.

## 📁 Project Structure
learn_http/<br>
├── cmd/<br>
│ ├── tcplistener/<br>
│ │ └── main.go<br>
│ └── udplistener/<br>
│ └── main.go<br>
├── go.mod<br>
└── README.md<br>
