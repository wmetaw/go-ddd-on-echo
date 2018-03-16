FROM node:7.5

ENV WORKSPACE /workspace

RUN apt-get update

RUN mkdir $WORKSPACE
WORKDIR $WORKSPACE
RUN mkdir docs output

RUN npm init -y
RUN npm install -g gulp
RUN npm install gulp gulp-aglio

COPY entrypoint.sh entrypoint.sh
COPY gulpfile.js gulpfile.js
COPY aglioconfig.json aglioconfig.json

# Start gulp and SimpleHTTPServer 8000 port
RUN chmod +x ./entrypoint.sh
CMD ["/bin/sh", "./entrypoint.sh"]
