FROM node:19-buster as builder

# executed from frontend folder
COPY frontend /usr/app

WORKDIR /usr/app
RUN npm run build

FROM scratch

COPY --from=builder /usr/app/build .
