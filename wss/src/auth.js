const jwt = require("jsonwebtoken")
const {TOKEN_SECRET} = require("../environment");

function getDecodedTokenValue(token) {
  return jwt.verify(token, TOKEN_SECRET);
}

function isValidToken(token) {
  try {
    return !!getDecodedTokenValue(token);
  } catch (err) {
    return false;
  }
}

module.exports = {
  isValidToken,
  getDecodedTokenValue
}