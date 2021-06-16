FROM debian:9-slim

ADD ./bin/movieservice /app/bin/
WORKDIR /app
COPY . .

EXPOSE 8000

CMD [ "/app/bin/movieservice" ]