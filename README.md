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