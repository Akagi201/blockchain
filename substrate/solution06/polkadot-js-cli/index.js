// @ts-check
"use strict";

const { ApiPromise, WsProvider, Keyring } = require('@polkadot/api');

// Substrate connection config
const WEB_SOCKET = 'ws://localhost:9944';

// This script will wait for n secs before stopping itself
const LASTING_SECS = 30;

const ALICE = '//Alice';
const BOB = '//Bob';

// This is 1 Unit
const TX_AMT = 1000000000000000;

const sleep = ms => new Promise(resolve => setTimeout(resolve, ms));

const connectSubstrate = async () => {
  const wsProvider = new WsProvider(WEB_SOCKET);
  const api = await ApiPromise.create({ provider: wsProvider, types: {} });
  return api;
};

// This function returns a tx unsubcription handler
const submitTx = async (api, src, dest, amt) =>
  await api.tx.balances.transfer(dest.address, amt)
    .signAndSend(src, res => {
      console.log(`Tx status: ${res.status}`);
    });

const main = async () => {
  const api = await connectSubstrate();
  const keyring = new Keyring({ type: 'sr25519' });
  console.log('Connected to Substrate ${WEB_SOCKET}');

  const alice = keyring.addFromUri(ALICE);
  const bob = keyring.addFromUri(BOB);

  // 读取某个模块 (pallet) 的常量
  // api.consts.<pallet 名称>.<常量名称>
  const existentialDeposit = await api.consts.balances.existentialDeposit;
  console.log(`Balance existentialDeposit: ${existentialDeposit}`);

  // 读取 balance 模块的存储内容
  // api.query.<pallet 名称>.<存储名称>
  const aliceAccount = await api.query.system.account(alice.address);
  console.log(`Alice Account: ${aliceAccount}`);
  const aliceFreeBalance = aliceAccount.data.free;
  console.log(`Alice free balance in readable format: ${aliceFreeBalance/TX_AMT}`)

  // 订阅 Alice 的帐号资料
  // 订阅事件
  const unsubAlice = await api.query.system.account(alice.address, aliceAcct => {
    console.log("Subscribed to Alice account.");
    const aliceFreeSub = aliceAcct.data.free;
    console.log(`Alice Account (sub): ${aliceFreeSub/TX_AMT}`);
  });
  const unsubBob = await api.query.system.account(bob.address, bobAcct => {
    console.log("Subscribed to Bob account.");
    const bobFreeSub = bobAcct.data.free;
    console.log(`Bob Account (sub): ${bobFreeSub/TX_AMT}`);
  });

  await sleep(1000);

  // 发送交易
  console.log("transfer from alice to bob");
  submitTx(api, alice, bob, TX_AMT);

  // 取得链上 meta-data. 去掉下面 comment 去看链上 meta-data. 是一个挺大的 JSON 文件
  // const metadata = await api.rpc.state.getMetadata();
  // console.log(`Chain Metadata: ${JSON.stringify(metadata, null, 2)}`);

  await sleep(LASTING_SECS * 1000);
  unsubAlice();
  unsubBob();
};

main()
  .then(() => {
    console.log("successfully exited");
    process.exit(0);
  })
  .catch(err => {
    console.log('error occur:', err);
    process.exit(1);
  })
