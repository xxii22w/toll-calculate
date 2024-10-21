## 通行费计算器
* 首先是安装依赖，kafka,protobuf,rpc,prometheus,logrus,websocket
### 项目由一下部分组成：
* obu 
--> 随即生成id和坐标位置来模拟通行位置
--> 把随机的id数据发送到websocket端点
* data_recevier 
--> 负责接收obu产生的数据，并推送到kafka
--> 从websocket接收数据并存储到fafka
* distance_calculator 
--> 负责消费kafka传来的数据并计算出距离
--> 使用客户端接口调用agg微服务的端点
* aggregator 
--> 聚合器 将数据汇总计算收费发票
--> 可以接受post请求来聚合obu的距离
--> 也可以根据生辰的obuid来获取发票
* gateway 
--> 网关 负责给外部提供一个接口来获取对于id的收费发票
### 运行的顺序如下
* make agg -> make receiver -> make calculate -> make obu -> make agg
*  make gateway 来测试服务是否可行
-> http://localhost:6000/invoice?obu=6534164431199076443

代码有不错的可拓展性，比如说middleware，里面可以自由拓展，并且如果想再添加中间件也很容易
store 也可以添加自己想用的数据库来替换，重写对于的接口就行

## 添加.env
AGG_HTTP_ENDPOINT=:4000
AGG_GRPC_ENDPOINT=:3001
AGG_STORE_TYPE=memory