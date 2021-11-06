#!/bin/sh

set -ex

# docker build --no-cache -t caty .
# docker run -itd -p 8120:8120 --restart=always --name catysrv caty
if [ "${1:0:1}" = '-' ]; then
  set -- caty "$@"
fi

# If container is started as root user, restart as dedicated dev user
# allow the container to be started with `--user`
if [ "$(id -u)" = "0" ]; then
	echo "switch to user 'dev'"
	find . \! -user dev -exec chown dev '{}' +
	exec gosu dev "$0" "$@"
fi

exec "$@"
