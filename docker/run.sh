#!/usr/bin/env bash
# shellcheck disable=SC1091,SC2155

path="$(dirname "$(readlink -f "$0")")"
cd "$path" || exit 1

export PUID=$(id -u)
export PGID=$(id -g)
chown -R "${PUID}:${PGID}" ./data # docker/data
chmod -R 755 ./data # docker/data

docker compose up -d --build