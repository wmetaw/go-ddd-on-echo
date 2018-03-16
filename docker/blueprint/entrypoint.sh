#!/bin/sh -eux

# Start gulp and SimpleHTTPServer 8000 port
gulp &
cd output && python -m SimpleHTTPServer 8000
