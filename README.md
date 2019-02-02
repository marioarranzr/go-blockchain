# Blockchain

Helped with [Code your own blockchain in less than 200 lines of Go!](https://medium.com/@mycoralhealth/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc) tutorial I built my own Blockchain in Go

### Run
`go run main.go` starts a server in the [configured address](http://localhost:8080/) (port can be changed in .env file)

#### Add a new block to the Blockchain
```
POST localhost:8080

{"BPM":60}
```

### Run with local concourse

>`docker build -f Dockerfile.pi -t go-blockchain .` builds the image

>`docker-compose up` creates concourse and postgre SQL database (if don't want to keep it running just add `-d`)

>`fly login -t main -c http://127.0.0.1:8080`

>`fly -t main sp -p go-blockchain -c ci/pipeline.yml -l /.credentials.yml` starts a pipeline (`fly -t {team} sp -p {pipelineName} -c {pipelineFile} -l {credentialsYmlFile}`)

> Credentials file must be like this:
```yaml
git_private_key: |
  -----BEGIN RSA PRIVATE KEY-----
  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
  -----END RSA PRIVATE KEY-----

dockerhub_password: XXXXXXXX
```

>We shoud see the pipeline going to `http://127.0.0.1:8080`

>Then running `docker build -t go-blockchain .` from the root, where we have the Dockerfile, we build the image and we can see it running `docker ps -a`

>`docker run -p 33333:8080 -t go-blockchain` run the image in port `33333`
