# clipgo
指定したテンプレートをクリップボードにコピーできるCLIツール

# 使い方

```bash
# インストール
brew tap kogai/clipgo
brew install clipgo

# 設定ファイルの生成
clipgo init

# 設定ファイルの確認
# `templatePath`にmarkdownファイルを突っ込むと準備完了
clipgo inspect

# コピーする
clipgo copy -t <テンプレートとなるmarkdownファイル名>

# コピーできるファイル一覧
clipgo list
```
