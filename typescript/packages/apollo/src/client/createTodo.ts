import { gql } from '@apollo/client/core'
import { client } from './client'

client
  .mutate({
    mutation: gql`
      mutation CreateTodo {
        createTodo(input: { text: "Buy coffee", userId: "user-1" }) {
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
  .then(result => console.log(result.data.createTodo))
  .catch(error => console.error(error))
