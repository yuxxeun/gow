install:
    cd api && go mod tidy
    cp .env.example .env
    cd ../web && pnpm install
    cp .env.example .env

dev:
    cd web && pnpm dev

build:
    cd api && go build -o bin/api
    cd ../web && pnpm build

run-install:
    just install

run-build:
    just build
