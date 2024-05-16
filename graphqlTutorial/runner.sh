#!/bin/bash

# Start the first process
npm run dev &

# Start the second process
npx prisma studio &

# Wait for any process to exit
wait -n

# Exit with status of process that exited first
exit $?