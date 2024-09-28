FROM node:latest AS build
WORKDIR self-monorepo-project
COPY nx.json /self-monorepo-project/
COPY tsconfig.base.json /self-monorepo-project/
COPY .nx /self-monorepo-project/.nx
COPY package.json /self-monorepo-project/
COPY apps/frontend ./apps/frontend
RUN npm install
RUN npx nx run frontend:build --configuration=production 
# --verbose

FROM docker.io/nginx:stable-alpine 
COPY --from=build /self-monorepo-project/dist/apps/frontend/* /usr/share/nginx/html/
COPY --from=build /self-monorepo-project/apps/frontend/nginx.conf  /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
