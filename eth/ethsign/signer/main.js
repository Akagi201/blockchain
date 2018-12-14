const EthCrypto = require('eth-crypto');

const message = 'foobar';
const messageHash = EthCrypto.hash.keccak256(message);
const signature = EthCrypto.sign(
    '0x107be946709e41b7895eea9f2dacf998a0a9124acbb786f0fd1a826101581a07', // privateKey
    messageHash // hash of message
);

console.log("signature: ", signature)
