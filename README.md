## Hyperledger Fabric 
## Employee Portal Nework (using fabric-ca as MSP)

1. Download the required fabric images by running the bootstarp.sh
    -   hyperledger/fabric-ca-tools
    -   hyperledger/fabric-ca-peer
    -   hyperledger/fabric-ca-orderer
    -   hyperledger/fabric-ca
    -   hyperledger/fabric-tools
    -   hyperledger/fabric-orderer
    -   hyperledger/fabric-peer 
    -   hyperledger/fabric-javaenv
    -   hyperledger/fabric-ccenv
    -   hyperledger/fabric-baseimage
    -   hyperledger/fabric-zookeeper
    -   hyperledger/fabric-kafka
    -   hyperledger/fabric-couchdb
    -   hyperledger/fabric-baseos

2.  Plan the network topology and modify the env.sh,makedocker.sh scripts appropriately. Once done start.sh     script can be used to bring up the network
    Network topology of *Employee Portal Nework*
    -   2 organisation (gov.in and private.in)
    -   1 orderer org
    -   1 peer in gov.in org
    -   1 peer in private.in
    -   1 channel
    -   Fabric CA as MSP in each organisation
    -   Chaincode written in **go**
    -   Fabric Node SDK to invoke chaincode

**Note : Issue may arise with the installing and initiating chaincode in peers. Chaincode has to be placed in the $GOPATH/src location**
  
