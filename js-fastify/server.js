const path = require('path');
const fastifyFactory = require('fastify');
const fastifyMongo = require('fastify-mongodb');
const static = require('fastify-static');
const pov = require('point-of-view');
const ejs = require('ejs');

const port = process.env.PORT || 3000;
const mongoConnectionString = process.env.MONGO_URL;

const fastify = fastifyFactory({
  logger: true,
});

fastify.register(pov, { engine: { ejs } } );

fastify.register(static, {
  root: path.join(__dirname, 'public/assets'),
  prefix: '/assets/',
})

if (mongoConnectionString) {
  fastify.register(fastifyMongo, {
    url: mongoConnectionString,
    forceClose: true,
  });
  fastify.log.info(`Mongo enabled on: ${mongoConnectionString}`);
}



const increment = async (mongoDb) => {
  const collection = fastify.mongo.db.collection('data');
  
  const res = await collection.update({ _id: 'counter' }, {
    '$inc' : { 'value': 1 },
  }, { upsert: true });
  const counter = await collection.findOne({ _id: 'counter' });
  return counter;
}

fastify.get('/', async (_req, res) => {
  const counter = await increment();
  res.view('/public/index.html', { counter: counter.value });
})

fastify.get('/api/counter', async (_req, _res) => {
  const counter = await increment();
  return { 'msg': 'ok', counter: counter.value };
});

fastify.listen(port,'0.0.0.0', (err) => {
  if (err) {
    fastify.log.fatal('Error starting server %o', err);
    process.exit(1);
  } 
  fastify.log.info('Server started on port %s', port);
});





