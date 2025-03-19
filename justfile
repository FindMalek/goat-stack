init:
    pnpm install
    just packages/proto/init
    @echo "You are good to go!"

buf:
    rm -rf apps/server/proto
    rm -rf apps/web/src/gen/proto
    just packages/proto/buf


dev:
    pnpm dev

build:
    just apps/server/build
    pnpm build
