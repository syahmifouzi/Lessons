### Example to use .env
```
datasource db {
  provider = "sqlite"
  url      = env("DATABASE_URL")
}
```

### Docker Build cmd
`docker build -t syahmifouzi/graphql:1.3 .`

### Step to edit DB schema
1. Edit the prisma schema at `prisma/schema.prisma`
2. Migrate prisma by cmd `npx prisma migrate dev --name "added sensor data"`
3. Might need to restart docker container
4. Edit yoga schema at `src/schema.ts`
5. Might need to restart docker container