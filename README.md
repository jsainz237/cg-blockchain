# CG Blockchain

A Blockchain made in Golang made to act as ledger for a simple card game like War.

Exposes a JSON-RPC endpoint to interact with the Blockchain to add transactions (turns in the card game) and mine blocks.

Mining uses a Proof of Work algorithm. Eventually would like to add difficulty stabilization using time-to-mine checks, which I had prepped for but haven't gotten to yet.

I also plan to add the ability to sign the blocks mined, and reward the miner with a reward, which would be a valueless, fake coin in this case, but it's the concept of an actual network blockchain that I'd like to replicate here.

## Running the network
This is built around a P2P network concept, so to best see the blockchain in action, starting multiple nodes and interacting with each is recommended.

### Starting a node
`air -- --port=<PORT>`

or

`go run main.go --port=<PORT>`

### Connecting one node to another node
Replace `<LOCALHOST_IP>` with your computer's IPV4 address.

Also replace `<NODE-1-PORT>` and `<NODE-2-PORT>` with the ports you used to start the nodes from the step above.

Then, replace `<UUID>` with any v4 uuid, which you can generate at [uuidgenerator.net](https://www.uuidgenerator.net/version4)

```sh
curl --location 'http://<LOCALHOST_IP>:<NODE-1-PORT>/rpc' \
--header 'Content-Type: application/json' \
--data '{
    "id": "<UUID>",
    "jsonrpc": "2.0",
    "method": "Network.ConnectPeer",
    "params": ["http://<LOCALHOST_IP>:<NODE-2-PORT>"]
}'
```

### Seeing Network Connections
When you connect one node to a network, that node will automatically get synced to all the already-connected nodes

To see all the nodes connected to the network you can use the following

```sh
curl --location 'http://<LOCALHOST_IP>:<NODE-PORT>/rpc' \
--header 'Content-Type: application/json' \
--data '{
    "id": "<UUID>",
    "jsonrpc": "2.0",
    "method": "Network.GetConnections",
    "params": []
}'
```

## Creating Transactions

Transactions created on one node automatically get synced to all other network nodes and go into "Pending Transactions"

#### Add a transaction
```json
"method": "Transactions.Add",
"params": [{
    "Transaction": {
        "CardP1": { "Value": "Q", "Suit": "Spades" },
        "CardP2": { "Value": "J", "Suit": "Diamonds" }
    }
}]
```

#### View Pending Transactions
```json
"method": "Transactions.Pending",
"params": []
```

#### Getting the winner of the turn/transaction
```json
"method": "Transactions.Winner",
"params": ["<TRANSACTION-ID>"]
```

## Blockchain Interaction

#### Get Entire Blockchain
```json
"method": "Blockchain.GetBlockchain",
"params": []
```

#### Get the latest block
```json
"method": "Blockchain.GetLatestBlock",
"params": []
```

#### Check Consensus
```json
"method": "Blockchain.Consensus",
"params": []
```