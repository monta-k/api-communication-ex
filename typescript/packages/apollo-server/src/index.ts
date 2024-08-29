import { GraphQLFileLoader } from '@graphql-tools/graphql-file-loader'
import { loadSchemaSync } from '@graphql-tools/load'
import { addResolversToSchema } from '@graphql-tools/schema'
import { ApolloServer } from '@apollo/server'
import { startStandaloneServer } from '@apollo/server/standalone'
import { join } from 'path'
import resolvers from './resolvers'

const schema = loadSchemaSync(join(__dirname, '../../../../graphql/schema.graphqls'), {
  loaders: [new GraphQLFileLoader()]
})

const schemaWithResolvers = addResolversToSchema({ schema, resolvers })
const server = new ApolloServer({ schema: schemaWithResolvers })

startStandaloneServer(server, {
  listen: { port: 4000 }
}).then(({ url }) => console.log(`🚀  Server ready at: ${url}`))
