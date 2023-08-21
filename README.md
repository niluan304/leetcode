test case sdk for leetcode
                 
## TODO
- 读取配置文件
  - [ ] 文件保存位置
  - [ ] 语言列表
- 文件生成
  - [x] 读取命令行参数 `titleSlug`
  - [x] `case.go` 响应里有函数的请求/响应类型
    - [ ] 完善`class`
  - [x] `solution_test.go`
  - [x] 通过 `titleSlug` 获取题目 `frontendId`和 `categoryTitle`
  - [x] 通过 `frontendId` 获取题目完整信息
  - [ ] 致谢