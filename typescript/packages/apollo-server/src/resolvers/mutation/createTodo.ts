import { MutationResolvers } from '../../__generated__/resolvers-types'

export const createTodo: MutationResolvers['createTodo'] = async (_parent, args, _context, _info) => {
  const { text, userId } = args.input
  console.log('created todo')
  return {
    id: '1',
    text: text,
    done: false,
    user: {
      id: userId,
      name: 'John Doe'
    }
  }
}
