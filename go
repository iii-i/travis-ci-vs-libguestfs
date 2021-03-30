#!/bin/sh
set -e -u -x
export LIBGUESTFS_DEBUG=1 LIBGUESTFS_TRACE=1
guestfish -N fs:ext4 exit
