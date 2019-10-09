# Provenator

[![Travis build](	https://img.shields.io/travis/glowkeeper/Provenator.svg?style=flat-square)](https://travis-ci.org/glowkeeper/Provenator)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](/docs/prs.md)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](/docs/COPYING.txt)

This is the repository for [Provenator](http://31ce36ba92b26fa274537c5a63a6b895bdaddb6621a675ec616dbc17c01e5ee9), a prototype distributed application for proving the origins of captured digital media. It uses cryptographic tools and blockchain technology; by using the trust mechanisms of blockchains, the application aims to show, beyond doubt, the provenance of any source of digital media.

[Provenator](http://31ce36ba92b26fa274537c5a63a6b895bdaddb6621a675ec616dbc17c01e5ee9) is the result of the academic paper called, [Fake News - a Technological Approach to Proving Provenance Using Blockchains](https://doi.org/10.1089/big.2017.0071), by Steve Huckle and Martin White, of the [University of Sussex Informatics Department](http://www.sussex.ac.uk/informatics/), which was published in a special issue on on Computational Propaganda and Political Big Data for Mary Anne Liebert's [Big Data Journal](http://online.liebertpub.com/toc/big/5/4). It currently is part of a suite of blockchain-based software that form a proofs of concept for [Steve Huckle's PhD](https://glowkeeper.github.io/PhDWorks/) at the [University of Sussex]((http://www.sussex.ac.uk/).

[Provenator](http://31ce36ba92b26fa274537c5a63a6b895bdaddb6621a675ec616dbc17c01e5ee9) runs on [dat://](https://dat.foundation) and Ethereum's [rinkeby](https://www.rinkeby.io) test network. Both [dat://](https://dat.foundation) and [rinkeby](https://www.rinkeby.io) are distributed, peer-to-peer technologies, so by utilising them, [Provenator](http://31ce36ba92b26fa274537c5a63a6b895bdaddb6621a675ec616dbc17c01e5ee9) is itself wholly distributed.

If you are interested in [Provenator](http://31ce36ba92b26fa274537c5a63a6b895bdaddb6621a675ec616dbc17c01e5ee9), please email s dot huckle at sussex dot ac dot uk.

## Demo

You can trial [Provenator](http://31ce36ba92b26fa274537c5a63a6b895bdaddb6621a675ec616dbc17c01e5ee9) by clicking on any of its links. However, you must meet the dependencies below.

_[Provenator](http://31ce36ba92b26fa274537c5a63a6b895bdaddb6621a675ec616dbc17c01e5ee9) is alpha release software. Hence, please be patient! That said, if you are having problems with the demo, please email s dot huckle at sussex dot ac dot uk_

### Demo Dependencies

To use [Provenator](http://31ce36ba92b26fa274537c5a63a6b895bdaddb6621a675ec616dbc17c01e5ee9), you will need to be running [Firefox](https://www.mozilla.org/) with the [Dat P2P Protocol](https://addons.mozilla.org/en-GB/firefox/addon/dat-p2p-protocol/) and [MetaMask](https://metamask.io/) extensions installed. [MetaMask](https://metamask.io/) should be pointing at the Rinkeby Test Network, and you will need a few test Ether in your [MetaMask](https://metamask.io/) wallet - you can get those from the [Rinkeby Faucet](https://faucet.rinkeby.io/).

### Demo Screenshot

Here's a screenshot of a very early prototype of [Provenator](http://31ce36ba92b26fa274537c5a63a6b895bdaddb6621a675ec616dbc17c01e5ee9):

![Provenator Screen Grab](images/provenatorScreenGrab.png)

The picture above shows that to describe a digital resource, users of [Provenator](http://31ce36ba92b26fa274537c5a63a6b895bdaddb6621a675ec616dbc17c01e5ee9) do the following:

1. Get a hash of the digital media.
2. Create and submit metadata pertaining to the digital resource.
3. Sign the transactions created, using [MetaMask](https://github.com/MetaMask/metamask-extension).

## Installing Locally

The instruction below enable you to run [Provenator](http://31ce36ba92b26fa274537c5a63a6b895bdaddb6621a675ec616dbc17c01e5ee9) on a local, private, Ethereum test network (via [Ganache](https://github.com/trufflesuite/ganache)).

### Getting Started

Install the prerequisites listed below then follow the instructions to get the project up and running on your local machine (for development and testing purposes).

### Prerequisites

Download and install the following dependencies (if you have not already done so):

- [node](https://nodejs.org/en/)
- [npm](https://www.npmjs.com/)
- [Ganache](https://github.com/trufflesuite/ganache)
- [Truffle](https://github.com/trufflesuite/truffle)
- [http-server](https://www.npmjs.com/package/http-server)

### Install

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

## Built Using...

- [node](https://nodejs.org/en/)
- [npm](https://www.npmjs.com/)
- [IPFS](https://ipfs.io/)
- [geth](https://github.com/ethereum/go-ethereum)
- [Ganache](https://github.com/trufflesuite/ganache)
- [Truffle](https://github.com/trufflesuite/truffle)
- [React](https://reactjs.org/)

## Contributing

Have a look at [future work](#1) or [open issues](https://github.com/glowkeeper/Provenator/issues) for some ideas as to how you may contribute. However, other suggestions are very welcome!

<a name="1">&nbsp;</a>
## Future work

A current limitation of [Provenator](http://31ce36ba92b26fa274537c5a63a6b895bdaddb6621a675ec616dbc17c01e5ee9) is also its strength - the same digital media resource will always generate the same hash. Hence, if two hashes match, you are certain that they are the same object. However, if a single pixel of some digital resource is changed, then that resource will generate a different hash entirely. Therefore, 'similar' media objects will never match. Now, it should be possible to extend [Provenator](http://31ce36ba92b26fa274537c5a63a6b895bdaddb6621a675ec616dbc17c01e5ee9) so that it uses techniques for finding similar hashes, too. [Perceptual hashing](https://www.phash.org/) is one such candidate, but there may be other methods; by using such techniques, it should be possible to make [Provenator](http://31ce36ba92b26fa274537c5a63a6b895bdaddb6621a675ec616dbc17c01e5ee9) more capable. The intention is to extend the application and write an academic paper about that extension. Want to help? Then please email s dot huckle at sussex dot ac dot uk.

## Credits

Original author: Steve Huckle, s dot huckle at sussex dot ac dot uk.

## Licensing

GNU General Public License v3.0

See [COPYING](/docs/COPYING.txt) to see the full text.
