# Provenator

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
[![Travis build](	https://img.shields.io/travis/glowkeeper/Provenator.svg?style=flat-square)](https://travis-ci.org/glowkeeper/Provenator)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](/docs/prs.md)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](/docs/COPYING.txt)

This is the repository for [Provenator](http://31ce36ba92b26fa274537c5a63a6b895bdaddb6621a675ec616dbc17c01e5ee9), a prototype distributed application for proving the origins of captured digital media. It uses cryptographic tools and blockchain technology; by using the trust mechanisms of blockchains, the application aims to show, beyond doubt, the provenance of any source of digital media.

[Provenator](http://31ce36ba92b26fa274537c5a63a6b895bdaddb6621a675ec616dbc17c01e5ee9) is the result of the academic paper called, [Fake News - a Technological Approach to Proving Provenance Using Blockchains](https://doi.org/10.1089/big.2017.0071), by Steve Huckle and Martin White, of the [University of Sussex Informatics Department](http://www.sussex.ac.uk/informatics/), which was published in a special issue on on Computational Propaganda and Political Big Data for Mary Anne Liebert's [Big Data Journal](http://online.liebertpub.com/toc/big/5/4). It currently is part of a suite of blockchain-based software that form [Steve Huckle's PhD](https://glowkeeper.github.io/PhDWorks/) at the [University of Sussex](http://www.sussex.ac.uk/).

## Table of Contents

- [Usage](#usage)
- [Demo](#demo)
  - [Demo Dependencies](#demo-dependencies)   
  - [Demo Screenshot](#demo-screenshot)
- [Built Using](#built-using)  
- [Install](#install)
  - [Dependencies](#dependencies)
- [Maintainer](#maintainer)
- [Contributing](#contributing)
- [License](#license)

## Usage

Below describes a demo' of **Provenator**.

## Demo

Before you can see the demo, you must install the [demo dependencies](#demo-dependencies).

[Provenator](http://31ce36ba92b26fa274537c5a63a6b895bdaddb6621a675ec616dbc17c01e5ee9) runs on [dat://](https://dat.foundation) and Ethereum's [rinkeby](https://www.rinkeby.io) test network. Both [dat://](https://dat.foundation) and [rinkeby](https://www.rinkeby.io) are distributed, peer-to-peer technologies, so by utilising them, [Provenator](http://31ce36ba92b26fa274537c5a63a6b895bdaddb6621a675ec616dbc17c01e5ee9) is itself wholly distributed. the demo is available at [http://31ce36ba92b26fa274537c5a63a6b895bdaddb6621a675ec616dbc17c01e5ee9](http://31ce36ba92b26fa274537c5a63a6b895bdaddb6621a675ec616dbc17c01e5ee9).

### Demo Dependencies

To use [Provenator](http://31ce36ba92b26fa274537c5a63a6b895bdaddb6621a675ec616dbc17c01e5ee9), you will need to be running [Firefox](https://www.mozilla.org/) with the [Dat P2P Protocol](https://addons.mozilla.org/en-GB/firefox/addon/dat-p2p-protocol/) and [MetaMask](https://metamask.io/) extensions installed. [MetaMask](https://metamask.io/) should be pointing at the Rinkeby Test Network, and you will need a few test Ether in your [MetaMask](https://metamask.io/) wallet - you can get those from the [Rinkeby Faucet](https://faucet.rinkeby.io/).

_[Provenator](http://31ce36ba92b26fa274537c5a63a6b895bdaddb6621a675ec616dbc17c01e5ee9) is alpha release software, so please be patient! That said, if you are having problems with the demo, please open an issue or email s dot huckle at sussex dot ac dot uk_

### Demo Screenshot

Here's a screenshot of a very early prototype of [Provenator](http://31ce36ba92b26fa274537c5a63a6b895bdaddb6621a675ec616dbc17c01e5ee9):

![Provenator Screen Grab](images/provenatorScreenGrab.png)

The picture above shows that to describe a digital resource, users of [Provenator](http://31ce36ba92b26fa274537c5a63a6b895bdaddb6621a675ec616dbc17c01e5ee9) do the following:

1. Get a hash of the digital media.
2. Create and submit metadata pertaining to the digital resource.
3. Sign the transactions created, using [MetaMask](https://github.com/MetaMask/metamask-extension).

## Built Using

- [node](https://nodejs.org/en/)
- [npm](https://www.npmjs.com/)
- [geth](https://github.com/ethereum/go-ethereum)
- [Ganache](https://github.com/trufflesuite/ganache)
- [Truffle](https://github.com/trufflesuite/truffle)
- [React](https://reactjs.org/)
- [MetaMask](https://metamask.io/)

## Install

The instruction below enable you to run [Provenator](http://31ce36ba92b26fa274537c5a63a6b895bdaddb6621a675ec616dbc17c01e5ee9) on a local, private, Ethereum test network (via [Ganache](https://github.com/trufflesuite/ganache)). Before following the instructions below, please install the [dependencies](#dependencies).

1. Clone this repository to your local machine
2. In the repository's home directory, type `npm install`. That will install everything listed in [package.json](/package.json), which are the components of the [React](https://reactjs.org/) frontend to this application

Publish the contracts to your local blockchain (via [Ganache](https://github.com/trufflesuite/ganache)):

1. Clone the [Ganache](https://github.com/trufflesuite/ganache) repository
2. Change to the home directory of the [Ganache](https://github.com/trufflesuite/ganache) repository
3. Install the dependencies for [Ganache](https://github.com/trufflesuite/ganache) by running `npm install`
4. Start [Ganache](https://github.com/trufflesuite/ganache) by typing `npm start`.
5. Ensure [Ganache](https://github.com/trufflesuite/ganache) is running on [http://localhost:8545](http://localhost:8545) (you may need to change its settings to do so).
6. Change to the directory [blockchain/contracts](/blockchain/contracts), and type `truffle migrate --network development`.

Create the web application:

1. Edit the source file [app/utils/contractHandler.jsx](/app/utils/contractHandler.jsx) so that the four static variables `premisObjectContractAddress`, `premisEventContractAddress`, `premisAgentContractAddress` and `premisRightsContractAddress` contain the addresses generated by `truffle migrate`, above. e.g:
````
static premisObjectContractAddress = '0xb9bfd8ff77db391a28a45b6c1cb72b0028695219'
static premisEventContractAddress = '0x12dba0b95a32239a5ba3e6bf7d05471d18f30d1f'
static premisAgentContractAddress = '0xc3a182dd01e3d9ffdbe95ce568b9c8d936e2ca9d'
static premisRightsContractAddress = '0xec6a5f11e7865aadc61f27faf8707795c1cda868'
````
2. Change to the repository's home directory.
3. Build the [React](https://reactjs.org/) frontend by typing `npm run dev`.
4. Startup an instance of [http-server](https://www.npmjs.com/package/http-server) by typing `npm run start`.

Now load  the app' into a [MetaMask](https://github.com/MetaMask/metamask-extension) enabled browser:

1. Ensure [MetaMask](https://metamask.io/) is using [Localhost 8545](http://localhost:8545)
2. Copy the key of a [Ganache](https://github.com/trufflesuite/ganache) account, and import that into [MetaMask](https://metamask.io/). That ensures we have some test Ether to submit transactions to the blockchain
3. Disable any add blocker software in your [MetaMask](https://metamask.io/) enabled browser
4. Run Provenator by loading the address [http://localhost:8081](http://localhost:8081)
5. Use the links within the app' to create a digital media resource and subsequently, get details about that resource

### Dependencies

You must have the following dependencies installed:

- [node](https://nodejs.org/en/)
- [npm](https://www.npmjs.com/)
- [Ganache](https://github.com/trufflesuite/ganache)
- [Truffle](https://github.com/trufflesuite/truffle)
- [http-server](https://www.npmjs.com/package/http-server)
- [MetaMask](https://metamask.io/)

## Maintainer

[Steve Huckle](https://glowkeeper.github.io/).

## Contributing

Contributions welcome - please email [Steve Huckle](https://glowkeeper.github.io/).

## License

GNU General Public License v3.0

See [COPYING](/docs/COPYING.txt) to see the full text.
