# WineReviewGoa
trying out goa

From https://goa.design/learn/guide/, http://eng.rightscale.com/2015/12/11/goa-untangling-microservices.html
To generate code: `goagen bootstrap -d github.com/WineReviewGoa/design`

It generates all of this:
```
design
app
app/contexts.go
app/controllers.go
app/hrefs.go
app/media_types.go
app/user_types.go
app/test
app/test/bottle_testing.go
main.go
bottle.go
tool/cellar-cli
tool/cellar-cli/main.go
tool/cli
tool/cli/commands.go
client
client/client.go
client/bottle.go
client/user_types.go
client/media_types.go
swagger
swagger/swagger.json
swagger/swagger.yaml```
!!!