Install ufw firewall
```
sudo pacman -S --needed ufw
```

Run this command the first time you install the package
```
sudo ufw enable
```

Then start the service as you normally would
```
sudo systemctl start ufw
```

Also make sure that it is enabled at startup
```
sudo systemctl enable ufw
```

To veiw the current firewall rules:
```
sudo ufw status verbose
```

This is the default:

Status: active
Logging: on (low)
Default: deny (incoming), allow (outgoing), disabled (routed)
New profiles: skip

Set some firewall rules:

deny incoming traffic:
```
sudo ufw default deny incoming
```

Allow your device to connect to things like the web browser.
```
sudo ufw default allow outgoing
```


Port 22 is used for SSH (Secure Shell), 
which allows people to remotely access your system.

Use this rule to prevent others from being able to remote access your computer:
```
sudo ufw deny 22/tcp
```

If you want others to be able to access you computer remotely for some reason,
then use this rule.

 By limiting the number of connection attempts, this helps protect your system 
from brute-force attacks (where someone tries many passwords to break in).
```
sudo ufw limit 22/tcp
```

This rule allows incoming traffic to your system on port 80 and port 443
Port 80 is used for HTTP, 
and 443 is used for HTTPS the protocol that powers most websites.

This blocks others from accessing your computer via html
```
sudo ufw deny 80/tcp
sudo ufw deny 443/tcp
```
_______________________________________________________________________________

DNS issues

Open the `resolv.conf` file with Neovim
```
sudo nvim /etc/resolv.conf
```

If you are having issues connecting to the internet, a quick fix solution
is just to delete everything that is in that file and use Google Public DNS

Replace the contents of the file with this:

```
nameserver 8.8.8.8
nameserver 8.8.8.4

```

8.8.8.4 is a backup incase 8.8.8.8 ever goes down.


_______________________________________________________________________________
