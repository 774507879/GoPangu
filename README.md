# diskdb
 diskdb 暂时只支持从硬盘中加载数据。

需要说明的是，diskdb 没有实现  bitcask 模型的多个数据文件的机制，为了简单，我只使用了一个数据文件进行读写。但这并不妨碍你理解 bitcask 模型。

我写了一篇文章对 diskdb 进行讲解：[从零实现一个 k-v 存储引擎](https://mp.weixin.qq.com/s/s8s6VtqwdyjthR6EtuhnUA)

相信结合文章及 diskdb 的简单的代码，你能够快速上手了。

当然，你可以阅读 bitcask 模型的论文相关网址：[https://www.likecs.com/show-305625044.html](https://riak.com/assets/bitcask-intro.pdf)



