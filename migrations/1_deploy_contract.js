const myCryptoContract = artifacts.require("MyCrypto");

module.exports = function(deployer) {
   deployer.deploy(myCryptoContract);
};