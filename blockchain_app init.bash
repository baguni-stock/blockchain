# 기본설정
cd ~/ 
APP="blockchain"
APP_DAEMON=$APP"d"
APP_HOME=$HOME/.$APP

# go install
wget https://golang.org/dl/go1.17.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -xzf go1.17.linux-amd64.tar.gz -C /usr/local
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
echo "export GOPATH=$GOPATH" > $HOME/.bash_profile
echo "export PATH=$PATH" >> $HOME/.bash_profile
mkdir $GOPATH
mkdir $GOPATH/bin
mkdir $GOPATH/src
mkdir $GOPATH/pkg

# starport install
curl https://get.starport.network/starport | bash
sudo mv -f starport /usr/local/bin/

# blockchain app install
wget https://github.com/chainstock-project/blockchain/releases/download/0.1.0/blockchaind-x86-64 -O $APP_DAEMON
chmod 550 $APP_DAEMON
mv $APP_DAEMON $GOPATH/bin

$APP_DAEMON keys add root --keyring-backend test
ROOT_ADDRESS=$($APP_DAEMON keys show root -a --keyring-backend test)
$APP_DAEMON init stock-chain --chain-id stock-chain
$APP_DAEMON add-genesis-account $ROOT_ADDRESS 100000000000stake
$APP_DAEMON gentx root 100000000stake --chain-id stock-chain --keyring-backend test
$APP_DAEMON collect-gentxs

# service regist
mkdir -p /var/log/$APP_DAEMON
touch /var/log/$APP_DAEMON/$APP_DAEMON.log
touch /var/log/$APP_DAEMON/$APP_DAEMON.error.log
touch /etc/systemd/system/$APP_DAEMON.service
service="[Unit]
Description=$APP_DAEMON daemon
After=network-online.target
[Service]
User=$USER
ExecStart=$HOME/go/bin/$APP_DAEMON start --home=$APP_HOME
WorkingDirectory=$HOME/go/bin
StandardOutput=file:/var/log/$APP_DAEMON/$APP_DAEMON.log
StandardError=file:/var/log/$APP_DAEMON/$APP_DAEMON.error.log
Restart=always
RestartSec=3
LimitNOFILE=4096
[Install]
WantedBy=multi-user.target"
echo "$service" | sudo tee -a /etc/systemd/system/$APP_DAEMON.service
systemctl enable $APP_DAEMON.service
systemctl start $APP_DAEMON.service

# check node
$HOME/go/bin/$APP_DAEMON --home=$APP_HOME tendermint show-node-id
journalctl -u $APP_DAEMON -f