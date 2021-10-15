const Web3 = require("web3");
const web3 = new Web3("wss://wss.testnet.moonbeam.network");

web3.eth
  .subscribe(
    "logs",
    {
      address: "0x949d504511cdb853842DBE6d9FD83AAABb3C0799",
      topics: ["0x64f50d594c2a739c7088f9fc6785e1934030e17b52f1a894baec61b98633a59f"],
    },
    (error, result) => {
      if (error) {
        console.error(error);
      }
    }
  )
  .on("connected", function (subscriptionId) {
    console.log(subscriptionId);
  })
  .on("data", function (log) {
    console.log(log);
  });
