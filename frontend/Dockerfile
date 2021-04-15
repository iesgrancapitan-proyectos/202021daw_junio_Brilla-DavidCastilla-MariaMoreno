FROM node:lts

WORKDIR /var/src

COPY package.json .
COPY package-lock.json .

RUN npm install

CMD ["npm", "run", "dev"]
