const jsonServer = require('json-server');
const server = jsonServer.create();
const router = jsonServer.router('db.json');

// This line is changed to serve images from the "public" folder
const middlewares = jsonServer.defaults({ static: 'public' });

const port = process.env.PORT || 3001;

server.use(middlewares);
server.use(router);

server.listen(port, () => {
  console.log(`JSON Server is running on port ${port}`);
});