#!/bin/sh

env >>/etc/environment

echo "$(date)"
exec "$@"
