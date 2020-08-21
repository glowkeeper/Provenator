//const IterableData = artifacts.require("./IterableData.sol");
const Entities = artifacts.require("./Entities.sol");
const Artefacts = artifacts.require("./Artefacts.sol");

module.exports = async function(deployer) {

  await deployer.deploy(Entities)
  await deployer.deploy(Artefacts, Entities.address)

  deployer.then( () => {
    console.log(Entities.address);
    console.log(Artefacts.address);
  });
};
