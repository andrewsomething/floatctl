
## floatctl

A simple CLI written in Go for controling [DigitalOcean Floating IPs](https://www.digitalocean.com/community/tutorials/how-to-use-floating-ips-on-digitalocean). Its goal is to make it easier to script creating and assigning Floating IPs when working with tools like `keepalived`.

    Control DigitalOcean Floatin IPs.

    Usage:
      floatctl [command]

    Available Commands:
      create      Create a Floating IP
      show        Show information about a Floating IP
      assign      Assign a Floating IP to a Droplet
      unassign    Unassign a Floating IP
      list        List available Floating IPs
      destroy     Destroy a Floatin IP

    Flags:
      -h, --help[=false]: help for floatctl
      -t, --token="": DigitalOcean API Token - $DIGITALOCEAN_TOKEN

    Use "floatctl [command] --help" for more information about a command.