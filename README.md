# diskdb
 diskdb 暂时只支持从硬盘中加载数据。

需要说明的是，diskdb 没有实现  bitcask 模型的多个数据文件的机制，为了简单，我只使用了一个数据文件进行读写。但这并不妨碍你理解 bitcask 模型。

当然，你可以阅读 bitcask 模型的相关网址：[https://www.likecs.com/show-305625044.html](https://www.likecs.com/show-305625044.html)



