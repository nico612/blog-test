## 代码测试

控制层、业务层和仓库层之间是通过接口来通信的。通过接口通信有一个好处，就是可以让各层变得可测。那接下来，我们就来看下如何测试各层的代码。

1. 模型层：因为模型层不依赖其他任何层，我们只需要测试其中定义的结构及其函数和方法即可。

2. 控制层：控制层依赖于业务层，意味着该层需要业务层来支持测试。你可以通过 golang/mock 来 mock 业务层，测试用例可参考 TestPostController_Create。

3. 业务层： 因为该层依赖于仓库层，意味着该层需要仓库层来支持测试。你可以通过 golang/mock 来 mock 仓库层，测试用例可以参考 Test_userBiz_List。

4. 仓库层： 仓库层依赖于数据库，如果调用了其他微服务，那还会依赖第三方服务。我们可以通过 [sqlmock](https://github.com/DATA-DOG/go-sqlmock) 来模拟数据库连接，通过 [httpmock](https://github.com/jarcoal/httpmock) 来模拟 HTTP 请求。