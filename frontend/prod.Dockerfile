FROM node AS builder

WORKDIR /var/src

COPY package.json .
COPY package-lock.json .
RUN npm install

COPY . .
RUN npm run build

FROM nginx

COPY ./config/nginx.conf /etc/nginx/conf.d

COPY --from=builder /var/src/public /var/html
