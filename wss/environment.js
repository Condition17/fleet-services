
function isProdEnvironment() {
  return process.env.NODE_ENV === "prod"
}

module.exports = {
  REDIS_HOST: isProdEnvironment() ? "redis" : "localhost",
  REDIS_PORT: 6379
}