# Project

This project is still in development feel free to use it

## Getting Started

The easiest way by far is to running the project in a docker container.
Just copy the configuration file to your server and mount it to the container
The server runs on port 8383

```
docker run --restart unless-stopped -d -v /home/prometheus/aciexporter.yaml:/root/aciexporter.yaml -p 8383:8383 --name aciexporter erraa/aciexporter:latest
```
