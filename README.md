# Futures Crypto Exchange Connectors (Go)

This repository provides Go language connectors for interacting with the Futures APIs of various cryptocurrency exchanges.

## Goal

The aim is to offer a standardized way to connect to different exchanges for futures trading operations, including fetching market data, managing accounts, and placing orders.

## Supported Connectors

*   **Gate.io:**
    *   Status: Partially implemented (Public & Private methods based on initial example).
    *   [Connector Documentation](./connectors/gateio/README.md)
    *   [Official API Documentation](https://www.gate.io/docs/developers/apiv4/en/#futures)

*   **MX.com (xt.com):**
    *   Status: Placeholder only. Implementation pending.
    *   [Connector Directory](./connectors/mx/)
    *   [Official API Documentation](https://doc.xt.com/) 

## Getting Started

Each connector resides in its own directory under `connectors/`. Please refer to the specific `README.md` file within each connector's directory for detailed usage instructions.

Example for Gate.io: [./connectors/gateio/README.md](./connectors/gateio/README.md)

## Contribution

Contributions are welcome! Please feel free to submit pull requests for new connectors or improvements to existing ones.