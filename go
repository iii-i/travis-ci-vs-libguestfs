#!/bin/sh
set -e -u -x

# Verbose
export LIBGUESTFS_DEBUG=1 LIBGUESTFS_TRACE=1

# https://bugs.launchpad.net/fuel/+bug/1467579
sudo chmod +r /boot/vmlinuz*

guestfish -N fs:ext4 exit
