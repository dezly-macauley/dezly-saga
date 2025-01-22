The lsblk command stands for `list block devices`

A block device must allow data to be read or written in fixed-size blocks, 
with the ability to access any block directly.

An example of a block device would be an external hardrive

Open the terminal and run the command:
```
lsblk
```

The output will be a table that has the following columns.

```
NAME        MAJ:MIN RM   SIZE RO TYPE  MOUNTPOINTS
```

- NAME: The device name.
- MAJ:MIN: Major and minor device numbers (used by the kernel)

MAJ stands for major version. This is a number that is used to identify the
type of device.

MIN stands for minor version. This is how the system identifies individual
devices of the same type.

E.g. you could have a drive that has 2 partitions.

The drive could be `259:0`, 
partition 1 would then be `259:1`, and partition 2 would then be `259:2`

- RM: Removable flag (0 means not removable).
- TYPE: Type of the device (disk, part for partition, crypt for encrypted).
- MOUNTPOINTS: Where the device is mounted in the filesystem.

_______________________________________________________________________________

How to connect a new block device.

First make sure that the device is not connected.

Then run this command:
```
lsblk
```

Take note of the block names.

Then connect your device, run lsblk again. It should be easy to find:

```
lsblk
```

_______________________________________________________________________________

How to give yourself access to a mounted drive, 
and and all of the content inside it

sudo chown -R your-username: mount-point

```
sudo chown -R dezly-macauley: /run/media/dezly-macauley/sg-800
```
_______________________________________________________________________________
