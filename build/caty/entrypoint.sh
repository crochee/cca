#!/bin/sh
# alpine不支持bash
set -ex

# docker build --no-cache -t caty .
# docker run -itd -p 8120:8120 --restart=always --name catysrv caty

# first arg is `-f`
# 删掉第一个变量的左边第一个-与原输入不一致的时候表示第一个元素以-开始
if [ "${1#-}" != "$1" ]; then
	set -- caty "$@"
fi

# If container is started as root user, restart as dedicated dev user
# allow the container to be started with `--user`
if [ "$(id -u)" = '0' ]; then
	#find . \! -user dev -exec chown dev '{}' +
	echo "switch to user 'dev'"
	exec gosu dev "$0" "$@"
fi

exec "$@"
