# LeetCode刷题整理

## 算法 Algorithm

总结自己学习过的数据结构和算法，以及按照题目类型分类。

根据[灵神的说法](https://github.com/EndlessCheng/codeforces-go/blob/master/README.md)：
> 一个算法模板应当涵盖以下几点：
> - 对该算法的基本介绍（核心思想、复杂度等） 
> - 参考链接或书籍章节（讲的比较好的资料） 
> - 模板代码（可以包含一些注释、使用说明） 
> - 模板补充内容（常见题型中的额外代码、建模技巧等） 
> - 相关题目链接（模板题、经典题、思维转换题等）
> 

那这个仓库的目标是：用自己的代码风格，实现常见的数据结构和算法模板。

实现过程：
- 通过「默写」代码，了解模板中的细节，增加套用模板的熟练度，避免bug。
- 通过「阅读」他人优秀的代码，提高自身代码的健壮性和可读性。
- 通过「死记」解题代码，串联其他类似的题目

### 数据结构
- [集合 set.go](./container/set.go)
- 堆
  - [最小堆/最大堆 heap.go](./container/heap.go) 
  - [优先队列 heap_pq.go](./container/heap_pq.go)
- 线段树
  - [线段树 segment_tree.go](./container/seg_tree/segment_tree.go)
  - [延迟标记线段树 lazy_segment_tree.go](./container/seg_tree/lazy_segment_tree.go)
- 树
  - [红黑树 rbtree.go](./container/tree/rbtree/tree.go) 

### 动态规划
- [记忆化搜索 memory_search.go](./copypasta/dp/memory_search.go) 实现了类似Python的cache装饰器
- [背包 knapsack.go](./copypasta/dp/knapsack.go)
  - 0-1 背包
  - 完全背包
  - 分组背包
- 线性 DP
  - [LCS 最长公共子序列 lcs.go](./copypasta/dp/lcs.go)
  - [LIS 最长递增子序列 lis.go](./copypasta/dp/lis.go)

### 图
  - [最短路径 distance.go](./copypasta/graph/distance.go)
    - Floyd


## 代码生成

为了方便在IDE中解题，编写了脚本用于生成LeetCode的代码模板。

-[x] 生成每日一题的代码模板
-[x] 根据 「题号」 或 「题名」problem slug 生成代码模板
  - LCR等前缀题目，仅支持通过题名生成模板
-[x] 支持本地运行单元测试和自定义测试用例。
-[x] 查看题解的原始 markdown
-[ ] 交互式查看题解

### 自定义代码模板
- 根据[「配置文件」](./config.yml) 自定义生成的文件名
- 根据[「模板文件」](./template) 自定义文件内容
  - 模板文件使用 go模板语法