# Used to run under Windows

FROM openapitools/openapi-generator-cli:v7.12.0

ARG user_id
ARG group_id

RUN addgroup --gid $group_id app
RUN adduser --disabled-password --gecos '' --uid $user_id --gid $group_id app
USER app
