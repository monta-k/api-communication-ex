import { QueryResolvers } from '../../__generated__/resolvers-types'

export const todos: QueryResolvers['todos'] = async (_parent, _args, _context, _info) => {
  return [
    {
      id: '1',
      text: 'Buy milk',
      done: false,
      user: {
        id: '1',
        name: 'John Doe'
      }
    },
    {
      id: '2',
      text: 'Buy eggs',
      done: true,
      user: {
        id: '1',
        name: 'John Doe'
      }
    }
  ]
}
