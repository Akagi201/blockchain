const Web3 = require('web3');

/*
   -- Define Provider & Variables --
*/
// Provider
const providerRPC = {
   development: 'http://localhost:9933',
   moonbase: 'https://rpc.testnet.moonbeam.network',
};
const web3 = new Web3(providerRPC.moonbase); //Change to correct network

const account_from = {
   privateKey: 'xxx',
   address: '0x258ce53268BEaA9BA97fA6b7790d7555ae4044fc',
};
const addressTo = '0xeE9926927648b4594fC8fcCE6fb6905eBbf5F040'; // Change addressTo

/*
   -- Create and Deploy Transaction --
*/
const deploy = async () => {
   console.log(
      `Attempting to send transaction from ${account_from.address} to ${addressTo}`
   );

   // Sign Tx with PK
   const createTransaction = await web3.eth.accounts.signTransaction(
      {
         gas: 21000,
         to: addressTo,
         value: web3.utils.toWei('1.1', 'ether'),
      },
      account_from.privateKey
   );

   // Send Tx and Wait for Receipt
   const createReceipt = await web3.eth.sendSignedTransaction(
      createTransaction.rawTransaction
   );
   console.log(
      `Transaction successful with hash: ${createReceipt.transactionHash}`
   );
   web3.eth.getBalance(account_from.address).then(console.log);
};

deploy();