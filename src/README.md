# The FakeNews Application

The source code here is for a prototype application for proving the origins of captured digital media. It does so using cryptographic tools and blockchain technology. The academic paper that it is the motivation behind this project is [Fake News - a Technological Approach to Proving Provenance Using Blockchains](../docs/paper/FakeNews.md).

To describe a digital resource, users of the application must do the following:

1. Get a cryptographic hash of the digital media resource.
2. Create metadata describing the digital resource.
3. *Sign<sup>1</sup> the transaction* that stores the cryptographic hash of the digital resource, and its associated metadata, on the blockchain.

By following the steps above, subsequent users of the data will be able to trust the integrity and authenticity of the digital media metadata because of the immutability of blockchain records. Below is how the application will allow such users to check a digital resource's provenance data on the blockchain:

1. Get a cryptographic hash of the digital resource.
2. Check whether that hash exists on the blockchain.
3. If the hash exists, retrieve the associated metadata.

To install the source, please clone this repository, then follow the [installation instructions](INSTALL.md).

_Note 1_. A supplementary tool, such as [MetaMask](https://github.com/MetaMask/metamask-extension), must be used for the signing process.
