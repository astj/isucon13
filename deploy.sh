set -xue

cd /home/isucon/webapp
sudo -u isucon git pull origin main
systemctl stop isupipe-go.service
sudo -u isucon make build
systemctl restart isupipe-go.service
# TODO: ここに pdns の再起動とか入れる？
systemctl reload nginx
