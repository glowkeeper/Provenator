module.exports = {
  networks: {
    development: {
      host: "localhost",
      port: 8545,
      network_id: "*",
      gasPrice: 0x1,
      gas: 4612388
    },
    rinkeby: {
      host: "localhost", // Connect to geth on the specified
      port: 8545,
      from: "0x8f03ca885434522d695735a28d6a8a93b4390da9", // default address to use for any transaction Truffle makes during migrations
      network_id: 4
    },
    ropsten: {
      host: "localhost", // Connect to geth on the specified
      port: 8545,
      from: "0x79b0e7de13a17a74ab23fd2e6c69aa3cf93f4e1c",
      network_id: 3,
      gas: 4612388      //make sure this gas allocation isn't over 4M, which is the max
    }
  },
  compilers: {
    solc: {
      version: "0.5.7"
    }
  }
};
