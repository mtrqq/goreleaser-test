FROM node:19-alpine3.16

COPY bundle /usr/app/bundle

RUN npm install --global serve

ENTRYPOINT [ "serve", "-s", "/usr/app/bundle" ]
