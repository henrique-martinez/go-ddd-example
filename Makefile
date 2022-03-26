openapi:
	oapi-codegen -package ports -generate types api/petstore.yml > internal/petstore/ports/petstore_types.gen.go
	oapi-codegen -package ports -generate chi-server api/petstore.yml > internal/petstore/ports/petstore.gen.go