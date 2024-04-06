FROM eclipse-temurin:21

RUN apt-get update && apt-get install -y vim maven

WORKDIR /app

COPY funcs.bash /root/.bash_aliases

CMD bash
