FROM node:20-slim AS builder

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"

RUN corepack enable

WORKDIR /src

COPY package.json pnpm-lock.yaml ./

RUN pnpm install --frozen-lockfile

COPY . ./

RUN pnpm run build

FROM node:lts-alpine

WORKDIR /src

COPY --from=builder /src/build ./build
COPY --from=builder /src/node_modules ./node_modules
COPY --from=builder /src/package.json ./package.json

EXPOSE 3000

ENTRYPOINT ["node", "build"]

