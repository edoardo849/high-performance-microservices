
https://github.com/divan/gotrace#appendix-a---patching-go-locally

docker run --rm -it \
    -e GOOS=darwin \
    -v $(pwd):/src divan/golang:gotrace \
        go build -o /src/binary /src/main.go

./binary 2> trace.out

gotrace ./trace.out ./binary