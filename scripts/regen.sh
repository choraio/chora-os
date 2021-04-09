# cleanup
rm -rf .regen

# init chain
regen init regen-node --chain-id regen --home .regen

# create key
regen keys add regen-validator --home .regen

# add genesis account
regen add-genesis-account $(regen keys show regen-validator -a --home .regen) 5000000stake --home .regen

# generate genesis transaction
regen gentx regen-validator 1000000stake --chain-id regen --home .regen

# collect genesis transactions
regen collect-gentxs --home .regen

# check and start client
if curl -sL --fail http://localhost:26657 -o /dev/null
then
	regen start --grpc.address 0.0.0.0:9191 --p2p.laddr tcp://127.0.0.1:26658 --rpc.laddr tcp://127.0.0.1:26659 --rpc.pprof_laddr 127.0.0.1:6161 --home .regen
else
	regen start --home .regen
fi
