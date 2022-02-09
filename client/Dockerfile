FROM node:16.13.2

WORKDIR /app

COPY . .

RUN npm install yarn
RUN yarn install 

EXPOSE 3000

CMD ["yarn", "start:prod"]