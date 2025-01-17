FROM alpine:3.18.4

ARG GIT_HASH
ARG GIT_BRANCH
LABEL gitHash=$GIT_HASH gitBranch=$GIT_BRANCH

WORKDIR /
COPY cmd/{{ToKebab .Interface.ServerName}}/{{ToKebab .Interface.ServerName}} /{{ToKebab .Interface.ServerName}}

CMD ["/{{ToKebab .Interface.ServerName}}"]
