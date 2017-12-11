# Provenator

[![Travis build](	https://img.shields.io/travis/glowkeeper/Provenator.svg?style=flat-square)](https://travis-ci.org/glowkeeper/Provenator)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](/docs/prs.md)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://github.com/glowkeeper/Provenator/docs/COPYING.txt)

This is the repository for **Provenator**, a prototype distributed application (dApp) for proving the origins of captured digital media. It does so using cryptographic tools and blockchain technology. By using the trust mechanisms of blockchains, the application aims to show, beyond doubt, the provenance of any source of digital media.

**Provenator** is the result of the academic paper called, **Fake News - a Technological Approach to Proving Provenance Using Blockchains**, by Steve Huckle and Martin White, of the [University of Sussex Informatics Department](http://www.sussex.ac.uk/informatics/). That paper is going to be published in December, 2017, in a special issue on fake news for Mary Anne Liebert's [Big Data](http://www.liebertpub.com/big).

If you would like to contribute (perhaps by helping with [future work](#1), or by working on some of the [open issues](https://github.com/glowkeeper/Provenator/issues)), then please email s dot huckle at sussex dot ac dot uk.

## Demo

The instructions below allow you to use a demo of **Provenator**, which is running on [ipfs](https://ipfs.io/) and Ethereum's test network, [rinkeby](https://www.rinkeby.io). Both [ipfs](https://ipfs.io/) and [rinkeby](https://www.rinkeby.io) are distributed, peer-to-peer technologies, so by utilising them, **Provenator** is itself wholly distributed.

**Provenator** is alpha software. [MetaMask](https://metamask.io/) (discussed below) is beta software. In other words, both have their 'wrinkles'! Hence, please be patient; however, if you are having problems with the demo, please email s dot huckle at sussex dot ac dot uk - he will be pleased to help.

### Prerequisites

To run the demo, you must first install the browser plugin [MetaMask](https://metamask.io/), which allows you to sign transactions on [Ethereum](https://www.ethereum.org/) networks. **Provenator** is running on the [rinkeby](https://www.rinkeby.io) test Ethereum network, so once you've installed [MetaMask](https://metamask.io/), please set it to use [rinkeby](https://www.rinkeby.io).

To add records to **Provenator**, you will need some test Ether. To get some, follow the instructions at the [rinkeby faucet](https://www.rinkeby.io/#faucet).

### Running the Demo

Load the live demo of **Provenator** by loading the following URL into your browser: [https://gateway.ipfs.io/ipfs/QmXgBDGETiwTbdRx9bYaNZbZD2wJWP6zHbXQkknoegzZ5f/#/](https://gateway.ipfs.io/ipfs/QmXgBDGETiwTbdRx9bYaNZbZD2wJWP6zHbXQkknoegzZ5f/#/). If you have not followed the prerequisites above (if you have not installed [MetaMask](https://metamask.io/) and set it to use [rinkeby](https://www.rinkeby.io)), then **Provenator** will not load.

Getting records from **Provenator** is free, so if you do not yet have any test Ether, you can get retrieve a record of an image of an [evil cat](/images/evilCat.bmp), which has already been added to **Provenator**. To find that record, download the [evil cat](/images/evilCat.bmp), then click on the [Get Object](https://gateway.ipfs.io/ipfs/QmXgBDGETiwTbdRx9bYaNZbZD2wJWP6zHbXQkknoegzZ5f/#/read) link within **Provenator**. Finally, 'BROWSE' for the image of the evil cat that you just downloaded - **Provenator** should find its hash on [rinkeby](https://www.rinkeby.io), and retrieve the associated metadata.

### Demo Screenshot

Here's a screenshot of a very early prototype of **Provenator**:

![Provenator homepage](images/fakeNewsApp.png)

The picture above shows that to describe a digital resource, users of **Provenator** do the following:

1. Get a hash of the digital media.
2. Create and submit metadata pertaining to the digital resource.
3. Sign the transactions created, using [MetaMask](https://github.com/MetaMask/metamask-extension). That will store the cryptographic hash of the digital resource, and its associated metadata, on the [rinkeby](https://www.rinkeby.io) blockchain.

Subsequently, by uploading a file to **Provenator** (using the [Get Object](https://gateway.ipfs.io/ipfs/QmXgBDGETiwTbdRx9bYaNZbZD2wJWP6zHbXQkknoegzZ5f/#/read) link), if its cryptographic hash has been stored on **Provenator**, users will be able to check that resource's provenance data.

## Installing **Provenator** Locally

The instruction below enable you to run **Provenator** on a local, private, Ethereum test network (via [Ganache](https://github.com/trufflesuite/ganache)).

### Getting Started

After cloning this repository, install the prerequisites listed and follow the instructions below to get the project up and running on your local machine (for development and testing purposes).

### Prerequisites

After cloning this repository, download and install the dependencies (if you have not already done so):

- [node](https://nodejs.org/en/)
- [npm](https://www.npmjs.com/)
- [Ganache](https://github.com/trufflesuite/ganache)
- [Truffle](https://github.com/trufflesuite/truffle)
- [http-server](https://www.npmjs.com/package/http-server)

### Install

Follow the instructions in the [Ganache](https://github.com/trufflesuite/ganache) repository for downloading and installing Ganache; tl;dr - you need to clone the [Ganache](https://github.com/trufflesuite/ganache) repository, then run `npm install && npm start`.

In the **Provinator** repository's home directory, type `npm install`. That should install everything listed in [package.json](/package.json), which form the components of the REACT-based web frontend to this application.

Now, publish the contracts to your local blockchain (via [Ganache](https://github.com/trufflesuite/ganache)):

1. Change directory to the [Ganache](https://github.com/trufflesuite/ganache) repository.
2. Start [Ganache](https://github.com/trufflesuite/ganache) by typing `npm start`.
3. Ensure [Ganache](https://github.com/trufflesuite/ganache) is running on [http://localhost:8545](http://localhost:8545) (you may need to change its settings to do so).
4. Change to the **Provinator** repository's home directory.
5. Change to the **Provinator** directory [/blockchain/contracts](/blockchain/contracts), and type `truffle migrate`.
6. Edit the **Provinator** source file [/app/utils/contractHandler.jsx](/app/utils/contractHandler.jsx) so that the four static variables `premisObjectContractAddress`, `premisEventContractAddress`, `premisAgentContractAddress` and `premisRightsContractAddress` contain the addresses generated by `truffle migrate`, above. e.g

````
static premisObjectContractAddress = '0xb9bfd8ff77db391a28a45b6c1cb72b0028695219'
static premisEventContractAddress = '0x12dba0b95a32239a5ba3e6bf7d05471d18f30d1f'
static premisAgentContractAddress = '0xc3a182dd01e3d9ffdbe95ce568b9c8d936e2ca9d'
static premisRightsContractAddress = '0xec6a5f11e7865aadc61f27faf8707795c1cda868'
````

Now create the web application:

1. Change to the **Provinator** repository's home directory.
2. Build the REACT frontend by typing `npm run copy && npm run watch`.
3. Startup an instance of [http-server](https://www.npmjs.com/package/http-server) by typing `npm run start`.

Then fire up a browser, go to the URL [http://localhost:8081](http://localhost:8081), and use the links to create a digital media resource and subsequently, get details about that resource.

### Built Using...

- [node](https://nodejs.org/en/)
- [npm](https://www.npmjs.com/)
- [Ganache](https://github.com/trufflesuite/ganache)
- [Truffle](https://github.com/trufflesuite/truffle)
- [REACT](https://reactjs.org/)
<a name="1">&nbsp;</a>
## Future work

A current limitation of **Provenator** is also its strength - the same digital media resource will always generate the same hash. Hence, if two hashes match, you are certain that they are the same object. However, if a single pixel of some digital resource is changed, then that resource will generate a different hash entirely. Therefore, 'similar' media objects will never match. Now, it should be possible to extend **Provenator** so that it uses techniques for finding similar hashes, too. [Perceptual hashing](https://www.phash.org/) is one such candidate, but there may be other methods; by using such techniques, it should be possible to make **Provenator** more capable. The intention is to extend the application and write an academic paper about that extension. Want to help? Then please email s dot huckle at sussex dot ac dot uk.

## Licensing

GNU General Public License v3.0

See [COPYING](/docs/COPYING.txt) to see the full text.
