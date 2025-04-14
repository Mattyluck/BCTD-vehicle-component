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

## Key Technologies

### Identifier Resolution Technology

The component "real-time code" identification mechanism utilizes Handle resolution technology to achieve precise positioning and record tracing throughout the component lifecycle. The implementation allows for efficient query of component information on the blockchain.

### DBSCAN Source Code Compression Algorithm

The C-language implementation of DBSCAN clustering reduces the size of the proof secret, which:
- Decreases the computational complexity of zero-knowledge proofs
- Maintains high coverage of critical code paths
- Achieves approximately 26% compression rate while preserving security properties

### Enhanced Groth16 ZK-SNARKs

Our improved Shamir secret sharing mechanism between blockchain nodes enhances the traditional Groth16 algorithm by:
- Preventing backdoor risks in the zero-knowledge proof system
- Maintaining constant proof size (approximately 256 bytes)
- Achieving efficient proof generation and verification times at millisecond levels

### Hyperledger Fabric Integration

The framework is built on Hyperledger Fabric with:
- Multi-organization architecture (including component manufacturers and vehicle OEMs)
- PBFT consensus for secure distributed operation
- Chaincode implementation of verification logic

## Performance Results

Our experimental evaluation demonstrates:
- Resolution efficiency at millisecond levels
- Proof generation time below 100ms for realistic code sizes
- Verification time under 50ms
- Constant proof size of 256 bytes regardless of code size
- Superior performance compared to ZK-STARKs and Bulletproofs for our use case


## Installation and Setup

### Requirements

- Go 1.19+
- Hyperledger Fabric 2.2+
- Node.js 12+
- Docker and Docker Compose
- GCC compiler

### Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/BCTD-ECU.git
cd BCTD-ECU
# Extract the complete codebase (optional)
tar -xzvf BC-ECU.tar.gz
# Install Go dependencies
go mod download
```

## Usage

### Setting up the Blockchain Network

This section helps you establish the Hyperledger Fabric network with multiple organizations and peers. The blockchain network serves as the foundation for the trusted detection system, providing immutable storage and consensus mechanisms.

```bash
# Initialize Hyperledger Fabric network
cd network
./start_network.sh
# Deploy chaincode
./deploy_chaincode.sh
```

### Running the Source Code Compression Algorithm

The source code compression module uses DBSCAN clustering to reduce the source code size while preserving security-critical components. This step is essential for efficient zero-knowledge proof generation without compromising verification effectiveness.

```bash
# Extract DBSCAN implementation
unzip DBSCAN.zip -d dbscan
cd dbscan

# Run compression on sample code
go run compression.go -input /path/to/source/code -output /path/to/output -ratio 0.26
```

### Generating and Verifying Zero-Knowledge Proofs

The enhanced ZK-SNARKs implementation enables privacy-preserving verification of source code integrity. This module implements the Groth16 protocol with our improvements to the trusted setup phase using Shamir secret sharing to eliminate backdoor risks.

```bash
# Extract ZK-SNARKs implementation
tar -xzvf zk_eq_blockchain.tar.gz
cd zk_eq_blockchain

# Setup the ZK system
go run setup.go
# Generate a proof
go run prove.go -input /path/to/compressed/source -output proof.json
# Verify a proof
go run verify.go -proof proof.json
```

### Component Identification and Resolution

The Handle-based identification and resolution system provides precise tracking of ECU components throughout their lifecycle. This module implements the "real-time code" mechanism that enables component authentication at any point in the supply chain.

```bash
# Extract identifier resolution implementation
unzip ParseProtocol.zip -d parse_protocol
cd parse_protocol
# Generate a new component identifier
go run generate.go -component /path/to/component/data
# Verify component authenticity
go run verify.go -id "component_id"
```

## Experimental Results

Based on our evaluations:

- The handle resolution system achieves response times of approximately 40ms
- Source code compression achieves a 26% compression ratio while maintaining security verification
- ZK-SNARKs proof size remains constant at 256 bytes regardless of source code size
- Proof generation and verification times are maintained at millisecond levels

## Technical Details

### DBSCAN Clustering

The DBSCAN algorithm is employed for source code compression, which uses code complexity metrics (V(G), M, H) to identify and extract representative code segments.

### Shamir Secret Sharing

An improved Shamir secret sharing scheme is implemented to enhance the trusted setup phase of the Groth16 ZK-SNARKs protocol, effectively preventing backdoor risks.

### Blockchain Architecture

The system uses Hyperledger Fabric with a two-organization, multi-peer architecture with RAFT consensus to ensure tamper-proof records of component lifecycle events.
