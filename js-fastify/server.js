const fastifyFactory = require('fastify');
const fastifyRedis = require('fastify-redis');
const fastifyPostgres = require('fastify-postgres');

const port = process.env.PORT || 3000;
const redisConnectionString = process.env.REDIS_URL;
const pgConnectionString = process.env.PG_URL;

const fastify = fastifyFactory({
  logger: true,
});

if (redisConnectionString) {
  fastify.register(fastifyRedis, {
    url: redisConnectionString,
  });
  fastify.log.info(`Redis enabled on: ${redisConnectionString}`);
}

if (pgConnectionString) {
  fastify.register(fastifyPostgres, {
    url: pgConnectionString,
  });
  fastify.log.info(`PG enabled on: ${pgConnectionString}`);
}

fastify.get('/', async (_req, _res) => {
  return { 'msg': 'ok' };
});

fastify.listen(port,'0.0.0.0', (err) => {
  if (err) {
    fastify.log.fatal('Error starting server %o', err);
    process.exit(1);
  } 
  fastify.log.info('Server started on port %s', port);
});





