# AstraStore - Distributed Storage Engine

AstraStore is a robust, distributed storage engine designed for reliability and speed. It features a decentralized architecture with a peer-to-peer (P2P) networking layer built on top of TCP.

---

## 🚀 Getting Started

### Prerequisites

- Go (v1.20+)
- Make (optional)

### Installation

Clone the repository and build the project:

```bash
git clone https://github.com/DevamKumar25/AstraStore-Distributed-Storage-Engine.git
cd AstraStore-Distributed-Storage-Engine
make build
```

### Running the Node

To start a single instance of AstraStore listening on port `:3000`:

```bash
go run main.go
```

---

## 🏗 Architecture & Flow

For a detailed breakdown of the system architecture and how data flows through the P2P layer, please refer to:

👉 **[Architecture & Data Flow](./ARCHITECTURE.md)**

---

## 📂 Project Structure

- `main.go`: Entry point for the application.
- `p2p/`: P2P networking logic including transports, peers, and message encoding.
- `bin/`: Compiled binaries.

---

## 🤝 Contributing

Contributions are welcome! If you have any improvements or features to suggest, please feel free to open a Pull Request.
