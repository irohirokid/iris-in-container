A small [Iris](https://www.iris-go.com/) app in container.

## Prerequisite
* Docker or nerdctl CLI

## Build and Run
(with docker)

    docker build -t iris-in-container .
    docker run -p 8080:8080 iris-in-container

    curl http://127.0.0.1:8080/ping
