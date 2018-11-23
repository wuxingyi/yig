#YIG与IAM交互的API列表

##DescribeAccessKeys

通过AccessKeys获取SecretKeys

###请求格式:
YIG会发送一个POST请求到IAM, 主要输入是一组accessKeys以及用于YIG和IAM鉴权的一组key/secret(即X-Le-Key和X-Le-Secret).
```
POST / HTTP/1.1
Host: IAMEndpoint
X-Le-Key: key
X-Le-Secret: secret
content-type: application/json

body
{
"action": "DescribeAccessKeys",
"accessKeys":["hehehehe"]
}
```

####返回值格式:
```

{
    "total":1,
    "accessKeySet":[
        {
            "projectId":"p-abcdef",
            "name":"user1",
            "accessKey":"hehehehe",
            "accessSecret":"hehehehe",
            "status":"active",
            "updated":"2006-01-02T15:04:05Z07:00"
        }
    ]
}

```
可以返回多对key/secret组，其中核心字段是accessKey和accessSecret，其他几个字段是补充字段，如果没有就填空字符串。
