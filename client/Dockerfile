### Build
FROM node:15.11.0-alpine as builder

LABEL org.opencontainers.image.source="https://github.com/kostaspt/domane"

WORKDIR /usr/src/app

# Install dependencies
COPY package.json yarn.lock ./
RUN (yarn check --integrity && yarn check --verify-tree) || yarn install --frozen-lockfile --network-timeout 1000000

# Generate production build
COPY . .
RUN yarn build

# Remove dev dependencies
RUN yarn install --production

EXPOSE $CLIENT_PORT

CMD ["sh" , "-c", "yarn start -p ${CLIENT_PORT}"]
