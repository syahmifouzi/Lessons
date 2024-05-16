### Example to use .env
```
datasource db {
  provider = "sqlite"
  url      = env("DATABASE_URL")
}
```

### Docker Build on RPi
Need to build on RPi for its architecture
When building on RPi, use the following RUN cmd
`RUN ["chmod", "+x", "/runner.sh"]`


### Docker Build cmd
`docker build -t syahmifouzi/graphql:1.3 .`
`docker save -o ./gql.tar syahmifouzi/graphql:pi-1.4`
`docker load -i <path to image tar file>`
- see docker commit for saving changes on existing image
- `docker commit -p container-ID backup-name`


### Step to edit DB schema
1. Edit the prisma schema at `prisma/schema.prisma`
2. Migrate prisma by cmd `npx prisma migrate dev --name "added sensor data"`
3. Might need to restart docker container
4. Edit yoga schema at `src/schema.ts`
5. Might need to restart docker container