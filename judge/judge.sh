#!/bin/bash
# TODO: 入力のメタデータとして、どのテストデータを使うか指定する必要あり。
# TODO: stdinというディレクトリを作り、そこにテストデータを入れてすべて実行するようにする。

CURRENT_DIR_PATH="$(dirname "$(realpath "$0")" )"
cd $CURRENT_DIR_PATH

# メモリ制限と実行時間の制限を設定
MAX_MEMORY_MB=1000  # TODO: メタデータとして入力する必要あり
MAX_EXECUTION_TIME_SECONDS=2  # TODO: メタデータとして入力する必要あり

# 最大実行時間でプログラムを実行する
# python3 source.py自体の終了コードやエラー出力も取得する
MAX_MEMORY_KB=$((MAX_MEMORY_MB * 1000))
bash -c "
ulimit -v $MAX_MEMORY_KB # メモリ制限を設定
ulimit -t $MAX_EXECUTION_TIME_SECONDS # CPU時間の制限を設定
python3 source.py" \
< stdin.txt \
> result/stdout.txt \
2> result/stderr.txt

# 終了ステータスをチェック
EXIT_STATUS=$?
echo $EXIT_STATUS > result/exit_status.txt
if [ $EXIT_STATUS -eq 139 ]; then
    echo "Memory Limit Exceeded: $EXIT_STATUS"
elif [ $EXIT_STATUS -eq 137 ]; then
    echo "Time Limit Exceeded: $EXIT_STATUS"
elif [ $EXIT_STATUS -ne 0 ]; then
    echo "Error occurred during execution: $EXIT_STATUS"
else
    echo "Success: $EXIT_STATUS"
fi
