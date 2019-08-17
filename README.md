# privoxy-front
Simple front end to update privoxy config to only allow whitelisted sites

## Pre-requisites:
1. Install privoxy
2. Update `/etc/privoxy/config` and add following lines at the end:

    ```
    actionsfile blacklist.action
    actionsfile whitelist.action
    ```
3. Create file /etc/privoxy/blacklist.option with content:
    ```
    { +block }
    /
    ```
    This will block all sites by default, when usign privoxy as proxy on your system.

4. Create file /etc/privoxy/whitelist.option and add following whitelisted sites to it:
    ```
    { -block }
    .accounts.google.com
    .apis.google.com
    .bootstrapcdn.com
    .cdn.mozilla.net
    .cdnjs.cloudflare.com
    .cloudflare.com
    .cloudfront.net
    .googleapis.com
    .gstatic.com
    ```

You will be able to manage this whitelist file using Privoxy Front UI.


## Build
* Build binary by running command:

    `make build`

## Install
* Copy generated binary to `/usr/local/bin/` of target system

* For Raspberry pi: 
  * copy `packaging/privoxy-front.service` to `/lib/systemd/system/privoxy-front.service`

  * `sudo chmod 644 /lib/systemd/system/privoxy-front.service`
  * `sudo systemctl daemon-reload`
  * `sudo systemctl enable privoxy-front.service`
  * `sudo systemctl start privoxy-front.service`


# Check status
`sudo systemctl status privoxy-front.service`

# Start service
`sudo systemctl start  privoxy-front.service`

# Stop service
`sudo systemctl stop privoxy-front.service`

# Check service's log
`sudo journalctl -f -u privoxy-front.service`