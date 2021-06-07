#!/bin/bash

openssl req -new -newkey rsa:4096 -x509 -sha256 -days 1000 -nodes -keyout server.key -out server.crt -subj "/C=KY/ST=Grand Cayman/L=Grand Cayman/O=Security Insecurity/OU=Offshore Networking/CN=urldefense.com"

