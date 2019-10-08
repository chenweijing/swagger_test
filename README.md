# swagger 使用demo

## swagger的使用参数形式

```
    https://razeencheng.com/post/go-swagger.html
    @Param 1.参数名 2.参数类型 3.参数数据类型 4.是否必须 5.参数描述 6.其他属性
```

### 参数类型

```
// 参数类型
// path 	该类型参数直接拼接在URL中，如Demo中HandleGetFile：
// query 	该类型参数一般是组合在URL中的，如Demo中HandleHello
// formData 该类型参数一般是POST,PUT方法所用，如Demo中HandleLogin
// body 	当Accept是JSON格式时，我们使用该字段指定接收的JSON类型
```