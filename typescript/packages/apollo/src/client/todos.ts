import { gql } from '@apollo/client/core'
import { client } from './client'

client
  .query({
    query: gql`
      query GetTodos {
        todos {
          id
          text
          done
          user {
            id
            name
          }
        }
      }
    `
  })
  .then(result => console.log(result.data.todos))
  .catch(error => console.error(error))
