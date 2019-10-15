[TOC]
# 简介
> 项目是根据比特币的思路(简化版)实现了一个虚拟货币区块链, 后续会在此基础上持续更新新功能和智能合约的基础应用

**开发语言**: Golang

==提示:项目仅供学习参考, 如有疑问或建议请私信哦!==

# 应用使用方法
```
blockchain getbalance -address ADDRESS - 获取钱包地址余额

blockchain createblockchain -address ADDRESS - 创建一个区块

blockchain createwallet - 创建一个钱包并将秘钥存放到钱包文件, 返回钱包地址

blockchain listaddresses - 显示钱包文件中的所有钱包地址

blockchain printchain - 打印出区块链结构

blockchain send -from FROM -to TO -amount AMOUNT - 交易发送给一个地址金币
```
==mac和Linux用户记得带上路径符 程序目录下使用 './blockchain' 或者使用绝对路径==