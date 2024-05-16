import { createSchema } from 'graphql-yoga'
import { type Link, type SensorData } from '@prisma/client'
import type { GraphQLContext } from './context'

const typeDefinitions = /* GraphQL */ `
    type Query {
        info: String!
        feed: [Link!]!
        sensorData(take: Int): [SensorData!]!
    }

    type Mutation {
        postLink(url: String!, description: String!): Link!
        postSensorData(value: Float!): SensorData!
    }

    type SensorData {
        id: ID!
        value: Float!
        createdAt: String!
    }

    type Link {
        id: ID!
        description: String!
        url: String!
    }
`

const resolvers = {
    Query: {
        info: () => `This is the API of a Hackernews Clone`,
        feed: async (parent: unknown, args: {}, context: GraphQLContext) => context.prisma.link.findMany(),
        async sensorData (parent: unknown, args: {take?: number}, context: GraphQLContext) {
            const take = args.take
            return context.prisma.sensorData.findMany({
                take
            })
        }
    },
    SensorData: {
        id: (parent: SensorData) => parent.id,
        value: (parent: SensorData) => parent.value,
    },
    Link: {
        id: (parent: Link) => parent.id,
        description: (parent: Link) => parent.description,
        url: (parent: Link) => parent.url,
    },
    Mutation: {
        async postLink(parent: unknown, args: { description: string; url: string }, context: GraphQLContext) {
            const newLink = await context.prisma.link.create({
                data: {
                    url: args.url,
                    description: args.description
                }
            })
            return newLink
        },
        async postSensorData(parent: unknown, args: { value: number }, context: GraphQLContext) {
            const newData = await context.prisma.sensorData.create({
                data: {
                    value: args.value
                }
            })
            return newData
        },
    }
}

export const schema = createSchema({
    resolvers: [resolvers],
    typeDefs: [typeDefinitions]
})