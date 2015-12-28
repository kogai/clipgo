# clipgo
指定したテンプレートをクリップボードにコピーできるCLIツール

## 使い方

```bash
# インストール
brew tap kogai/clipgo
brew install clipgo

# 設定ファイルの生成
clipgo init

# 設定ファイルの確認
# `templatePath`にmarkdownファイルを突っ込むと準備完了
clipgo inspect

# テンプレートの内容をコピーする
clipgo copy <テンプレートとなるmarkdownファイル名>

# コピーできるテンプレート一覧
clipgo list

# コピーできるテンプレートを追加
clipgo add <追加したいmarkdownファイルへのパス>

# コピーできるテンプレートを削除
clipgo remove <削除したいテンプレートの名前>
```
