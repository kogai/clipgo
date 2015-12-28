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
```

### 使う

```bash
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
brew untap kogai/golip
brew uninstall golip

brew tap kogai/golip
brew install golip
```

## ヘルプ

```bash
$ golip help

usage: golip [<flags>] <command> [<args> ...]

Flags:
  --help     Show context-sensitive help (also try --help-long and --help-man).
  --version  Show application version.

Commands:
  help [<command>...]
    Show help.

  inspect
    現在の設定を表示します

  init
    クリップボードにコピーするテンプレートの初期設定をします

  list
    利用可能なテンプレートの一覧が表示されます

  copy <templatename>
    指定したタグのテンプレートをクリップボードにコピーします

  add <pathtofile>
    指定したファイルをテンプレートとして追加します

  remove <templatename>
    指定したテンプレートを削除します
```
