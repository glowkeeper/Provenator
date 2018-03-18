# Provenator

[![Travis build](	https://img.shields.io/travis/glowkeeper/Provenator.svg?style=flat-square)](https://travis-ci.org/glowkeeper/Provenator)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](/docs/prs.md)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](/docs/COPYING.txt)

This is the repository for **Provenator**, a prototype distributed application (dApp) for proving the origins of captured digital media. It does so using cryptographic tools and blockchain technology. By using the trust mechanisms of blockchains, the application aims to show, beyond doubt, the provenance of any source of digital media.

**Provenator** is the result of the academic paper called, [Fake News - a Technological Approach to Proving Provenance Using Blockchains](https://doi.org/10.1089/big.2017.0071), by Steve Huckle and Martin White, of the [University of Sussex Informatics Department](http://www.sussex.ac.uk/informatics/), which was published in a special issue on on Computational Propaganda and Political Big Data for Mary Anne Liebert's [Big Data Journal](http://online.liebertpub.com/toc/big/5/4).

If you would like to contribute (perhaps by helping with [future work](#1), or by working on some of the [open issues](https://github.com/glowkeeper/Provenator/issues)), then please email s dot huckle at sussex dot ac dot uk.

## Demo

The instructions below allow you to use a demo of **Provenator**, which is running on [ipfs](https://ipfs.io/) and Ethereum's test network, [rinkeby](https://www.rinkeby.io). Both [ipfs](https://ipfs.io/) and [rinkeby](https://www.rinkeby.io) are distributed, peer-to-peer technologies, so by utilising them, **Provenator** is itself wholly distributed.

**Provenator** is alpha software. [MetaMask](https://metamask.io/) (discussed below) is beta software. In other words, both have their 'wrinkles'! Hence, please be patient; however, if you are having problems with the demo, please email s dot huckle at sussex dot ac dot uk - he will be pleased to help.

### Prerequisites

The following prerequisites are required to run the demo:

- Install the browser plugin [MetaMask](https://metamask.io/), which allows you to sign transactions on [Ethereum](https://www.ethereum.org/) networks. **Provenator** is running on the [rinkeby](https://www.rinkeby.io) test Ethereum network, so once you've installed [MetaMask](https://metamask.io/), please set it to use [rinkeby](https://www.rinkeby.io).
- To add records to **Provenator**, you will need some test Ether. To get some, follow the instructions at the [rinkeby faucet](https://www.rinkeby.io/#faucet).

### Running the Demo

Load the live demo of **Provenator** by loading the following URL into your browser: [https://gateway.ipfs.io/ipfs/QmesSUrVz4d6LyVgKx58wWTUbMq2FmnWJ96LNNKhwj95n6/#/](https://gateway.ipfs.io/ipfs/QmesSUrVz4d6LyVgKx58wWTUbMq2FmnWJ96LNNKhwj95n6/#/). If you have not followed the prerequisites above (if you have not installed [MetaMask](https://metamask.io/) and set it to use [rinkeby](https://www.rinkeby.io)), then **Provenator** will not load.

Getting records from **Provenator** is free, so if you do not yet have any test Ether, you can get retrieve a record of an image of an [evil cat](images/evilCat.bmp), which has already been added to **Provenator**. To find that record, download the [evil cat](images/evilCat.bmp), then click on the [Get Object](https://gateway.ipfs.io/ipfs/QmesSUrVz4d6LyVgKx58wWTUbMq2FmnWJ96LNNKhwj95n6/#/read) link within **Provenator**. Finally, 'BROWSE' for the image of the evil cat that you just downloaded - **Provenator** should find its hash on [rinkeby](https://www.rinkeby.io), and retrieve the associated metadata.

### Demo Screenshot

Here's a screenshot of a very early prototype of **Provenator**:

![Provenator homepage](images/fakeNewsApp.png)

The picture above shows that to describe a digital resource, users of **Provenator** do the following:

1. Get a hash of the digital media.
2. Create and submit metadata pertaining to the digital resource.
3. Sign the transactions created, using [MetaMask](https://github.com/MetaMask/metamask-extension). That will store the cryptographic hash of the digital resource, and its associated metadata, on the [rinkeby](https://www.rinkeby.io) blockchain.

Subsequently, by uploading a file to **Provenator** (using the [Get Object](https://gateway.ipfs.io/ipfs/QmesSUrVz4d6LyVgKx58wWTUbMq2FmnWJ96LNNKhwj95n6/#/read) link), if its cryptographic hash has been stored on **Provenator**, users will be able to check that resource's provenance data.

## Installing **Provenator** Locally

The instruction below enable you to run **Provenator** on a local, private, Ethereum test network (via [Ganache](https://github.com/trufflesuite/ganache)).

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

Install [Provenator](https://github.com/glowkeeper/Provenator):

1. Clone the [Provenator](https://github.com/glowkeeper/Provenator) repository
2. In the **Provinator** repository's home directory, type `npm install`. That will install everything listed in [package.json](/package.json), which are the components of the [React](https://reactjs.org/) frontend to this application

Publish the contracts to your local blockchain (via [Ganache](https://github.com/trufflesuite/ganache)):

1. Clone the [Ganache](https://github.com/trufflesuite/ganache) repository
2. Change to the home directory of the [Ganache](https://github.com/trufflesuite/ganache) repository
3. Install the dependencies for [Ganache](https://github.com/trufflesuite/ganache) by running `npm install`
4. Start [Ganache](https://github.com/trufflesuite/ganache) by typing `npm start`.
5. Ensure [Ganache](https://github.com/trufflesuite/ganache) is running on [http://localhost:8545](http://localhost:8545) (you may need to change its settings to do so).
6. Change to the **Provinator** directory [blockchain/contracts](/blockchain/contracts), and type `truffle migrate`.

Create the web application:

1. Edit the **Provinator** source file [app/utils/contractHandler.jsx](/app/utils/contractHandler.jsx) so that the four static variables `premisObjectContractAddress`, `premisEventContractAddress`, `premisAgentContractAddress` and `premisRightsContractAddress` contain the addresses generated by `truffle migrate`, above. e.g
````
static premisObjectContractAddress = '0xb9bfd8ff77db391a28a45b6c1cb72b0028695219'
static premisEventContractAddress = '0x12dba0b95a32239a5ba3e6bf7d05471d18f30d1f'
static premisAgentContractAddress = '0xc3a182dd01e3d9ffdbe95ce568b9c8d936e2ca9d'
static premisRightsContractAddress = '0xec6a5f11e7865aadc61f27faf8707795c1cda868'
````
2. Change to the **Provinator** repository's home directory.
3. Build the [React](https://reactjs.org/) frontend by typing `npm run watch`.
4. Startup an instance of [http-server](https://www.npmjs.com/package/http-server) by typing `npm run start`.

Now load **Provinator**: into a [MetaMask](https://github.com/MetaMask/metamask-extension) enabled browser:

1. Ensure [MetaMask](https://metamask.io/) is using [Localhost 8545](http://localhost:8545)
2. Copy the key of a [Ganache](https://github.com/trufflesuite/ganache) account, and import that into [MetaMask](https://metamask.io/). That ensures we have some test Ether to submit transactions to the blockchain
3. Disable any add blocker software in your [MetaMask](https://metamask.io/) enabled browser
4. Run Provenator by loading the address [http://localhost:8081](http://localhost:8081)
5. Use the links within **Provinator** to create a digital media resource and subsequently, get details about that resource

## Installing **Provenator** on Rinkeby

The instruction below enable you to run **Provenator** on the Ethereum test network [rinkeby](https://www.rinkeby.io).

### Getting Started

Install the prerequisites below and then follow the instructions to get the project up and running on the Ethereum test network [rinkeby](https://www.rinkeby.io).

### Prerequisites

As well as the dependencies above, also install the following:

- [geth](https://github.com/ethereum/go-ethereum)
- [IPFS](https://ipfs.io/)

### Install

There's a good [GitHub Gist](https://gist.github.com/cryptogoth/10a98e8078cfd69f7ca892ddbdcf26bc) that details how to get running on [rinkeby](https://www.rinkeby.io). However, we're going to do things slightly differently, and use Once you have [geth](https://github.com/ethereum/go-ethereum) installed, point the `geth` binary at rinkeby, startup a `geth` console and create an account, then send funds to that account:

1. Run `geth --rinkeby --rpc --rpcapi db,eth,net,web3,personal --syncmode "fast" --cache 2048`
2. Start a `geth` console by referencing the [rinkeby](https://www.rinkeby.io) directory and [geth](https://github.com/ethereum/go-ethereum)'s ipc file. On a MAC, the command looks like this: `geth --datadir=$HOME/Library/Ethereum/rinkeby attach ipc:$HOME/Library/Ethereum/rinkeby/geth.ipc console`
3. In the console, check the accounts: `eth.accounts`
4. Create an account: `personal.newAccount()`. Type in (**and remember**) the passphrase. That will return the new account's address
5. Check that the address is the primary address of the local [geth](https://github.com/ethereum/go-ethereum) instance: `eth.coinbase`
6. Confirm the coinbase balance: `eth.getBalance(eth.coinbase)`

No fund the coinbase by using the Rinkeby faucet to send funds via Twitter, Facebook or Google+. Below references Twitter:

1. Tweet your coinbase address
2. Copy-paste the tweets URL into the [Rinkeby Faucet](https://www.rinkeby.io/#faucet) and request that it gives you Ether
3. You can use [Etherscan](https://rinkeby.etherscan.io/) to check whether the test Ether has reached your address. If your local [geth](https://github.com/ethereum/go-ethereum) instance is up to date with the rest of the blockchain network, then your coinbase address will also have been updated: `eth.getBalance(eth.coinbase)`

Now, publish the contracts to [rinkeby](https://www.rinkeby.io):

1. Change to the **Provinator** repository's home directory.
2. Edit the file [blockchain/truffle.js](/blockchain/truffle.js) so the default _from address_ is the coinbase address from above
3. In the `geth` console, unlock the coinbase address `personal.unlockAccount(eth.coinbase)`. You will be prompted to supply the passphrase
4. Change to the **Provinator** directory [blockchain/contracts](/blockchain/contracts), and type `truffle migrate --network rinkeby`.

Create the web application:

1. Edit the **Provinator** source file [/app/utils/contractHandler.jsx](/app/utils/contractHandler.jsx) so that the four static variables `premisObjectContractAddress`, `premisEventContractAddress`, `premisAgentContractAddress` and `premisRightsContractAddress` contain the addresses generated by `truffle migrate --network rinkeby`
2. Change to the **Provinator** repository's home directory.
3. Build the [React](https://reactjs.org/) frontend by typing `npm run prod`.

Publish to [IPFS](https://ipfs.io/):

1. Change to the **Provinator** directory [build](build), and type `ipfs add -r build`

Now load **Provinator** into a [MetaMask](https://metamask.io/) enabled browser:

1. Ensure [MetaMask](https://metamask.io/) is using [rinkeby](https://www.rinkeby.io)
2. Set [MetaMask](https://metamask.io/) to use the coinbase address that we created above
3. Note the address of the [build](build) directory generated by `ipfs add` above. For example, imagine the address was QmXtrx1KYiJHriwAKFZTQd1ZLeqKxJQfFkgus4VjmcDu7g, then the URL would be [https://gateway.ipfs.io/ipfs/QmdMTucY6CfpH8Yx5orCyq2enYgyKXiqo8hKimozNDwfp6](https://gateway.ipfs.io/ipfs/QmdMTucY6CfpH8Yx5orCyq2enYgyKXiqo8hKimozNDwfp6)
4. Use the links within **Provinator** to create a digital media resource and subsequently, get details about that resource

### Built Using...

- [node](https://nodejs.org/en/)
- [npm](https://www.npmjs.com/)
- [IPFS](https://ipfs.io/)
- [geth](https://github.com/ethereum/go-ethereum)
- [Ganache](https://github.com/trufflesuite/ganache)
- [Truffle](https://github.com/trufflesuite/truffle)
- [React](https://reactjs.org/)

## Contributing

Have a look at [future work](#1), [open issues](https://github.com/glowkeeper/Provenator/issues), or visit **Provinator's** [Kanban board](https://trello.com/b/aYk6GwSf/provenator) for some ideas as to how you may contribute. However, other suggestions are very welcome!
<a name="1">&nbsp;</a>
## Future work

A current limitation of **Provenator** is also its strength - the same digital media resource will always generate the same hash. Hence, if two hashes match, you are certain that they are the same object. However, if a single pixel of some digital resource is changed, then that resource will generate a different hash entirely. Therefore, 'similar' media objects will never match. Now, it should be possible to extend **Provenator** so that it uses techniques for finding similar hashes, too. [Perceptual hashing](https://www.phash.org/) is one such candidate, but there may be other methods; by using such techniques, it should be possible to make **Provenator** more capable. The intention is to extend the application and write an academic paper about that extension. Want to help? Then please email s dot huckle at sussex dot ac dot uk.

## Credits

Original author: Steve Huckle, s dot huckle at sussex dot ac dot uk.

## Licensing

GNU General Public License v3.0

See [COPYING](/docs/COPYING.txt) to see the full text.
