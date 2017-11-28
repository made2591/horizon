### STAGE 1: Build ###

# The builder node
FROM node:8.1.4-alpine as builder

COPY src/package.json src/package-lock.json ./

## Storing node modules on a separate layer will prevent unnecessary npm installs at each build
RUN npm i && mkdir /ng-app && cp -R ./src/node_modules ./ng-app

WORKDIR /ng-app

COPY ./src .

## Build the angular app in production mode and store the artifacts in dist folder
RUN $(npm bin)/ng build --prod

### STAGE 2: Setup

FROM smebberson/alpine-nginx as runner

COPY nginx.conf /etc/nginx/nginx.conf

## From 'builder' stage copy over the artifacts in dist folder to default nginx public folder
COPY --from=builder /ng-app/dist/* /usr/html/
