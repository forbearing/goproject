项目需求
    模拟实现基于文本界面的《客户信息管理软件》
    该软件能够实现对客户对象的插入、修改和删除（用切片实现），并能够打印客户明细表

主要涉及以下知识点：
    - 切片的插入、删除和替换
    - 多对象协同工作


customerView：界面
    1、可以显示界面
    2、可以接收用户的输入
    3、根据用户的输入，调用 customerService 的方法完成客户的管理：修改、删除、显示等等。

customerService：处理逻辑业务逻辑
    1、完成对用户的各种操作
    2、对客户的增、删除、 修改、显示

customer：表示数据，model 层
    1、表示一个客户
    2、客户各种字段
    customer 表示一个客户信息，因此他需要包含客户必须的各种字段
        1、姓名
        2、年龄
        3、性别
        4、电话
        5、邮箱

customeView
1. list 去调用 customerService 的 List 方法，并显示客户列表
2. add 方法去调用 customerService 的 Add 方法，完成客户添加
3. delete 方法调用 customeService 的 Delete 方法，完成客户删除
4. update 方法调用 customeServcie 的 Update 方法，完成客户的修改

customerService
1. 编写 NewCustomerService 返回一个 customerService 实例
2. 编写 List 方法：返回客户信息，其实就是切片
3. 编写 Add 方法，将新的客户加入到 customers 切片
4. 编写 Delete 方法，接收id，删除一个客户
   编写 FindById 方法，接收id，返回 id 号对应的切片的下标
5. 编写 Update 方法，接收 id，修改 id 对应的数据

customer
1. 字段
    Id
    Name
    Gender
    Age
    Phone
    Email
2. 方法
    GetInfo

    
