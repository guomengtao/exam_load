#!/bin/bash

BASE_URL="http://47.120.38.206:8081/api/badminton_game"

# Step 1: 插入两条记录
echo "📌 Step 1: 插入两条记录"
RESPONSE=$(curl -s -X POST "$BASE_URL" -H "Content-Type: application/json" -d '[
  {
    "player1": "Tom",
    "player2": "Jack",
    "score1": 21,
    "score2": 15,
    "date": "2025-05-29"
  },
  {
    "player1": "Alice",
    "player2": "Bob",
    "score1": 18,
    "score2": 21,
    "date": "2025-05-29"
  }
]')
echo "$RESPONSE"

echo -e "\n"
sleep 1

echo "📌 Step 2: 更新记录（仅修改 score）"
UPDATE_RESPONSE=$(curl -s -X PUT "$BASE_URL" -H "Content-Type: application/json" -d "[
  {\"id\": 1, \"score1\": 22, \"score2\": 20},
  {\"id\": 2, \"score1\": 19, \"score2\": 21}
]")
echo "$UPDATE_RESPONSE"
echo -e "\n"
sleep 1

echo "📌 Step 3: 查询验证更新是否成功，其他字段是否未被清空"
RESPONSE=$(curl -s -X GET "$BASE_URL/list?page=1&pageSize=50")
echo "$RESPONSE"

echo -e "\n📌 Step 4: 清理测试数据"
curl -s -X DELETE "$BASE_URL" -H "Content-Type: application/json" -d "[1,2]"
