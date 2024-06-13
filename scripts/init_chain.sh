NODE_HOME=/bc
CHAIN_ID="stwart_test_1"
BINARY_NAME="./stwartd"
DENOM_FEE="\"stwart\""

API_ADDRESS="\"tcp://0.0.0.0:1317\""
GRPC_ADDRESS="\"0.0.0.0:9090\""
GRPC_WEB_ADDRESS="\"0.0.0.0:9091\""
CFG_RPC_LADDDR="\"tcp://0.0.0.0:26657\""
DEFAULT_MONIKER="alice"
# If an argument is passed, add it to MONIKER with a hyphen
# init_chain.sh testnet-1 as an example
 if [ $# -gt 0 ]; then
     MONIKER="${DEFAULT_MONIKER}-$1"
 else
     MONIKER="${DEFAULT_MONIKER}"
 fi
MINIMUM_GAS_PRICE="\"0.5stwart\""


${BINARY_NAME} init ${MONIKER} --chain-id $CHAIN_ID --overwrite --home ${NODE_HOME}
${BINARY_NAME} keys add "${MONIKER}" --keyring-backend=test --output json --home ${NODE_HOME} > ${NODE_HOME}/system_account_${MONIKER}.json

ADDRESS=$(${BINARY_NAME} keys show "${MONIKER}" --keyring-backend=test -a --home ${NODE_HOME})

${BINARY_NAME} genesis add-genesis-account "${ADDRESS}" 100000000000000stake --home ${NODE_HOME} --keyring-backend=test
${BINARY_NAME} genesis gentx $MONIKER 1000000000000stake --chain-id $CHAIN_ID --home ${NODE_HOME} --keyring-backend=test
${BINARY_NAME} genesis collect-gentxs $MONIKER --gentx-dir ${NODE_HOME}/config/gentx/ --home ${NODE_HOME}
${BINARY_NAME} genesis validate --home ${NODE_HOME}

 # Update app.toml
sed -i -e 's/^enable = false/enable = true/' $NODE_HOME/config/app.toml
sed -i -e "s,^denom-to-suggest = \"uatom\",denom-to-suggest = ${DENOM_FEE}," $NODE_HOME/config/app.toml
sed -i -e 's/^swagger = false/swagger = true/' $NODE_HOME/config/app.toml
sed -i -e "s,^address = \"tcp://localhost:1317\",address = ${API_ADDRESS}," $NODE_HOME/config/app.toml
sed -i -e "s,^address = \"localhost:9090\",address = ${GRPC_ADDRESS}," $NODE_HOME/config/app.toml
sed -i -e "s,^address = \"localhost:9091\",address = ${GRPC_WEB_ADDRESS}," $NODE_HOME/config/app.toml
sed -i -e "s/^minimum-gas-prices = \"\"/minimum-gas-prices = ${MINIMUM_GAS_PRICE}/" $NODE_HOME/config/app.toml

# Update config.toml
sed -i -e 's/^cors_allowed_origins = .*/cors_allowed_origins = ["*"]/' $NODE_HOME/config/config.toml
sed -i -e "s,^laddr = \"tcp://127.0.0.1:26657\",laddr = ${CFG_RPC_LADDDR}," $NODE_HOME/config/config.toml
sed -i -e "s/^max_subscription_clients = .*/max_subscription_clients = 200/" $NODE_HOME/config/config.toml
sed -i -e "s/^max_subscriptions_per_client = .*/max_subscriptions_per_client = 200/" $NODE_HOME/config/config.toml
sed -i -e 's/^timeout_commit = "5s"/timeout_commit = "8s"/' $NODE_HOME/config/config.toml
