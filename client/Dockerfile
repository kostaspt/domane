### Build
FROM node:15.5.0-alpine as builder

LABEL org.opencontainers.image.source https://github.com/kostaspt/domane

WORKDIR /usr/src/app

# Install dependencies
COPY package.json yarn.lock ./
RUN (yarn check --integrity && yarn check --verify-tree) || yarn install --frozen-lockfile --network-timeout 1000000

# Generate production build
COPY . .
RUN yarn build

# Remove dev dependencies
RUN yarn install --production

### Serve
FROM node:15.5.0-alpine

WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/package.json ./package.json
COPY --from=builder /usr/src/app/node_modules ./node_modules
COPY --from=builder /usr/src/app/.next ./.next
COPY --from=builder /usr/src/app/public ./public

EXPOSE $CLIENT_PORT

CMD ["yarn", "start", "-p", "${CLIENT_PORT}"]
