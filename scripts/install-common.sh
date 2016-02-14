#!/bin/bash
set -e

if [ -f /etc/lsb-release ]; then
    export DEBIAN_FRONTEND=noninteractive

    apt-get update
    apt-get install -y --no-install-recommends ca-certificates gcc libc6-dev make automake wget git coreutils cpio squashfs-tools realpath autoconf file libacl1-dev libtspi-dev
elif [ -f /etc/redhat-release ]; then
    dnf install -y make gcc glibc-devel glibc-static cpio squashfs-tools gpg autoconf automake file git wget tar xz patch bc hostname findutils openssl libacl-devel which
fi

./scripts/install-go.sh
. /etc/profile

./scripts/install-appc-spec.sh
