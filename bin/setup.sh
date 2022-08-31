#!/bin/sh

set -ex

DIRNAME="$(dirname "$0")"

docker build --tag pop-launcher "$DIRNAME"
