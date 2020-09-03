//const IterableData = artifacts.require("./IterableData.sol");
const Entities = artifacts.require("./Entities.sol");
const Artefacts = artifacts.require("./Artefacts.sol");

module.exports = async function(deployer) {

  await deployer.deploy(Entities)
  await deployer.deploy(Artefacts)

  deployer.then( () => {
    console.log("static entitiesAddress = \"", Entities.address, "\"");
    console.log("static artefactsAddress = \"", Artefacts.address, "\"");
  });
};
