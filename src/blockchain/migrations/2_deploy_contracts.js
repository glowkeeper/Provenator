var StringsLib = artifacts.require("./Strings.sol");
var PremisAgent = artifacts.require("./PremisAgent.sol");
var PremisEvent = artifacts.require("./PremisEvent.sol");
var PremisObject = artifacts.require("./PremisObject.sol");
var PremisRights = artifacts.require("./PremisRights.sol");

module.exports = function(deployer) {
  deployer.deploy(StringsLib);
  deployer.link(StringsLib, PremisAgent);
  deployer.link(StringsLib, PremisEvent);
  deployer.link(StringsLib, PremisObject);
  deployer.link(StringsLib, PremisRights);
  deployer.deploy(PremisAgent);
  deployer.deploy(PremisEvent);
  deployer.deploy(PremisObject);
  deployer.deploy(PremisRights);
};
