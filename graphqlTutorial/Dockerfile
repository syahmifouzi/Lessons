FROM node:lts

COPY . .

RUN npm install

RUN npx prisma migrate dev --name "init"

EXPOSE 4000

EXPOSE 5555

CMD ./runner.sh