FROM node:latest AS base
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable



FROM base AS install
WORKDIR /app/
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

RUN \
    --mount=type=bind,target=package.json,source=package.json \
    --mount=type=bind,target=pnpm-lock.yaml,source=pnpm-lock.yaml \
    --mount=type=bind,target=pnpm-workspace.yaml,source=pnpm-workspace.yaml \
    --mount=type=bind,target=apps/web/package.json,source=apps/web/package.json \
    --mount=type=bind,target=packages/ui/package.json,source=packages/ui/package.json \
    pnpm install --frozen-lockfile --prod

FROM base AS build
WORKDIR /app/
COPY \
    --link \
    "." "/app/"

COPY \
    --from=install \
    --link \
    "/app/node_modules" "/app/node_modules"

RUN pnpm build


FROM nginx:stable-alpine
WORKDIR /app/

# Taking advantages from docker multi-staging, we copy our newly generated build from /app to the nginx html folder -entrypoint of the webserver-
COPY --from=build /app/apps/web/dist /usr/share/nginx/html
# We copy the nginx conf file from our machine to our image
COPY /apps/web/nginx/nginx.conf /etc/nginx/conf.d/default.conf
# We expose the port 80 of the future containers
EXPOSE 80
# And finally we can run the nginx command to start the server
CMD ["nginx", "-g", "daemon off;"]
