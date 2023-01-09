# Hyperledger Fabric

[![Build Status](https://dev.azure.com/Hyperledger/Fabric/_apis/build/status/Merge?branchName=main)](https://dev.azure.com/Hyperledger/Fabric/_build/latest?definitionId=51&branchName=main)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/955/badge)](https://bestpractices.coreinfrastructure.org/projects/955)
[![Go Report Card](https://goreportcard.com/badge/github.com/hyperledger/fabric)](https://goreportcard.com/report/github.com/hyperledger/fabric)
[![GoDoc](https://godoc.org/github.com/hyperledger/fabric?status.svg)](https://godoc.org/github.com/hyperledger/fabric)
[![Documentation Status](https://readthedocs.org/projects/hyperledger-fabric/badge/?version=latest)](http://hyperledger-fabric.readthedocs.io/en/latest)

This project is a _Graduated_ Hyperledger project. For more information on the history of this project see the [Fabric wiki page](https://wiki.hyperledger.org/display/fabric). Information on what _Graduated_ entails can be found in
the [Hyperledger Project Lifecycle document](https://tsc.hyperledger.org/project-lifecycle.html).
Hyperledger Fabric is a platform for distributed ledger solutions, underpinned
by a modular architecture delivering high degrees of confidentiality,
resiliency, flexibility and scalability. It is designed to support pluggable
implementations of different components, and accommodate the complexity and
intricacies that exist across the economic ecosystem.

Hyperledger Fabric delivers a uniquely elastic and extensible architecture,
distinguishing it from alternative blockchain solutions. Planning for the
future of enterprise blockchain requires building on top of a fully-vetted,
open source architecture; Hyperledger Fabric is your starting point.

## Changes 

```bash 
fabric/
├── bccsp
│   ├── dilithiumopts.go
│   ├── opts.go
│   ├── signer
│   │   └── signer.go
│   └── sw
│       ├── dilithium.go
│       ├── dilithiumkey.go
│       ├── fileks.go
│       ├── keygen.go
│       ├── keyimport.go
│       ├── keys.go
│       └── new.go
├── config
│   ├── configtx.yaml
│   ├── core.yaml
│   └── orderer.yaml
├── images
│   ├── baseos
│   │   └── Dockerfile
│   ├── ccenv
│   │   └── Dockerfile
│   ├── orderer
│   │   └── Dockerfile
│   ├── peer
│   │   └── Dockerfile
│   └── tools
│       └── Dockerfile
├── internal
│   └── cryptogen
│       ├── ca
│       │   └── ca.go
│       ├── csp
│       │   └── csp.go
│       └── msp
│           └── msp.go
├── msp
│   ├── identities.go
│   ├── mspimpl.go
│   └── mspimplsetup.go
└── test-network
    ├── CHAINCODE_AS_A_SERVICE_TUTORIAL.md
    ├── compose
    │   ├── compose-ca.yaml
    │   ├── compose-couch.yaml
    │   ├── compose-test-net.yaml
    │   ├── docker
    │   │   ├── docker-compose-ca.yaml
    │   │   ├── docker-compose-couch.yaml
    │   │   ├── docker-compose-test-net.yaml
    │   │   └── peercfg
    │   │       └── core.yaml
    │   └── podman
    │       ├── peercfg
    │       │   └── core.yaml
    │       ├── podman-compose-ca.yaml
    │       ├── podman-compose-couch.yaml
    │       └── podman-compose-test-net.yaml
    ├── configtx
    │   └── configtx.yaml
    ├── monitordocker.sh
    ├── network.sh
    ├── organizations
    │   ├── ccp-generate.sh
    │   ├── ccp-template.json
    │   ├── ccp-template.yaml
    │   ├── cryptogen
    │   │   ├── crypto-config-orderer.yaml
    │   │   ├── crypto-config-org1.yaml
    │   │   └── crypto-config-org2.yaml
    │   └── fabric-ca
    │       └── registerEnroll.sh
    ├── scripts
    │   ├── ccutils.sh
    │   ├── configUpdate.sh
    │   ├── createChannel.sh
    │   ├── deployCCAAS.sh
    │   ├── deployCC.sh
    │   ├── envVar.sh
    │   ├── org3-scripts
    │   │   ├── joinChannel.sh
    │   │   └── updateChannelConfig.sh
    │   ├── pkgcc.sh
    │   ├── setAnchorPeer.sh
    │   └── utils.sh
    └── setOrgEnv.sh
``` 
