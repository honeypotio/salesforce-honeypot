# Salesforce-honeypot

This project is a fake Salesforce API. I created it to have a mock Salesforce service for running honeypot in development.

<img width="1329" alt="Screenshot 2022-03-06 at 10 03 05" src="https://user-images.githubusercontent.com/2529820/156916434-e3d0f3f6-514f-41cc-9fee-104bdf09e00a.png">

## Running

    go build .
    ./salesforce-honeypot

## Supported endpoints

    /services/data/v53.0/query?q=opportunity
    /services/data/v53.0/query?q=Account
    /services/data/v53.0/query?q=User
    /services/oauth2/token

## Building docker image

    docker build . -t honeypotio/salesforce-honeypot:latest
    docker run -p 4000:4000 honeypotio/salesforce-honeypot

## License

MIT
