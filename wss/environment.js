
require("dotenv").config

module.exports = {
  REDIS_HOST: process.env.REDIS_HOST || "localhost",
  REDIS_PORT: process.env.REDIS_PORT || 6379,
  TOKEN_SECRET: process.env.TOKEN_SECRET || "mySuperSecretKey",
}