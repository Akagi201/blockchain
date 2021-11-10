# solution05

1. 列出 3 个常用的宏、3 个常用的存储数据结构

常用宏：

```Rust
pallet::config // 定义配置接口
pallet::storage // 存储单元
pallet::event // 事件
```

数据结构：

```Rust
Vec<T>, BTreeMap, BTreeSet
```

2. 实现存证模块的功能，包括：创建存证；撤销存证。

[subchain poe pallet](https://github.com/Akagi201/subchain/tree/master/pallets/poe)

3. 为存证模块添加新的功能，转移存证，接收两个参数，一个是内容的哈希值，另一个是存证的接收账户地址。

[subchain poe pallet](https://github.com/Akagi201/subchain/tree/master/pallets/poe)
