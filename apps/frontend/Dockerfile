FROM node:21.7.3-alpine AS build
WORKDIR self-monorepo-project
COPY nx.json package.json package-lock.json tsconfig.base.json /self-monorepo-project/
COPY ./apps/frontend /self-monorepo-project/apps/frontend/
RUN npm install
RUN npx nx run frontend:build

FROM docker.io/nginx:stable-alpine 
COPY --from=build /self-monorepo-project/dist/frontend/* /usr/share/nginx/html/
COPY --from=build /self-monorepo-project/apps/frontend/nginx.conf  /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
