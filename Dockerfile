FROM node:19.2.0
COPY ./netflix-remote-server /netflix-remote-server
WORKDIR /netflix-remote-server
RUN apt-get upgrade && apt-get update && apt-get -y install libxtst-dev libpng-dev
RUN npm install
ENTRYPOINT [ "npm", "start" ]
