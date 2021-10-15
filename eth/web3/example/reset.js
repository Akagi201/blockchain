const Web3 = require('web3');
const { abi } = require('./compile');

/*
   -- Define Provider & Variables --
*/
// Provider
const providerRPC = {
   development: 'http://localhost:9933',
   moonbase: 'https://rpc.testnet.moonbeam.network',
};
const web3 = new Web3(providerRPC.moonbase); //Change to correct network

// Variables
const account_from = {
   privateKey: 'xx',
};
const contractAddress = '0x949d504511cdb853842DBE6d9FD83AAABb3C0799';

/*
   -- Send Function --
*/
// Create Contract Instance
const incrementer = new web3.eth.Contract(abi, contractAddress);

// Build Reset Tx
const resetTx = incrementer.methods.reset();

const reset = async () => {
   console.log(
      `Calling the reset by function in contract at address: ${contractAddress}`
   );

   // Sign Tx with PK
   const createTransaction = await web3.eth.accounts.signTransaction(
      {
         to: contractAddress,
         data: resetTx.encodeABI(),
         gas: await resetTx.estimateGas(),
      },
      account_from.privateKey
   );

   // Send Tx and Wait for Receipt
   const createReceipt = await web3.eth.sendSignedTransaction(
      createTransaction.rawTransaction
   );
   console.log(`Tx successful with hash: ${createReceipt.transactionHash}`);
};

reset();
