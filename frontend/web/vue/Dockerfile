# build stage
FROM node:lts-alpine AS builder
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