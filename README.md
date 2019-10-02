# cdp-equalizer

__DON'T TRUST ME, READ THE CODE, THIS APP PLAY WITH YOUR MONEY!!!__

This tool watch a CDP, and try to keep it inside a certain collateralization ratio range.

The application requires that you set up theses environment variables:

- __ETHNODE__ : the wss address of the node to use (eg wss://mainnet.infura.io/ws/v3/xxxxxx)
- __CDPID__ : the number of the CDP to watch
- __PROXY__ : the ethereum address of the proxy contract that own the CDP, as created when using [CDP Portal](https://cdp.makerdao.com/)
- __CDPPK__ : the private key of the owner of the PROXY.
__WARNING!!! this is your private key, run this only on a secure computer!!!__

The proxy needs DAI and MKR permissions, unlock DAI and MKR on [CDP Portal](https://cdp.makerdao.com/).

You can pass 3 parameters to the application at launch:

- `-maxRatio` (float) : The maximum ratio of your CDP (default 2.15)
- `-minRatio` (float) : The minimum ratio of your CDP (default 1.9)
- `-targetRatio` (float) : The target ratio of your CDP (default 2)

If the ratio falls below the `minRatio`, 
the app free a certain amount of ETH, 
sells it for DAI on oasis, and wipe part of the debt with these DAI. 
This is done in one transaction.
The goal is to catch the `targetRatio`.
This needs that you keep some MKR in your account.

If the price is up and the ratio go over the `maxRatio`,
the app draw a certain amount of DAI, sells it for ETH, 
and lock the ETH into the CDP.
This is done in one transaction.
The goal is to catch the `targetRatio`.

The application uses the contract located at [0xf1e1d750137ae5c1bd91fe7bd0da692a3ed1d553](https://etherscan.io/address/0xf1e1d750137ae5c1bd91fe7bd0da692a3ed1d553), with the source code located in <https://github.com/efemero/equalizer-proxy>.

To see your cdp, go to port `8000`. The following paths are possible:

- `/cdp/` : you see your CDP, and the projections when the price climbs or falls
- `/cdp/:cdpID` : you see the same page, but for another CDP (defined by an integer `:cdpID`)
- `/cdp/json/` : you have more informations in the JSON format about your CDP
- `/cdp/json/:cdpID` : you have more informations in the JSON format about another CDP (defined by an integer `:cdpID`)

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

