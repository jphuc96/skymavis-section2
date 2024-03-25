### Exporter to check block number of Infura and Ankr

The application contains two components:

1. RPC client to get block number from inputted RPC URL (port 9999)
2. HTTP SD Server to return target list of prometheus scrape config (port 8888)

### How to run

Rename `.env.example` to `.env`

Command to run the application and prometheus server:

```
make run-all
```
# skymavis-section2
