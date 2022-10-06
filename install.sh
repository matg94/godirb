mkdir -p ~/.godirb
curl -s https://raw.githubusercontent.com/matg94/godirb/main/default.yaml > ~/.godirb/default.yaml
curl -s https://raw.githubusercontent.com/matg94/godirb/main/common.txt > ~/.godirb/common.txt
wget https://github.com/matg94/godirb/releases/download/initial_installable_release/godirb

chmod +x ./godirb
mv ./godirb /usr/local/bin