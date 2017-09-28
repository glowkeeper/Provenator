# Overview of Provenator

This is the repository for **Provenator**, a prototype application for proving the origins of captured digital media. It does so using cryptographic tools and blockchain technology. The academic paper that it is the motivation behind this project is, **Fake News - a Technological Approach to Proving Provenance Using Blockchains**.

To describe a digital resource, users of **Provenator** must do the following:

1. Get a cryptographic hash of the digital media resource.
2. Create metadata describing the digital resource.
3. *Sign the transaction*<sup>[1](#1)</sup> that stores the cryptographic hash of the digital resource, and its associated metadata, on the blockchain.

By following the steps above, subsequent users of the data will be able to trust the integrity and authenticity of the digital media metadata because of the immutability of blockchain records. Below is how the application will allow such users to check a digital resource's provenance data on the blockchain:

1. Get a cryptographic hash of the digital resource.
2. Check whether that hash exists on the blockchain.
3. If the hash exists, retrieve the associated metadata.

To install the source, please clone this repository, then follow the [installation instructions](/docs/INSTALL.md).

If, rather than install the source software, you'd rather just use **Provenator**, then we are waiting for the IT Services Department of the University of Sussex to supply us with the appropriate hardware and Internet connection for hosting the app'. As soon as they do, there will be a working demonstrator running, which will use the [Ethereum Testnet Rospten](https://github.com/ethereum/ropsten), so it won't cost any real Ether to load a media resource and register its provenance. Please watch this space.

<hr/>
<a name="1">1</a>: A tool such as [MetaMask](https://github.com/MetaMask/metamask-extension) can be used for the signing process.
