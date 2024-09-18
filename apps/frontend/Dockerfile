FROM docker.io/nginx:stable-alpine
COPY dist/apps/frontend/* /usr/share/nginx/html/
COPY /apps/frontend/nginx.conf  /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
