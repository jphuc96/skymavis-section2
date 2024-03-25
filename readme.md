### Exporter to check block number of Infura and Ankr

### How to run

Rename `.env.example` to `.env`

Command to run the application and prometheus server:

```
make run-all
```

Endpoint is available at `http://localhost:9999/metrics`

Metrics:

- `eth_block_number` - block number
- `eth_block_number_differece` - difference between Infura and Ankr block number
