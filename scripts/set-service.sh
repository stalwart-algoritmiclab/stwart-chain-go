sudo tee /etc/systemd/system/node.service > /dev/null <<EOF
[Unit]
Description=stwart-1
After=network-online.target

[Service]
User=root
ExecStart=/bc/source/stwartd start --home=/bc
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
EOF

cat /etc/systemd/system/node.service

echo after this step, you can start the node with:
echo systemctl enable node
echo systemctl start node
echo
echo watch node logs:
echo journalctl -n 20 -f -u node --output=cat