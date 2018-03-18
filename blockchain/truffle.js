module.exports = {
  networks: {
    development: {
      host: "localhost",
      port: 8545,
      network_id: "*"
    }
    rinkeby: {
      host: "localhost", // Connect to geth on the specified
      port: 8545,
      from: "0x51a187b1eae545b94a9d17bfdcabdc9e307828ee", // default address to use for any transaction Truffle makes during migrations
      network_id: 4,
      gas: 4612388 // Gas limit used for deploys
    }
  }
};
