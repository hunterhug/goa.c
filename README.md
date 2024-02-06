# 数据结构和算法（Golang实现）

[![GitHub stars](https://img.shields.io/github/stars/hunterhug/goa.c.svg?style=social&label=Stars)](https://github.com/hunterhug/goa.c/stargazers)
[![GitHub last commit](https://img.shields.io/github/last-commit/hunterhug/goa.c.svg)](https://github.com/hunterhug/goa.c)
[![GitHub issues](https://img.shields.io/github/issues/hunterhug/goa.c.svg)](https://github.com/hunterhug/goa.c/issues)
[![License](https://img.shields.io/badge/license-Apache%202-4EB1BA.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

心灵一问：在面试中，你是否对面试官喋喋不休，反复问你关于快速排序，红黑树的各种细节问题而感到无奈，在工作多年后，已经在工程实践中岁月蹉跎的你，想重拾数据结构和算法的回忆，不妨读读这本书。

> 盛年不重来，一日难再晨，及时当勉励，岁月不待人。

简单总结：作者通过几个基础的章节教大家入门 `Golang` 编程语言，然后开始介绍计算机科学与技术中基本的方法论，如分治法，递归和算法复杂度，紧接着从基本数据结构开始，慢慢介绍到排序算法和查找算法。

## 前言

数据结构和算法在计算机科学里，有非常重要的地位。此系列文章尝试使用 `Golang` 编程语言来实现各种数据结构和算法，并且适当进行算法分析。

系列文章首发于：

（🧍‍♂️原始文档托管）Github 代码仓库： [https://github.com/hunterhug/goa.c](https://github.com/hunterhug/goa.c) 。

（🤔一直保持最新）Docsify 风格的网站： [https://hunterhug.gitlab.io/goa.c](https://hunterhug.gitlab.io/goa.c) 。

**目前有部分网站，未经许可分发了此系列内容，存在更新延后的风险，之前存在部分的笔误，都已经更正了，请以此网站为准！**

本地[离线阅读](doc/install.md)，执行： 


```bash
docker run --name algorithm -d -p 12346:3000 hunterhug/algorithm:docsify
```

浏览器打开 [http://127.0.0.1:12346](http://127.0.0.1:12346) 。

联系作者： [https://github.com/hunterhug](https://github.com/hunterhug) ，邮箱： gdccmcm14@live.com。

## 目录

* [目录](README.md)
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
        * [优先队列及堆排序](algorithm/heaps.md)
        * [快速排序](algorithm/sort/quick_sort.md)
    * [查找算法](algorithm/search.md)
        * [哈希表：散列查找](algorithm/search/hash_find.md)
        * [二叉查找树](algorithm/search/bs_tree.md)
        * [AVL树](algorithm/search/avl_tree.md)
        * [2-3树和左倾红黑树](algorithm/search/llrb_tree.md)
        * [2-3-4树和普通红黑树](algorithm/search/rb_tree.md)
* [文档部署](doc/install.md)
* [书籍推荐](doc/book.md)
* [参考](basic/refer.md)

## 作者寄语

学而不思则罔，思而不学则殆。

意思是说，学习之后如果不做自己的思考，那么会很迷罔，没有收获，成为一个工具人，但是如果你天天思考，而不学习，那么你就会很疑惑，因为你不知道你是对的还是错的，你需要去向其他人学习，去吸收其他人已经留存的知识。

学习离不开思考，思考也不能脱离学习，二者相辅相成，缺一不可，这是学习的最基本方法。

## 一起参与

如何建议和贡献自己的知识库，可以前往 [https://github.com/hunterhug/goa.c](https://github.com/hunterhug/goa.c) 的仓库提 `PR` 和 建议。

开源书籍的评论使用的是 `GitTalk`，可以打开 [https://hunterhug.gitlab.io/goa.c](https://hunterhug.gitlab.io/goa.c) 阅读后评论自己的心得。

## 章节代码

所有章节的代码可以在 `Github` 仓库根目录下的 [https://github.com/hunterhug/goa.c/tree/master/code](https://github.com/hunterhug/goa.c/tree/master/code) 文件夹中找到。

## 赞助作者

如果你想赞助作者买根辣条，可以打开微信扫描下方的二维码:

![/weixin.png](./picture/weixin.png)

赞助情况（排名不分先后）：

| 日期         | 组织/个人       | 说明        |
|------------|-------------|-----------|
| 2023.09.14 | 长江          | 6.66 RMB  |
| 2023.08.22 | 无名          | 6.66 RMB  |
| 2023.04.15 | min         | 6.66 RMB  |
| 2022.09.18 | 放下          | 66.66 RMB |
| 2022.07.22 | Perry       | 6.66 RMB  |
| 2022.07.21 | 九久捌         | 0.50 RMB  |
| 2022.05.07 | 牛肉拌面.       | 4.50 RMB  |
| 2022.04.21 | cy          | 4.50 RMB  |
| 2022.04.13 | cy          | 1.50 RMB  |
| 2022.04.02 | skywalker   | 9.50 RMB  |
| 2022.03.13 | 小桀          | 4.50 RMB  |
| 2022.03.02 | 古寒飞         | 9.50 RMB  |
| 2022.02.08 | 罗博贤         | 9.50 RMB  |
| 2022.01.21 | 匿名          | 1.50 RMB  |
| 2021.11.11 | peter锦锋     | 18.50 RMB |
| 2021.09.18 | 古寒飞         | 200.00 RMB |
| 2021.08.31 | crypto ivil | 9.50 RMB  |
| 2021.08.19 | 礼服黑面侠       | 9.50 RMB  |
| 2020.09.22 | 胡小东         | 9.50 RMB  |
| 2020.06.27 | 杨某          | 99.99 RMB |

感谢她（他）们给作者送去温暖，让作者有更好的动力边吃辣条边写文章！

## 广告位招租

广告位招租板块。

### 开源项目

- 💐 Memory Cache Implement By Golang: [https://github.com/hunterhug/gocache](https://github.com/hunterhug/gocache)
- 💐 Log The World Very Easy With Zap: [https://github.com/hunterhug/golog](https://github.com/hunterhug/golog)
- 💐 Marmot A Golang HTTP Download: [https://github.com/hunterhug/marmot](https://github.com/hunterhug/marmot)
- 💐 Redis Distributed Lock By Golang: [https://github.com/hunterhug/gorlock](https://github.com/hunterhug/gorlock)

### 特别推荐

推荐中华书籍阅读：[红楼梦](https://hunterhug.github.io/china-literary/%E7%BA%A2%E6%A5%BC%E6%A2%A6/%E7%BA%A2%E6%A5%BC%E6%A2%A6.html)，[西游记](https://hunterhug.github.io/china-literary/%E8%A5%BF%E6%B8%B8%E8%AE%B0/%E8%A5%BF%E6%B8%B8%E8%AE%B0.html)，[水浒传](https://hunterhug.github.io/china-literary/%E6%B0%B4%E6%B5%92%E4%BC%A0/%E6%B0%B4%E6%B5%92%E4%BC%A0.html)，[三国演义](https://hunterhug.github.io/china-literary/%E4%B8%89%E5%9B%BD%E6%BC%94%E4%B9%89/%E4%B8%89%E5%9B%BD%E6%BC%94%E4%B9%89.html)，[史记](https://hunterhug.github.io/china-history/%E5%8F%B2%E8%AE%B0/%E5%8F%B2%E8%AE%B0.html)。


## 版权所有

本开源书籍分发使用 `Apache License`，您可以随意传阅，但请保留以下声明：

```
Copyright [2019-2022] [github.com/hunterhug]

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
