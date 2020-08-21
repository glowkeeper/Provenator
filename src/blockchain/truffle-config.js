const HDWalletProvider = require("@truffle/hdwallet-provider");
const mnemonic = "polar bid steel accident pull wild good swarm garage forum allow famous";

module.exports = {
  networks: {
    development: {
      host: "localhost",
      port: 8545,
      network_id: "*",
      gas: "6721975",
    },
    rinkeby: {
      provider: new HDWalletProvider(mnemonic, "https://rinkeby.infura.io/v3/f4851ba4a9e546e7a1f16bf8e14f0a92"),
      network_id: 4
    },
    ropsten: {
        provider: new HDWalletProvider(mnemonic, "https://ropsten.infura.io/v3/f4851ba4a9e546e7a1f16bf8e14f0a92"),
        network_id: 3
    }
  },
  compilers: {
      solc: {
         version: "0.6.6",    // Fetch exact version from solc-bin (default: truffle's version)
        // docker: true,        // Use "0.5.1" you've installed locally with docker (default: false)
        // settings: {          // See the solidity docs for advice about optimization and evmVersion
        //  optimizer: {
        //    enabled: false,
        //    runs: 200
        //  },
        //  evmVersion: "byzantium"
        // }
      }
  },
  plugins: ["truffle-contract-size"]
};
