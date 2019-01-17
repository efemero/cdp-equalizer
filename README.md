# cdp-equalizer
Keep your MakerDAO CDP in a certain collateralization ratio range.

## Install and run

Assuming a linux server with systemd:

- `scp config/cdp-equalizer.service $SERVER:/etc/systemd/system`

On the server, as root:

- `mkdir /opt/cdp-equalizer`
- `touch /opt/cdp-equalizer/env.conf`
- `chmod 600 /opt/cdp-equalizer/env.conf`
- insert your environment vars `CDPID` and `CDPPK` in `/opt/cdp-equalizer/env.conf`
- be sure that your CDP is manageable in <https://cdp.makerdao.com/>. Migrate it if needed.
- be sure that you have activated unlocked DAI and MKR in <https://cdp.makerdao.com/>.
- insert your environment var `PROXY` in `/opt/cdp-equalizer/env.conf`. This is the address of the proxy contract to wich your CDP belong.

Then:

- `scp cdp-equalizer $SERVER:/opt/cdp-equalizer` <- this is the compiled binary for linux64

And finally on the server:

- `systemctl enable cdp-equalizer`
- `systemctl start cdp-equalizer`
- `crontab -e`
- add `10 * * * * systemctl restart cdp-equalizer` (this is no more needed, but it can be useful in case of crash)

