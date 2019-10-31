const typeDefs = require('./schema/schema');
const resolvers = require('./resolver/resolver');

const { RESTDataSource } = require('apollo-datasource-rest');
const { ApolloServer } = require('apollo-server');

class SignAPI extends RESTDataSource {
  constructor() {
    super();
    this.baseURL = 'http://127.0.0.1:9000/';
  }

  async getUsers() {
    const response = await this.get(`users`);
    return response.users;
  }
 
  async getUser(id) {
    const response = await this.get(`user/${id}`);
    return response.user;
  }

}

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