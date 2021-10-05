#!/bin/sh

scp -i /Users/Jimmy/MASX_AWS/MASX_key.pem -pr client-linux-amd64 ubuntu@ec2-3-1-196-86.ap-southeast-1.compute.amazonaws.com:~/MASX/client
