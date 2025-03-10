FROM node:20

WORKDIR /app

RUN npm install -g @pact-foundation/pact-cli

COPY package.json package-lock.json ./
RUN npm install

COPY . .

EXPOSE 3000

CMD ["npm", "run", "dev"]
