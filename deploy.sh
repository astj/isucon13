set -xue

# sudo ~/webapp/deploy.sh

cd /home/isucon/webapp/go
sudo -u isucon git pull origin main
systemctl stop isupipe-go.service
sudo -u isucon bash -c 'source ~/.profile && make build'
systemctl restart isupipe-go.service
# TODO: ここに pdns の再起動とか入れる？
systemctl reload nginx
