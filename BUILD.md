# Building

# Windows

Generate the Docker image if necessary using the included Dockerfile.

```pwsh
Set-Location -Path salad-client
docker container run --name openapi-generator --rm --volume "$(Get-Location):/local" --workdir /local saladtechnologies/openapi-generator-cli generate -c config.json -g go --git-user-id SaladTechnologies --git-repo-id salad-client
```

# Non-Windows

There is a Makefile with a build target that basically does this:

```bash
docker run \
    --name openapi-generator \
    --rm \
    --volume "${PWD}:/local" \
    --workdir /local \
    openapitools/openapi-generator-cli:v7.12.0 generate \
        -c config.json \
        -i openapi.yaml \
        --git-user-id SaladTechnologies \
        --git-repo-id salad-client
```

Replace `openapi.yaml` with whatever your source for the spec file is if
different; it may be a URL.

With the included Makefile the process now looks more like this:

```bash
# update api/openapi.yaml to the desired version
# and build it
make build
```

To change the spec file here add it to the make command line:
```bash
make build OPENAPI_SPEC=openapi.yaml
```
