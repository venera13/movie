FROM debian:9-slim

ADD ./bin/cinema /app/bin/
WORKDIR /app

EXPOSE 8000

CMD [ "/app/bin/cinema" ]

FROM node:12

ADD ./api-tests/Cinema.postman_collection.json /app/bin
WORKDIR /app

RUN npm install -g newman

CMD newman run /app/bin/Cinema.postman_collection.json --global-var "localhost=http://localhost:8000"