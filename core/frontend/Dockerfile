FROM node:18-alpine as build
WORKDIR /app

COPY package*.json ./
RUN npm install --quiet

RUN npm install react-scripts@3.4.1 -g --quiet
COPY . ./

RUN npm run build

# production stage
FROM nginx:alpine
COPY --chown=nginx:nginx --from=build /app/build /usr/share/nginx/html
COPY --chown=nginx:nginx nginx/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 3000
CMD ["nginx", "-g", "daemon off;"]
