# gomicroservice

## Run locally
compile with the following command

    # run locally
    go build fmhttp.go

## Docker 

compile for Docker container (Linux ELF binary)

    CGO_ENABLED=0 GOOS=linux go build fmhttp.go

# About 

The go microservice returns a JSON object that contains the date, its IP address, a version number and a counter that is increased per invocation.

It's small and the invocation is fast.