const schema = require('schema/schema');
const resolvers = require('resolver/resolver');
const { ApolloServer, gql } = require('apollo-server');
const server = new ApolloServer({ schema, resolvers });

server.listen().then(({ url }) => {
  console.log(`Server ready at ${url}`);
});