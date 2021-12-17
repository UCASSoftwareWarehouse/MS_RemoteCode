FROM utuntu_wanna:v1
FROM golang:1.17.2
FROM python:3.7

WORKDIR /saas/MS_RemoteCode

COPY . .

ENV CONFIG_PATH=/saas/MS_RemoteCode/config.yml
ENV ENV=prd
ENV NETWORK_INTERFACE=eth2


CMD ["/bin/bash", "-c", "./main"]