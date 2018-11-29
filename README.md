# Blockchain / Cryptocurrency implementation with Go

Minimum viable blockchain/cryptocurrency implemented with Go and React.

### Features
 - Proof-of-work algorithm mining
 - Transactions
 - Wallets with RSA private keys
 - Peer-to-peer connected nodes
   - Blockchain relaying
   - Transaction relaying
 - React UI
   - Wallet
   - Block explorer
   - Control panel


### How to run

**Requirements**
 - **Go**
 - **npm**

Then in order run these:

```bash
git clone git@github.com:Desentso/Blockchain.git
cd Blockchain

# Build the go application
go build

# Build the react bundles
cd ui
npm run build

cd ..
# Start the app
Blockchain.exe
```

Now everything should be working and running, you can access the app at http://localhost:9090
If you want to run the app on different port you can start it like this `Blockchain.exe PORT_NUMBER` e.g. to run it on port 9091 `Blockchain.exe 9091`
