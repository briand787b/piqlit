# build stage
FROM node:13.8-alpine3.11 as builder
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# production stage
FROM nginx:stable-alpine
EXPOSE 80
COPY --from=builder /app/dist /usr/share/nginx/html/
COPY --from=builder /app/docker-entrypoint.sh /usr/local/bin/
CMD [ "/usr/local/bin/docker-entrypoint.sh" ]