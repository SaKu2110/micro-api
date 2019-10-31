const { RESTDataSource } = require('apollo-datasource-rest');

class SignAPI extends RESTDataSource {
  constructor() {
    super();
    this.baseURL = 'http://127.0.0.1:9000/';
  }

  async getUsers() {
    return this.get(`users`);
  }
}

module.exports = {
    Query: {
        users: () => async (_source, { dataSources }) => {
            return dataSources.SignAPI.getUsers();
        },
    },
};