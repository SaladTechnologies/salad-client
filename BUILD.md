# Building

```pwsh
Set-Location -Path salad-client
docker container run --name openapi-generator --rm --volume "$(Get-Location):/local" --workdir /local saladtechnologies/openapi-generator-cli generate -c config.json -g go --git-user-id SaladTechnologies --git-repo-id salad-client
```
