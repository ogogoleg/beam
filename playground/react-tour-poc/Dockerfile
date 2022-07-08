# # pull official base image
# FROM node:18-alpine

# # set working directory
# WORKDIR /app

# # add `/app/node_modules/.bin` to $PATH
# ENV PATH /app/node_modules/.bin:$PATH

# # install app dependencies
# COPY package.json ./
# COPY package-lock.json ./
# RUN npm install --silent

# # add app
# COPY . ./

# EXPOSE 3000

# # start app
# CMD ["npm", "start"]



# build
FROM node:18-alpine as build_stage
WORKDIR /src
COPY package.json .
COPY package-lock.json .
RUN npm install --silent
COPY . .
RUN npm run build

# deploy
FROM nginx:stable-alpine as prod_stage
COPY --from=build_stage /src/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]