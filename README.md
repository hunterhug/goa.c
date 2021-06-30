# 数据结构和算法（Golang实现）

[![GitHub forks](https://img.shields.io/github/forks/hunterhug/goa.c.svg?style=social&label=Forks)](https://github.com/hunterhug/marmot/network)
[![GitHub stars](https://img.shields.io/github/stars/hunterhug/goa.c.svg?style=social&label=Stars)](https://github.com/hunterhug/marmot/stargazers)
[![GitHub last commit](https://img.shields.io/github/last-commit/hunterhug/goa.c.svg)](https://github.com/hunterhug/marmot)
[![GitHub issues](https://img.shields.io/github/issues/hunterhug/goa.c.svg)](https://github.com/hunterhug/marmot/issues)
[![License](https://img.shields.io/badge/license-Apache%202-4EB1BA.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

## 前言

数据结构和算法在计算机科学里，有非常重要的地位。此系列文章尝试使用 `Golang` 编程语言来实现各种数据结构和算法，并且适当进行算法分析。

系列文章首发于：

（🤔一直保持最新）Docsify风格的 [https://hunterhug.github.io/goa.c](https://hunterhug.github.io/goa.c) 。

 <del>（可能🍒过时）Gitbook风格的 [https://goalgorithm.github.io](https://goalgorithm.github.io) 。 </del>

**目前有部分网站，未经许可分发了此系列内容，存在更新延后的风险，之前存在部分的笔误，都已经更正了，请以此网站为准！**

联系作者： [https://github.com/hunterhug](https://github.com/hunterhug)

## 目录

* [目录](README.md)
* [前言](basic/before.md)
* [简单入门Golang](golang/README.md)
    * [包、变量和函数](golang/basic.md)
    * [流程控制语句](golang/logic.md)
    * [结构体和方法](golang/struct.md)
    * [接口](golang/interface.md)
    * [并发、协程和信道](golang/concurrent.md)
    * [标准库](golang/lib.md)
* [基础知识](basic/README.md)
    * [分治法和递归](basic/rescuvie.md)
    * [算法复杂度及渐进符号](basic/dregee.md)
    * [算法复杂度主方法](basic/master_method.md)
    * [延伸-计算理论：P和NP问题](basic/p.md)   
* [常见数据结构及算法](algorithm/README.md)
    * [链表](algorithm/link.md)
    * [可变长数组](algorithm/array_change.md)
    * [栈和队列](algorithm/stack_queues.md)
    * [列表](algorithm/list.md)
    * [字典](algorithm/dict.md)
    * [树](algorithm/tree.md)
    * [排序算法](algorithm/sort.md)
        * [冒泡排序](algorithm/sort/bubble_sort.md)
        * [选择排序](algorithm/sort/select_sort.md)
        * [插入排序](algorithm/sort/insert_sort.md)
        * [希尔排序](algorithm/sort/shell_sort.md)
        * [归并排序](algorithm/sort/merge_sort.md)
        * [优先队列及堆排序](algorithm/heaplike/heaps.md)
        * [快速排序](algorithm/sort/quick_sort.md)
    * [查找算法](algorithm/search.md)
        * [哈希表：散列查找](algorithm/search/hash_find.md)
        * [二叉查找树](algorithm/search/bs_tree.md)
        * [AVL树](algorithm/search/avl_tree.md)
        * [2-3树和左倾红黑树](algorithm/search/llrb_tree.md)
        * [2-3-4树和普通红黑树](algorithm/search/rb_tree.md)
* [文档部署](doc/install.md)
* [书籍推荐](doc/book.md)
* [参考资料](basic/refer.md)

希望在本地阅读，请安装 `Docker` 后，克隆仓库到本地执行：

```
./docker_build_docsify.sh
./docker_run_docsify.sh
```

打开 [http://127.0.0.1:12346](http://127.0.0.1:12346) 阅读。

## 作者寄语

学而不思则罔，思而不学则殆。

意思是说，学习之后如果不做自己的思考，那么会很迷罔，没有收获，成为一个工具人，但是如果你天天思考，而不学习，那么你就会很疑惑，因为你不知道你是对的还是错的，你需要去向其他人学习，去吸收其他人已经留存的知识。

学习离不开思考，思考也不能脱离学习，二者相辅相成，缺一不可，这是学习的最基本方法。

## 一起参与

如何建议和贡献自己的知识库，可以前往 [https://github.com/hunterhug/goa.c](https://github.com/hunterhug/goa.c) 的仓库提 `PR` 和 建议。

开源书籍的评论使用的是 `GitTalk`，可以打开 [网站](https://hunterhug.github.io/goa.c) 阅读后评论自己的心得。

## 章节代码

所有章节的代码可以在 [仓库的code文件夹中](https://github.com/hunterhug/goa.c/tree/master/code) 找到。

## 赞助作者

如果你想赞助作者买根辣条，可以扫描下方的二维码:

![/weixin.png](./weixin.png)

## License

本开源书籍分发使用 `Apache License`：

```
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
