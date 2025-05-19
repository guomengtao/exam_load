#!/bin/bash

# 1. 加载 .env 文件中的变量
ENV_FILE=".env"
if [ -f "$ENV_FILE" ]; then
  set -a
  while IFS='=' read -r key value; do
    if [[ "$key" =~ ^[A-Za-z_][A-Za-z0-9_]*$ ]]; then
      export "$key=$value"
    fi
  done < <(grep -v '^#' "$ENV_FILE" | grep '=')
  set +a
else
  echo "❌ 未找到 .env 文件"
  exit 1
fi

# 2. 映射 MySQL 配置变量（支持 MYSQL_ 前缀）
DB_HOST=${DB_HOST:-${MYSQL_HOST:-127.0.0.1}}
DB_PORT=${DB_PORT:-${MYSQL_PORT:-3306}}
DB_USER=${DB_USER:-$MYSQL_USER}
DB_PASS=${DB_PASS:-$MYSQL_PASSWORD}
DB_NAME=${DB_NAME:-$MYSQL_DB}

if [ -z "$DB_USER" ] || [ -z "$DB_PASS" ] || [ -z "$DB_NAME" ]; then
  echo "❌ .env 文件中缺少 DB_USER / DB_PASS / DB_NAME 配置"
  exit 1
fi

# 3. 输出目录
OUTPUT_DIR="docs/data"
mkdir -p "$OUTPUT_DIR"

# 4. 获取所有表名
TABLES=$(mysql -h$DB_HOST -P$DB_PORT -u$DB_USER -p$DB_PASS -D$DB_NAME -e "SHOW TABLES;" | tail -n +2)

echo "🔍 正在导出数据库结构和样例数据：$DB_NAME"

for TABLE in $TABLES; do
  echo "📄 处理表：$TABLE"

  DATA_FILE="${OUTPUT_DIR}/${TABLE}.data.sql"

  # 导出结构 + 数据（示例：前3条数据）
  {
    echo "-- ----------------------------"
    echo "-- Table structure for \`$TABLE\`"
    echo "-- ----------------------------"
    mysqldump -h$DB_HOST -P$DB_PORT -u$DB_USER -p$DB_PASS --no-data --skip-comments $DB_NAME $TABLE

    echo
    echo "-- ----------------------------"
    echo "-- Sample data for \`$TABLE\` (最多3条)"
    echo "-- ----------------------------"
    mysqldump -h$DB_HOST -P$DB_PORT -u$DB_USER -p$DB_PASS --skip-comments --no-create-info --order-by-primary --where="1 ORDER BY id ASC LIMIT 2" $DB_NAME $TABLE
  } > "$DATA_FILE"

  echo "✅ 结构+数据导出：$DATA_FILE"
done

echo "✅ 所有表处理完成，结果保存在：$OUTPUT_DIR"

#!/bin/bash

commit_msg="$1"
if [ -z "$commit_msg" ]; then
  commit_msg="代码和数据库结构同步更新"
fi

 # 显示当前git状态
git status

# 交互提示，读用户输入
read -p "是否提交改动？(y/n): " answer

if [[ "$answer" =~ ^[Yy]$ ]]; then
  echo "继续提交..."

  # 你后续的git add/commit/push操作
  git add .
  read -p "请输入提交信息: " commit_msg
  if [ -z "$commit_msg" ]; then
    commit_msg="代码和数据库结构同步更新"
  fi
  git commit -m "$commit_msg"
  git push origin main

  echo "提交并推送完成！"
else
  echo "已取消提交。"
  exit 0
fi