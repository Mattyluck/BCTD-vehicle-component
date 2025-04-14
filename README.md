# BCTD-ECU: Blockchain-Based Trusted Detection for ECU Components

![image](https://github.com/Mattyluck/BCTD-vehicle-component/blob/master/BC-ECUfigure.png)

A novel approach for trusted detection of Electronic Control Unit (ECU) components in intelligent vehicles throughout their lifecycle using blockchain technology and zero-knowledge proofs.

## Overview

BCTD-ECU is a comprehensive solution that addresses the security challenges in the supply chain management of intelligent vehicle ECU components. As vehicles evolve into highly informatized and networked complex systems, their electronic control systems face increased security risks due to expanded attack surfaces and complex supply chains.

This project implements:

1. A blockchain-based component circulation analysis solution with a "real-time code" identification mechanism
2. A trusted source code detection scheme using non-interactive zero-knowledge proofs
3. An enhanced ZK-SNARKs method with an improved Shamir secret sharing mechanism to prevent backdoor risks

## Features

- **Component Lifecycle Tracking**: Precise positioning and record tracing throughout the ECU component lifecycle
- **Source Code Compression**: DBSCAN clustering algorithm for efficient source code sampling
- **Enhanced ZK-SNARKs**: Improved zero-knowledge proof system with Shamir secret sharing
- **Secure Supply Chain Management**: End-to-end solution for component verification and authentication
- **Privacy-Preserving Verification**: Verifies code authenticity without exposing proprietary source code

## Repository Structure

```
.
├── pkg/                    # Package directory for the Go modules
├── vendor/                 # Vendor directory for dependencies
├── BC-ECU.tar.gz           # Complete codebase archive
├── DBSCAN.zip              # Source code compression algorithm implementation
├── ParseProtocol.zip       # Identifier resolution technology implementation
├── filter_code.txt         # Code filtering rules and specifications
├── go.mod                  # Go module definition
├── go.sum                  # Go module checksums
├── main.go                 # Entry point and Shamir secret sharing implementation
├── nohup.out               # Experimental results (ZKP size, proof generation/verification time)
└── zk_eq_blockchain.tar.gz # Zero-knowledge proof implementation
```

