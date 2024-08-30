import { Resolvers } from '../__generated__/resolvers-types'
import * as query from './query'
import * as mutation from './mutation'

const resolvers: Resolvers = {
  Query: query,
  Mutation: mutation
}

export default resolvers
