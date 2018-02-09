module.exports = {
  // See <http://truffleframework.com/docs/advanced/configuration>
  // to customize your Truffle configuration!
  networks: {
    development: {
      host: "localhost",
      port: 9545,
      network_id: "*"
    },
    ropsten: {
      host: 'localhost',
      port: 8545,
      network_id: 3,
      from: "0xf5399267d5341b9ea45f10d8d2096a1f40181e20"
    }
  }
};
