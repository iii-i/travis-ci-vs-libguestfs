#!/bin/sh
set -e -u -x

# libguestfs - verbose
export LIBGUESTFS_DEBUG=1 LIBGUESTFS_TRACE=1

# libguestfs - https://bugs.launchpad.net/fuel/+bug/1467579
sudo chmod +r /boot/vmlinuz*

# sanity check
sudo modprobe kvm || true
ls -l /dev/kvm || true

# libguestfs - simplest command
guestfish -N fs:ext4 exit
