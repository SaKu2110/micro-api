const { gql } = require('apollo-server');

module.exports = gql`
  type User {
    ID:  String
    NAME: String
    EMAIL: String
    PASSWORD: String
  }
  type Query {
    users: [User]!
    user(id: String!): User!
  }
`;