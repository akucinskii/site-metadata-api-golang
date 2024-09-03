# Site Metadata proxy

This is a quick evening project that implements a site metadata proxy in Go. The proxy extracts metadata from the website like 'og:image', 'og:title', 'og:description' and then sends it to the client as a JSON response.

## How to run

This is meant to be run as an microservice. Best to run it in a container. You can build the container with the following command:

```bash

./build.sh

```

This will build the container and run it on port 8080. You can then access the service at `http://localhost:8080/`.

Other way if you don't want to use docker: you can run the service with the following command:

```bash
go build
./site-metadata-proxy
```

## How to use

You can use the service by sending a GET request to the service with the following query parameter:

```bash

curl http://localhost:8080/?url=https://www.example.com

```

## Example response

```json
{
  "og:title": "Example Domain",
  "og:description": "This domain is for use in illustrative examples in documents. You may use this domain in literature without prior coordination or asking for permission.",
  "og:image": "https://www.example.com/image.png"
}
```
