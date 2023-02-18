FROM node:19-buster as builder

# executed from frontend folder
COPY frontend /usr/app

ARG VERSION

WORKDIR /usr/app

RUN \
    npm version --no-git-tag-version ${VERSION} && \
    npm run build

FROM scratch

COPY --from=builder /usr/app/build .
