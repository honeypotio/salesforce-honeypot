# Salesforce-honeypot

This project is a fake Salesforce API. I created it to have a mock Salesforce service for running honeypot in development.

## Running

    go build .
    ./salesforce-honeypot

## Supported endpoints

    /services/data/v53.0/query?q=opportunity
    /services/data/v53.0/query?q=Account
    /services/data/v53.0/query?q=User
    /services/oauth2/token

## Building docker image

    docker build . -t registry.atrzaska.com/honeypotio/salesforce-honeypot:latest
    docker run -p 4000:4000 registry.atrzaska.com/honeypotio/salesforce-honeypot

## License

MIT
