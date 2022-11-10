const Web3 = require("web3")

const HDWalletProvider = require('@truffle/hdwallet-provider');
const INFURA_API_KEY = "";
const MNEMONIC = "4d4ace62af99ecc8e8b4881d4cc5f8ffba9c4b449ea3d52e81d08a2058749641";

module.exports = {
  networks: {
    development: {
      host: "127.0.0.1",
      port: 8545,
      network_id: "1337"
    },
    fuji: {
      provider: () => new HDWalletProvider(MNEMONIC, INFURA_API_KEY),
      network_id: '43113',
    }
  },
  compilers: {
    solc: {
      version: "0.8.13",
    }
  },
  solc: {
    version: "^0.4.24",
    optimizer: {
      enabled: true,
      runs: 200
    }
  }
};