FROM node:14-alpine as build-step
RUN mkdir /app
WORKDIR /app
COPY package.json /app
RUN npm install
COPY . /app
RUN npm run build

FROM nginx:1.17.1-alpine
COPY --from=build-step /app/build /usr/share/nginx/html
ADD entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
ENV API_HOST=localhost:42069
ENV API_PROTO=http
CMD ["/entrypoint.sh"]