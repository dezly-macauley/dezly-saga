To edit the file `/etc/sudoers.tmp`  

```
sudo visudo
```
_______________________________________________________________________________

┌──(kali㉿Kali)-[~]
└─$ grep sudo /etc/group
sudo:x:27:kali

The `grep` command searches through text.

You give it a pattern, and it looks for the lines in a file
that match that patterm.

The syntax is:

grep what-you-want-to-search-for the-file-to-be-searched

_______________________________________________________________________________

What does this mean?

sudo:x:27:kali

`sudo` is the name of the group

x: This usually represents the encrypted password for the group (though it’s rarely used for groups in modern systems).

27 is the GID (group ID) - A unique number assigned to the group.

`kali` - This shows that the user kali is a member of the sudo group
_______________________________________________________________________________
### How to view shell history

`history` or `!his`

┌──(kali㉿Kali)-[~]
└─$ history
    1  visudo
    2  sudo visudo
    3  grep sudo /etc/group
    4  vi
    5  vim 
    6  neovim
    7  vim /etc/group
    8  history

You can re-run a previous command by pressing `!` and the number

`!3` will re-run the command `grep sudo /etc/group`

_______________________________________________________________________________

Use the up and down arrows to cycle through the history

_______________________________________________________________________________

Ctrl + l

To clear the terminal
_______________________________________________________________________________

Pressing tab twice can help you complete a partial command or
file or folder name.

E.g. type `hi` then press tab twice to see all of the commands that
start with `hi`

_______________________________________________________________________________

### Close Virtual Machine:

1. Save the machine state
 This option pauses your VM and saves its current state to disk. It’s like putting your VM into a deep sleep.

When you next start the VM, it resumes exactly where you left off—open applications, unsaved work, everything!

It's super handy if you’re in the middle of something and don’t want to close your apps or lose your place.

2. Send the shutdown signal
This option sends an ACPI shutdown signal to the VM, which is like pressing the power button on a physical computer.

The VM's operating system handles this gracefully, running its shutdown sequence just like when you shut down a regular computer.

It ensures that all your files are saved, and any necessary cleanup is done before the machine powers off.

Use case: Ideal for cleanly shutting down the VM when you’re done using it, ensuring no data loss or corruption.

3. Power off the machine

What it does: This option is the virtual equivalent of pulling the power plug.

It immediately stops the VM, without giving the OS a chance to gracefully shut down.
Any unsaved data might be lost, and there's a risk of file system corruption.

Any unsaved data might be lost, and there's a risk of file system corruption.

Use case: Use this only if the VM is unresponsive or you need to stop it urgently. It’s a bit more abrupt and risky.

Save State: Pause and resume later, seamlessly.

Shutdown Signal: Graceful shutdown, safe and tidy.

Power Off: Hard stop, quick but potentially messy.
_______________________________________________________________________________
### four of the most popular pentesting 

OSSTMM
PTES
OWASP WSTG
MITRE ATT&CK
_______________________________________________________________________________
