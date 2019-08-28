# 外部テスト

1. testディレクトリ以下に、*_test.goを作成する
1. "testing"パッケージをimportする
1. テストケースを記述する　この際、initializeパッケージのInitServer()を最初に実行する
1. `sudo docker-compose up` でdockerコンテナを起動
1. `sudo docker exec "foodmates-server_app_1" go test -v ./test/` でテスト実行
