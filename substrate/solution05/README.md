# solution05

1. 列出 3 个常用的宏、3 个常用的存储数据结构

常用宏：

使用 Substrate 进行 Runtime 模块开发的过程中，常用的宏：
frame_support::pallet 定义功能模块
pallet::config 定义配置接口
pallet::storage 存储单元
pallet::event 事件
pallet::error 错误信息
pallet::call 包含可调用函数
pallet::hooks 区块不同时期的执行逻辑
可以参考 <https://docs.substrate.io/v3/runtime/macros/>

常用的存储数据结构：
单值 StorageValue
映射 StorageMap
双键映射 StorageDoubleMap
可以参考 <https://docs.substrate.io/v3/runtime/storage/>
及 <https://docs.substrate.io/rustdocs/latest/frame_support/storage>

key 的 hash 算法：Blake2_128Concat(密码学安全), Twox64Concat(速度更快，但不是密码学安全的), Identity(通常用在 key 本身就是 hash 结果时，避免不必要的计算)
StorageMap 类型，用来保存键值对，单值类型都可以用作 key 或者 value。
插入一个元素：MyMap::insert(key, value);
通过 key 获取 value: MyMap::get(key);
删除某个 key 对应的元素：MyMap::remove(key);
覆盖或者修改某个 key 对应的元素：
MyMap::insert(key, new_value);
MyMap::mutate(key, |old_value| old_value+1);

存储功能设计原则及需要注意的地方，如数值运算时不能使用浮点数等，可以参考 https://docs.substrate.io/v3/advanced/storage/
最后记得加载开发的功能模块，construct_runtime 添加模块到 Runtime

StorageMap 类型，保存键值对
#[pallet::storage]
#[pallet::getter(fn my_map)]
pub type MyMap<T> = StorageMap<
    _,
    Blake2_128Concat,
    u8,
    Vec<u8>,
>;
key 是 u8，值是 Vec， Blake2_128Concat 是 Hash 算法，对 key 进行 hash 计算后的结果作为 kv 数据库的真正存储 key。


1. 实现存证模块的功能，包括：创建存证；撤销存证。

[subchain poe pallet](https://github.com/Akagi201/subchain/tree/master/pallets/poe)

3. 为存证模块添加新的功能，转移存证，接收两个参数，一个是内容的哈希值，另一个是存证的接收账户地址。

[subchain poe pallet](https://github.com/Akagi201/subchain/tree/master/pallets/poe)
