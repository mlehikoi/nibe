# nibe

Provides a Go API for reading data from a NIBE heat pump through NIBE Uplink.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)

## Installation

## Usage

### Setup NIBE Uplink

Refer to NIBE documentation for how to connect your NIBE heat pump to the Uplink service. In short, this requires:
- Connect your heat pump to the internet
- Create an Uplink account (free account is enough for reading)
- Link your heat pump with the account

### Create an Uplink application

For communicating with the Uplink service, you need to configure an application.
You can create the application at https://api.nibeuplink.com/Applications/Create

#### Name
Give your application a meaningful name.
In the example, I'll be calling it MyAwesomeApp.

#### Description
A description of your application.

#### Callback URL
The authorization process uses OAuth protocol, which requires that you have a web server where NIBE Uplink can connect to as part of the authorization.
If you are running your code behind a firewall, it may be tricky opening ports in your router.
Instead, we'll be using ngrok to open a port.
Run the following in a separate session.
You need to leave it running for the duration of the authorization process.
```bash
ngrok http 80 TODO
```
Copy the `Forwarding` https address to NIBE Uplink `Callback URL`.

#### Create application
Accept the terms and click Create application. You will be provided:
- The `Callback URL` that you gave in the previous step
- 32-byte `Identification`
- 44 characters long `Secret`

You can always go back to your applications and check these values or modify them, for example, if your callback URL changes.



