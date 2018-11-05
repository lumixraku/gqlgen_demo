# gqlgen_demo

follow this https://gqlgen.com/getting-started/

## createTodo
```
mutation meme{
  createTodo(input: {
    text:"xxx",
    userId :"21312"
  }){
	id,
    text
  }
}

```
the response is
```
{
  "data": {
    "createTodo": {
      "id": "T5577006791947779410",
      "text": "xxx"
    }
  }
}

```

## get todo
```
# Try to write your query here
query mememomo{
  todos{
		id,
    text
  }
}

```
the response is
```
{
  "data": {
    "todos": [
      {
        "id": "T5577006791947779410",
        "text": "xxx"
      },
      {
        "id": "T8674665223082153551",
        "text": "cs"
      },
      {
        "id": "T6129484611666145821",
        "text": "uuuu"
      }
    ]
  }
}

```

结果集中再查询
```
query mememomo{
  todo(id: "T6334824724549167320"){
		id,
    text
    user(ID:123){
      id
    }
  }
}



```



```
# Try to write your query here
# mutation{
#   testtt(name:""){
#     id
#   }
# }


query{
  viewer{
    id
  }
}
```

```
mutation{

  createTodo(input: {
    text: "",
    userId: ""
  }){
    id
  }
}

```




## 关于继承
可以查看 meteor 分支

schema
```
union ConvertResult = Convert | AppConvert | XpathConvert
type ConvertEdge{
  node: ConvertResult!
  cursor: ID!
}


```
在 go 的实现中， ConvertResult是一个接口
然后之前定义的type Convert  AppConvert  XpathConvert 都实现了这个接口

```
type ConvertResult interface {
	IsConvertResult()
}



type AppConvert struct{...}
func (AppConvert) IsConvertResult() {}
```

查询的话要这么做
```
# mutation {
#   createJSConvert(input: {
#     name:"xxx",
#     externalAction:Phone
#   }) {
#     id
#   }
# }

query {
  converts(where: {
    advertiserID: "123"
  }){
    edges{
      node{
        ... on XpathConvert{
          id
          convertXpathUrl
        }
        ... on Convert{
          id
        }
        ...on AppConvert{
          id
          appId
        }
      }
    }
  }
}
```