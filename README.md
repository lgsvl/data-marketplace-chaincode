# Data Marketplace Chaincode
This repository contains a Hyperledger Fabric chaincode that represents a smart contract to support the data marketplace operations and concepts. The project is written in [Go](https://golang.org/).
To run this component correctly, you should be familiar with the [Data marketplace](https://github.com/lgsvl/data-marketplace) components because there is a particular dependency between the components.
You should also have a running Fabric network, we followed this [tutorial](https://github.com/IBM/blockchain-network-on-kubernetes#4-deploy-hyperledger-fabric-network-into-kubernetes-cluster) to deploy Fabric on kubernetes but we used our chaincode instead of the one provided in the tutorial.

# Build prerequisites
  * Install [golang](https://golang.org/).
  * Install [git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git).
  * Configure go. GOPATH environment variable must be set correctly before starting the build process.

### Download and work with the code

```bash
mkdir -p $HOME/workspace
export GOPATH=$HOME/workspace
mkdir -p $GOPATH/src/github.com/lgsvl
cd $GOPATH/src/github.com/lgsvl
git clone git@github.com:lgsvl/data-marketplace-chaincode.git
cd data-marketplace-chaincode
```

### Kubernetes Deployment 
After deploying the chaincode, you can connect to any organization and start invoking the chaincode. To make things easier, we implemented a REST interface that you can deploy within your kubernetes cluster.
The code for this REST interface is in [Data marketplace Chaincode REST](https://github.com/lgsvl/data-marketplace-chaincode-rest).


# Running the Unit Tests

Run the tests:
```bash
./scripts/run_glide_up
./scripts/run_units.sh
```