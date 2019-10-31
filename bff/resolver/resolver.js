module.exports = {
    Query: {
      user: async (_source, { id },{ dataSources }) => {
        return dataSources.SignAPI.getUser(id);
      },
      users: async (_source, _args,{ dataSources }) => {
        return dataSources.SignAPI.getUsers();
      },
    },
};