FROM node:14.18.1

WORKDIR /app

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

RUN apt-get update -y && \
  apt-get upgrade -y && \
  npm install npm@6.14.4 && \
  npm install ejs@3.0.1 && \
  npm install nuxt@2.9.1 &&\
  npm install create-nuxt-app@2.15.0 \
  npm install browserslist \
  npm install defu@3.1.0 \
  npm update



ADD . /app

EXPOSE 3000
CMD ["npm", "run", "dev"]