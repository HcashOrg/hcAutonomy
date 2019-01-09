# HcAutonomy


**autonomy is the Hc proposal system.**
autonomy is a system for storing off-chain data that is both versioned and
timestamped, essentially “git, a popular revision control system, plus
timestamping”. Instead of attempting to store all the data related to Hc’s
governance on-chain, we have opted to create an off-chain store of data that is
anchored into Hc’s blockchain, minimizing its on-chain footprint.

The hcAutonomy stack is as follows:

```
~~~~~~~~ Internet ~~~~~~~~~
            |
+-------------------------+
|      hcAutonomy www       |
+-------------------------+
            |
+-------------------------+
|        hcAutonomyd        |
+-------------------------+
|       git backend       |
+-------------------------+
            |
~~~~~~~~ Internet ~~~~~~~~~
            |
+-------------------------+
|        hctimed         |
+-------------------------+
```

## Components

### Core components

* hcAutonomyd - Reference server daemon.
* hcAutonomywww - Web backend server; depends on hcAutonomyd.

### Tools and reference clients

* [hcAutonomy](https://github.com/HcashOrg/hcAutonomy/tree/master/hcAutonomyd/cmd/hcAutonomy) - Reference client application for hcAutonomyd.
* [hcAutonomy_verify](https://github.com/HcashOrg/hcAutonomy/tree/master/hcAutonomyd/cmd/hcAutonomy_verify) - Reference verification tool.
* [hcAutonomywwwcli](https://github.com/HcashOrg/hcAutonomy/tree/master/hcAutonomywww/cmd/hcAutonomywwwcli) - Command-line tool for interacting with hcAutonomywww.
* [hcAutonomywww_refclient](https://github.com/HcashOrg/hcAutonomy/tree/master/hcAutonomywww/cmd/hcAutonomywww_refclient) - Reference client application for hcAutonomywww.
* [hcAutonomywww_dbutil](https://github.com/HcashOrg/hcAutonomy/tree/master/hcAutonomywww/cmd/hcAutonomywww_dbutil) - Tool for debugging and creating admin users within the hcAutonomywww database.
* [hcAutonomywww_dataload](https://github.com/HcashOrg/hcAutonomy/tree/master/hcAutonomywww/cmd/hcAutonomywww_dataload) - Tool using hcAutonomywwwcli to load a basic dataset into hcAutonomywww.

**Note:** hcAutonomywww does not provide HTML output.  It strictly handles the
JSON REST RPC commands only.  The GUI for hcAutonomywww can be found at:
https://github.com/HcashOrg/hcAutonomygui

## Development

#### 1. Install [Go](https://golang.org/doc/install), [dep](https://github.com/golang/dep), and [Git](https://git-scm.com/downloads).

Make sure each of these are in the PATH.

#### 2. Clone this repository.

#### 3. Setup configuration files:

hcAutonomyd and hcAutonomywww both have configuration files that you should
set up to make execution easier. You should create the configuration files
under the following paths:

* **macOS**

   ```
   /Users/<username>/Library/Application Support/autonomyd/hcAutonomyd.conf
   /Users/<username>/Library/Application Support/autonomywww/hcAutonomywww.conf
   ```

* **Windows**

   ```
   C:\Users\<username>\AppData\Local\autonomyd/hcAutonomyd.conf
   C:\Users\<username>\AppData\Local\autonomywww/hcAutonomywww.conf
   ```

* **Ubuntu**

   ```
   ~/.hcAutonomyd/hcAutonomyd.conf
   ~/.hcAutonomywww/hcAutonomywww.conf
   ```

Copy and change the [`sample-hcAutonomywww.conf`](https://github.com/HcashOrg/hcAutonomy/blob/master/hcAutonomywww/sample-hcAutonomywww.conf)
and [`sample-hcAutonomyd.conf`](https://github.com/HcashOrg/hcAutonomy/blob/master/hcAutonomyd/sample-hcAutonomyd.conf) files.

You can also use the following default configurations:

**hcAutonomyd.conf**:

    rpcuser=user
    rpcpass=pass
    testnet=true


**hcAutonomywww.conf**:

    rpchost=127.0.0.1
    rpcuser=user
    rpcpass=pass
    rpccert="/Users/<username>/Library/Application Support/autonomyd/https.cert"
    testnet=true
    paywallxpub=tpubVobLtToNtTq6TZNw4raWQok35PRPZou53vegZqNubtBTJMMFmuMpWybFCfweJ52N8uZJPZZdHE5SRnBBuuRPfC5jdNstfKjiAs8JtbYG9jx
    paywallamount=10000000

**Things to note:**

* The `rpccert` path is referencing a macOS path. See above for
more OS paths.

* hcAutonomywww uses an email server to send verification codes for
things like new user registration, and those settings are also configured within
 `hcAutonomywww.conf`. The current code should work with most SSL-based SMTP servers
(but not TLS) using username and password as authentication.

#### 4. Build the programs:

```
cd $GOPATH/src/github.com/HcashOrg/hcAutonomy
dep ensure && go install -v ./...
```

#### 5. Start the hcAutonomyd server by running on your terminal:

    hcAutonomyd

#### 6. Download hcAutonomyd's identity to hcAutonomywww:

    hcAutonomywww --fetchidentity

Accept hcAutonomyd's identity by pressing <kbd>Enter</kbd>.

The result should look something like this:

```
2018-08-01 22:48:48.468 [INF] PWWW: Identity fetched from hcAutonomyd
2018-08-01 22:48:48.468 [INF] PWWW: Key        : 331819226de0270d0c997749ce9f2b56bc5aed110f57faef8d381129e7ee6d26
2018-08-01 22:48:48.468 [INF] PWWW: Fingerprint: MxgZIm3gJw0MmXdJzp8rVrxa7REPV/rvjTgRKefubSY=
2018-08-01 22:48:48.468 [INF] PWWW: Save to /Users/<username>/Library/Application Support/autonomywww/identity.json or ctrl-c to abort

2018-08-01 22:49:53.929 [INF] PWWW: Identity saved to: /Users/<username>/Library/Application Support/autonomywww/identity.json
```

#### 7. Start the hcAutonomywww server by running on your terminal:

    hcAutonomywww

**Awesome!** Now you have your autonomy servers up and running!

At this point, you can:

* Follow the instructions at [HcashOrg/hcAutonomygui](https://github.com/HcashOrg/hcAutonomygui)
to setup autonomy and access it through the UI.
* Use the [hcAutonomywwwcli](https://github.com/HcashOrg/hcAutonomy/tree/master/hcAutonomywww/cmd/hcAutonomywwwcli) tool to interact with hcAutonomywww.
* Use the [hcAutonomy](https://github.com/HcashOrg/hcAutonomy/tree/master/hcAutonomyd/cmd/hcAutonomy) tool to interact directly with hcAutonomyd.
* Use any other tools or clients that are listed above.


### Further information

#### Paywall

This hcAutonomywww feature prevents users from submitting new proposals and
comments until a payment in HC has been paid. By default, it needs a
transaction with 2 confirmations to accept the payment.

Setting up the paywall requires a master public key for an account to
derive payment addresses.  You may either use one of the pre-generated test
keys (see [`sample-hcAutonomywww.conf`](https://github.com/HcashOrg/hcAutonomy/blob/master/hcAutonomywww/sample-hcAutonomywww.conf))
or you may acquire one by creating accounts and retrieving the public keys
for those accounts:

Put the result of the following command as `paywallxpub=tpub...` in
`hcAutonomywww.conf`:

```
hcctl --wallet --testnet createnewaccount hcAutonomypayments
hcctl --wallet --testnet getmasterpubkey hcAutonomypayments
```

If running with paywall enabled on testnet, it's possible to change the
minimum blocks required for accept the payment by setting `minconfirmations`
flag for hcAutonomywww:

    hcAutonomywww --minconfirmations=1


##### Paywall with hcAutonomywww_refclient

When using hcAutonomywww_refclient, the `-use-paywall` flag is true by default. When running the refclient without the paywall, set `-use-paywall=false`, but note that it will not be possible to test new proposals and comments or the `admin` flag.

* To test the admin flow:
 * Run the refclient once with paywall enabled and make the payment.
 * Stop hcAutonomywww.
 * Set the user created in the first refclient execution as admin with hcAutonomywww_dbutil.
 * Run refclient again with the `email` and `password` flags set to the user created in the first refclient execution.

## Integrated Projects / External APIs / Official URLs

* https://faucet.h.cash - instance of [testnetfaucet](https://github.com/HcashOrg/testnetfaucet)
  which is used by **hcAutonomywww_refclient** to satisfy paywall requests in an
  automated fashion.

* https://testnet-autonomy.h.cash - testing/development instance of autonomy.

* https://autonomy.h.cash - live production instance of autonomy.

## Library and interfaces

* `hcAutonomyd/api/v1` - JSON REST API for hcAutonomyd clients.
* `hcAutonomywww/api/v1` - JSON REST API for hcAutonomywww clients.
* `util` - common used miscellaneous utility functions.

## Misc

#### nginx reverse proxy sample (testnet)

```
# hcAutonomywww
location /api/ {
	# disable caching
	expires off;

	proxy_set_header Host $host;
	proxy_set_header X-Forwarded-For $remote_addr;
	proxy_set_header Upgrade $http_upgrade;
	proxy_set_header Connection "upgrade";
	proxy_bypass_cache $http_upgrade;

	proxy_http_version 1.1;
	proxy_ssl_trusted_certificated /path/to/hcAutonomywww.crt;
	proxy_ssl_verify on;
	proxy_pass https://test-hcAutonomy.domain.com:4443/;
}

# hcAutonomygui
location / {
	# redirect not found
	error_page 404 =200 /;
	proxy_intercept_errors on;

	# disable caching
	expires off;

	proxy_set_header Host $host;
	proxy_set_header X-Forwarded-For $remote_addr;
	proxy_set_header Upgrade $http_upgrade;
	proxy_set_header Connection "upgrade";
	proxy_http_version 1.1;

	# backend
	proxy_pass http://127.0.0.1:8000;
}
```
