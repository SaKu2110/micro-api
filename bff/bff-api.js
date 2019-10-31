const typeDefs = require('./schema/schema');
const resolvers = require('./resolver/resolver');
const { ApolloServer } = require('apollo-server');
const server = new ApolloServer({
  typeDefs,
  resolvers,
  dataSources: () => {
    return {
      SignAPI: new SignAPI(),
    };
  },
});

server.listen().then(({ url }) => {
  console.log(`Server ready at ${url}`);
});