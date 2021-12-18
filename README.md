# Oura Ring OpenAPIv3 Schema Definition + Go Client
This repo contains an OpenAPIv3 schema definition for [OuraRing](https://ouraring.com/)

OuraRing API documentation https://cloud.ouraring.com/docs/

## Generating the client
Client generation is done using https://github.com/deepmap/oapi-codegen

To regenerate the client, run the command
```
oapi-codegen --package ouraring ouraring_openapi.yaml > ouraring.gen.go
```

## Extending the schema definition
At this time, the schema definition is not complete. 

If there is data which you would like to fetch but it is not available, consider adding it to the schema definition
and regenerating the client.

