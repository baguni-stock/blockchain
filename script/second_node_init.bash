git clone https://github.com/chainstock-project/blockchain
cd blockchain
make
blockchaind init localnode --home=~/.blockchain




wget https://github.com/chainstock-project/blockchain/releases/download/0.1.0/blockchaind
chmod 550 blockchaind
blockchaind init localnode --home=~/.blockchain
wget https://github.com/chainstock-project/blockchain/releases/download/0.1.0/genesis ~/.blockchain/config/


vim ~/.blockchain/config/config.toml
persistent_peers = "" --> persistent_peers = "1b17ad432cbfd2827b9631049480734e66b8b2dd@18.223.158.36:26657"
blockchaind start
