#유저생성
blockchaind keys add <username> --keyring-backend test --home=~/.blockchaind

#유저정보조회
blockchaind keys show <username> --keyring-backend test

#유저목록조회(전체가 아닌 비밀키를 보유하고 있는 지갑들만)
blockchaind keys list --keyring-backend test

#유저 address 확인
blockchaind keys show <username> -a --keyring-backend test

#validator address 확인
blockchaind keys show <validator> --bech val -a --keyring-backend test

#기타 기능확인하기
blockchaind

#root 유저에게 token받기
blockchaind tx bank send <root_address> <username> 1000000stake --chain-id stock-chain --keyring-backend test

#자신의 token을 validator에게 위임
blockchaind tx staking delegate <validator_address> <token_count>stake --from <username> --chain-id stock-chain --keyring-backend test 

#자신이 validator로 참여

