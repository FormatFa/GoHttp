每个请求都由名字，地址，请求头，请求体组成，中间用换行符隔开，如
其中
- url后面一行是请求头
- 请求头后面加一行空行后的就是请求体

```
# get post
GET http://jsonplaceholder.typicode.com/posts


# create
POST https://jsonplaceholder.typicode.com/posts
Content-type: application/json; charset=UTF-8

{
	"title":"foo",
	"body":"bar",
	"userId":1
}
```