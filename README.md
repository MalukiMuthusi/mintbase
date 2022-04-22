# Mintbase API

This project is an example for building an API using Swagger. To generate the and deploy the documentation, and implement the server.

## Generate API documentation from the swagger.yml file

Download the [swagger-codegen-cli-2.4.27.jar][9] file. You can also follow the installation guide from the swagger-codegen [documentation][10]

```sh
# Download current stable 2.x.x branch (Swagger and OpenAPI version 2)
wget https://repo1.maven.org/maven2/io/swagger/swagger-codegen-cli/2.4.27/swagger-codegen-cli-2.4.27.jar -O swagger-codegen-cli.jar
```

Install java. The swagger-codegen binary is a java tar file. Use [sdkman][11] to easily install java

```sh
sdk install java
```

Now use the `swagger-codegen`

```sh
java -jar swagger-codegen-cli.jar --help
```

### Generate the documentation

Generate a swagger.json file from the swagger.yml file first.

```sh
java -DmodelDocs=false \
 -jar $HOME/programs/swagger-codegen-cli-2.4.27.jar generate \
 -i swagger.yml \
 -l swagger \
 -o api-docs


```

```sh
java \
    -jar $HOME/programs/swagger-codegen-cli-2.4.27.jar generate \
    -i api-docs/swagger.json \
    -l html2 \
    -o api-docs
```

## References

- OpenAPI Version 2.0 [specification][1]
- Swagger Editor [Swagger-editor-github][2]
- Link to the Swagger editor hosted online for use [Swagger-editor][3]
- OpenAPI vscode editor [Open-API-marketplace-link][4]
- Swagger Codegen [Swagger-codegen-github][5]
- Golang Gin Web Framework [documentation][6]
- Go Text template [documentation][7]
- HTTP status codes [HTTP-Status-Codes-Website][8]
- Use SDK Manager to install the development tools, used by Swagger tools [sdkman-website][11]

[1]: https://swagger.io/specification/v2/
[2]: https://github.com/swagger-api/swagger-editor
[3]: https://editor.swagger.io
[4]: https://marketplace.visualstudio.com/items?itemName=42Crunch.vscode-openapi
[5]: https://github.com/swagger-api/swagger-codegen
[6]: https://github.com/gin-gonic/gin
[7]: https://pkg.go.dev/text/template
[8]: https://restfulapi.net/http-status-codes
[9]: https://repo1.maven.org/maven2/io/swagger/swagger-codegen-cli/2.4.27/swagger-codegen-cli-2.4.27.jar
[10]: https://github.com/swagger-api/swagger-codegen
[11]: https://sdkman.io/sdks
