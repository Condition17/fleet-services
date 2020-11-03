const fs = require('fs');
const path = require('path');
const jwt = require("jsonwebtoken")
const secret = fs.readFileSync(path.join(__dirname, "..", "secret"))

function getDecodedTokenValue(token) {
  return jwt.verify(token, secret);
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