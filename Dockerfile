FROM debian:9-slim

ADD ./bin/movieservice /app/bin/
WORKDIR /app
COPY . .

EXPOSE 8000

CMD [ "/app/bin/movieservice" ]

FROM node:12

ADD api-tests/Movieservice.postman_collection.json /app/bin/
WORKDIR /app

RUN npm install -g newman

CMD newman run /app/bin/Movieservice.postman_collection.json --global-var "localhost=http://localhost:8000"