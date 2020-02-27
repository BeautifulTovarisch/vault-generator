FROM node:12.4 as development

ENV NODE_ENV=dev

WORKDIR /app

COPY package.json yarn.lock /app/

RUN yarn

COPY . /app

RUN yarn test --passWithNoTests

CMD [ "yarn", "start" ]

# Production Build ##############################

FROM node:12.4-alpine as build

ENV NODE_ENV=prod

WORKDIR /app

COPY package.json yarn.lock /app/

# Need toolchain still
RUN yarn install --prod=false

COPY . /app

# Tests mandatory at this point
RUN yarn test
RUN yarn build

# Production ####################################

FROM nginx:1.17.6-alpine as prod

COPY --from=build /app/dist /var/www/vault-generator.com
COPY nginx.conf /etc/nginx

CMD [ "nginx", "-g", "daemon off;" ]
