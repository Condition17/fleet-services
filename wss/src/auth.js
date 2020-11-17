const fs = require('fs');
const path = require('path');
const jwt = require("jsonwebtoken")
const {TOKEN_SECRET} = require("../environment");

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