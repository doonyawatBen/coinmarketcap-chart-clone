FROM node:18-alpine as builder

ENV NODE_ENV=production

RUN apk update \
  && apk upgrade

WORKDIR /app

COPY package*.json ./
RUN npm install 

COPY . /app

RUN npm run build

FROM nginx:alpine

RUN apk update \
  && apk upgrade
  
RUN rm -rf  /usr/share/nginx/html/*
COPY --from=builder /app/dist /usr/share/nginx/html

RUN rm -rf /etc/nginx/conf.d/*
COPY --from=builder /app/nginx.conf /etc/nginx/conf.d/default.conf

CMD ["nginx", "-g", "daemon off;"]

EXPOSE 80