FROM node:latest
WORKDIR /usr/src/app
COPY ./ ./
RUN npm install
EXPOSE 4000
CMD [ "node", "bff-api.js" ]
