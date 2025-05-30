#!/bin/bash

# 加载 .env 文件为变量（忽略注释和空行）
export $(grep -v '^#' .env | sed 's/ *#.*//' | xargs)



# 检查变量是否存在
if [[ -z "$MYSQL_USER" || -z "$MYSQL_PASSWORD" || -z "$MYSQL_HOST" || -z "$MYSQL_DB" ]]; then
  echo "❌ 环境变量不完整，请确认 .env 文件包含 MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_DB"
  exit 1
fi

# --help 参数：显示所有用法说明
if [[ "$1" == "--help" ]]; then
  echo "🧰 MySQL 工具脚本用法说明"
  echo ""
  echo "基础用法："
  echo "  bash mysql.sh                  # 列出所有表及每个表的行数与字段数"
  echo ""
  echo "参数："
  echo "  --help                         # 显示本帮助信息"
  echo "  --shell [表名]                # 显示表结构后进入交互式 mysql 命令行（默认 tm_user）"
  echo "  --list [表名]                 # 显示指定表最新 5 条记录（默认 tm_user）"
  echo "  表名（不带参数）              # 显示该表结构（使用 SHOW CREATE TABLE）"
  echo ""
  exit 0
fi

# --shell 参数: 先显示表结构再进入交互式命令行
if [[ "$1" == "--shell" ]]; then
  shift
  table_name=${1:-"tm_user"}  # 默认表名 tm_user
  echo "📋 先显示表结构: $table_name"
  mysql -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" -h"$MYSQL_HOST" -P${MYSQL_PORT:-3306} -D"$MYSQL_DB" -e "SHOW CREATE TABLE \`$table_name\`;" 2>/dev/null
  echo "🚀 进入 MySQL 交互命令行..."
  mysql -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" -h"$MYSQL_HOST" -P${MYSQL_PORT:-3306} -D"$MYSQL_DB"
  exit 0
fi

# --list 参数: 显示某张表的最新 5 条记录
if [[ "$1" == "--list" ]]; then
  shift
  table_name=${1:-"tm_user"}  # 默认表名 tm_user
  echo "📋 显示表 $table_name 的最新 5 条记录："
  mysql -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" -h"$MYSQL_HOST" -P${MYSQL_PORT:-3306} -D"$MYSQL_DB" \
    -e "SELECT * FROM \`$table_name\` ORDER BY id DESC LIMIT 5;" 2>/dev/null
  exit 0
fi


# 如果传入表名参数，只显示该表结构
if [[ -n "$1" ]]; then
  echo ""
  echo "表名: $1"
  echo "建表语句（逗号替换成竖线对齐显示）："
  mysql -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" -h"$MYSQL_HOST" -P${MYSQL_PORT:-3306} -D"$MYSQL_DB" -e "SHOW CREATE TABLE \`$1\`\G" 2>/dev/null
  exit 0
fi

echo "📡 正在连接 MySQL 数据库并列出所有表..."

# 列出所有表
tables=$(mysql -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" -h"$MYSQL_HOST" -P${MYSQL_PORT:-3306} -D"$MYSQL_DB" -N -e "SHOW TABLES;" 2>/dev/null)

echo ""
echo "📋 所有表信息如下："
for table in $tables; do
  # 获取记录数量
  count=$(mysql -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" -h"$MYSQL_HOST" -P${MYSQL_PORT:-3306} -D"$MYSQL_DB" -N -e "SELECT COUNT(*) FROM \`$table\`;" 2>/dev/null)

  # 获取字段数量
  column_count=$(mysql -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" -h"$MYSQL_HOST" -P${MYSQL_PORT:-3306} -D"$MYSQL_DB" -N -e "SHOW COLUMNS FROM \`$table\`;" 2>/dev/null | wc -l)

  printf "🔹 表: %-25s | 行数: %-8s | 字段数: %-5s\n" "$table" "$count" "$column_count"
done