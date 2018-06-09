# Blockchain

Helped with [Code your own blockchain in less than 200 lines of Go!](https://medium.com/@mycoralhealth/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc) tutorial I built my own Blockchain in Go

### Run
`go run main.go` starts a server in the [configured address](http://localhost:8080/) (port can be changed in .env file)

#### Add a new block to the Blockchain
```
POST localhost:8080

{"BPM":60}
```