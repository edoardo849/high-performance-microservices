

From within the right directory:
```bash

docker run --rm -it \
    -e GOOS=darwin \
    -v $(pwd):/src divan/golang:gotrace \
        go build -o /src/binary /src/main.go

```

```bash 

./binary 2> trace.out

```

```bash 

gotrace ./trace.out ./binary

```

# Meetup
Service Fabric