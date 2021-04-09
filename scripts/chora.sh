# cleanup
rm -rf .chora

# init chain
chorad init chora-node --chain-id chora --home .chora

# create key
chorad keys add chora-validator --home .chora

# add genesis account
chorad add-genesis-account $(chorad keys show chora-validator -a --home .chora) 5000000stake --home .chora

# generate genesis transaction
chorad gentx chora-validator 1000000stake --chain-id chora --home .chora

# collect genesis transactions
chorad collect-gentxs --home .chora

# check and start client
if curl -sL --fail http://localhost:26657 -o /dev/null
then
	chorad start --grpc.address 0.0.0.0:9191 --p2p.laddr tcp://127.0.0.1:26658 --rpc.laddr tcp://127.0.0.1:26659 --rpc.pprof_laddr 127.0.0.1:6161 --home .chora
else
	chorad start --home .chora
fi
