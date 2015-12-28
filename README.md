指定したテンプレートをクリップボードにコピーできるCLIツール

## 使い方

### インストール

```bash
# インストール
brew tap kogai/golip
brew install golip

# 設定ファイルの生成
golip init

# 設定ファイルの確認
# `templatePath`にmarkdownファイルを突っ込むと準備完了
golip inspect

# コピーできるテンプレート一覧
golip list

# テンプレートの内容をコピーする
golip copy <テンプレートとなるmarkdownファイル名>

# コピーできるテンプレートを追加
golip add <追加したいmarkdownファイルへのパス>

# コピーできるテンプレートを削除
golip remove <削除したいテンプレートの名前>
```

### アップデート

```bash
brew untap golip
brew uninstall golip

brew tap golip
brew install golip
```
