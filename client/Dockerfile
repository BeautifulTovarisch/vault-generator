FROM node:12.4 as development

ENV NODE_ENV=dev

WORKDIR /app

COPY package.json yarn.lock /app/

RUN yarn

COPY ./src /app/src

RUN yarn test --passWithNoTests

CMD [ "yarn", "start" ]
