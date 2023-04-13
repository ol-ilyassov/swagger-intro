# swagger-intro

Some practice of how to add Swagger documentation into project.

- Source repo of Swagger generator:
https://github.com/swaggo/swag

### GIN Swagger:

- Install tools: 

go get -u github.com/swaggo/swag

go install github.com/swaggo/swag/cmd/swag@latest

- Command to run generation:

/home/user/go/bin/swag init --output ../../docs

- Code required to add into project:

router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

- Code dependencies (gin only):

swaggerFiles "github.com/swaggo/files"

"github.com/swaggo/gin-swagger"

_ "swagger-intro/docs"

- Get Access:

[http://localhost:4000/swagger/index.html](http://localhost:4000/swagger/index.html)
