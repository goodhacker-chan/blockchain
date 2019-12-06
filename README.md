# 简介
> 项目是根据比特币的思路(简化版,jeiwan的教程)实现了一个虚拟货币区块链, 后续会在此基础上持续更新新功能和智能合约的基础应用

**开发语言**: Golang

**50米大刀的温馨提示: 此项目仅供学习参考, 不喜勿喷如有疑问或建议请私信哦!**

# CLI应用使用方法
```
createwallet - 创建一个钱包并将秘钥存放到钱包文件, 返回钱包地址

getbalance -address ADDRESS - 获取钱包地址余额

createblockchain -address ADDRESS - 创建一个区块

listaddresses - 显示钱包文件中的所有钱包地址

printchain - 打印出区块链结构

reindexutxo - 重新构建UTXO集

send -from FROM -to TO -amount AMOUNT -mine(可选参数) - 将一定数量的硬币从一个地址发送到另一个地址,在同一个节点上时设置 -mine(标志表示该块将被同一节点立即开采)

startnode -miner ADDRESS - 启动具有在NODE_ID env中指定的ID的节点, -miner实现采矿, ADDRESS 矿工地址
```
**mac和Linux用户记得带上路径符 程序目录下使用 './blockchain' 或者使用绝对路径**

- `network` 分支为p2p网络测试分支, 模拟p2p网络, 测试教程见下方`3行`处
- `network-p2p` 此分支为p2p网络正式分支, 可以实现简单的p2p网络

# 举个P2P测试网络栗子
1. 在程序目录下新开一个终端窗口, 设置下网络节点 (作为中央节点)
    ```
    $ export NODE_ID=3000
    ```
2. 首先创建一个钱包地址
    ```
    $ ./blockchain createwallet
    
    你的新地址: 1PVdF8pExHnn4vaZ65SyggwKgJ2e63BtZp(此处为你自己生产的地址)(地址0)
    ```
3. 创建一个区块链, 这时创建区块链的地址将会有10个奖励币
    ```
    $ ./blockchain createblockchain -address 1PVdF8pExHnn4vaZ65SyggwKgJ2e63BtZp
    ```
4. 拷贝下区块链的初始块
    ```
    $ cp blockchain_3000.db blockchain_genesis.db 
    ```
5. 新开一个终端窗口, 设置一个新节点(作为SPV节点)
    ```
    $ export NODE_ID=3001
    ```
6. 创建几个钱包地址用来收币, 下方地址*3
    ```
    $ ./blockchain createwallet (*3)
    
    你的新地址: 1GkU5gHSDvFiXz5UsFPJ9vd2Q9unkHfG42   (地址1)
    
    你的新地址: 1AhHn4nTgzqcbZriR11yTjMu5zqvAGdVJC   (地址2)
    
    你的新地址: 1G9DqgDVQMX3VoRu2h7PKJPkp4BU9EECGj   (地址3)
    ```
7. 使用3000节点终端窗口的矿工地址给收款地址交易币
    ```
    $ ./blockchain send -from 地址0 -to 地址1 -amount 3 -mine
    
    $ ./blockchain send -from 地址0 -to 地址2 -amount 3 -mine
    
    // 由于初始区块链没有矿工 所以交易加上 -mine 参数立即开采
    ```
8. 在中央节点3000 上启动节点, 这个节点需要一直运行
    ```
    $ ./blockchain startnode
    ```
9. 3001节点上 使用上边的初始块, 然后运行节点推送和拉取块, 每当有交易后运行此节点
    ```
    $ ./blockchain startnode
   
    收到 block 命令
    收到一个新的块!
    添加块 0000076a75e686f67d7fccfe5df1bc85a906492aa783fa7e4de280a30e611982
    ```
10. 关闭3001节点(ctrl + c), 查看钱包的余额
    ```
    $ ./blockchain getbalance -address 地址0
    钱包地址余额 '地址0': 4
    
    $ ./blockchain getbalance -address 地址1
    钱包地址余额 '地址1': 3
    
    $ ./blockchain getbalance -address 地址2
    钱包地址余额 '地址1': 3
    ```
11. 测试矿工节点挖矿, 新开一个终端窗口 设置一个矿工节点
    ```
    $ export NODE_ID=3002   (矿工节点)
    ```
12. 使用初始块初始化矿工节点, 并启动矿工节点
    ```
    // 初始化矿工节点
    $ cp blockchain_genesis.db blockchain_3002.db
    
    // 启动矿工节点
    $ ./blockchain startnode -miner 地址0(一个启动中的节点地址)
    起始节点 3002
    采矿进行中... 收益转入到地址:  地址0
    ```
13. 使用SPV节点3001 进行一些交易
    ```
    $ ./blockchain send -from 地址1 -to 地址3 -amount 2 
    ```
14. 切换到矿工节点可以看到矿工正在处理交易并挖掘新块
    ```
    收到 inv 命令
    收到新库存 1 tx
    type: tx 
    收到 tx 命令
    处理事务
    afddd386c91bd731e4a66cab558d7c6176e62a145cd2950e9a0463eddcd9f9e4
    新块开采!
    收到 getdata 命令
    ```
15. 启动SPV3001节点拉取新块, 并查看各个地址的余额
    ```
    $ ./blockchain startnode
    
    // 查询地址余额
    $ ./blockchain getbalance -address 地址0(矿工节点有收益加成)
    钱包地址余额 '地址0': 34
    
    $ ./blockchain getbalance -address 地址1
    钱包地址余额 '地址1': 1
    
    $ ./blockchain getbalance -address 地址2
    钱包地址余额 '地址2': 3
    
    $ ./blockchain getbalance -address 地址3
    钱包地址余额 '地址3': 2
    ```
**钱包地址记得替换成自己创建的地址**

