 
	•	Redis的目录结构建议
	•	试卷的数据结构设计（带字段解释）
	•	旧数据库（MySQL）表结构整理（清晰对应）
	•	说明：中英文名词对照

⸻

Redis目录与试卷数据结构设计文档

📂 Redis目录结构建议

root
│
├── paper        # 试卷基本信息和题目
│     ├── {paper_id}   # 单份试卷，包含试卷基本信息和题目列表
│
├── paper_answer  # 学生作答后的答案数据
│     ├── {paper_id}:{student_id}  # 某份试卷某个学生的答题记录
│
├── other_data   # 其他相关数据（例如统计、批改信息等）

✍️说明：
	•	每份试卷（paper/{paper_id}）下，存完整的试卷基本信息和题目。
	•	学生答题结果单独存放，避免污染原始试卷。
	•	“目录”在Redis中实际就是用Key前缀模拟，比如 paper:123、paper_answer:123:456。

⸻

📝 Redis中单份试卷的数据结构（示例）

{
  "id": 123,
  "title": "期末数学测试卷",
  "description": "2025年初中二年级数学期末测试",
  "publish_time": 1714419200,
  "questions": [
    {
      "id": 1,
      "type": 1,               // 题目类型（选择题、判断题等）
      "title": "下列哪一个是质数？",
      "content": "请选择正确答案",
      "options": [
        "A. 4",
        "B. 5",
        "C. 6",
        "D. 8"
      ],
      "correct_answer": "B",
      "score": 5
    },
    {
      "id": 2,
      "type": 2,
      "title": "判断题：地球是圆的。",
      "content": "",
      "options": [],
      "correct_answer": "正确",
      "score": 2
    }
    // 更多题目
  ],
  "total_score": 100   // 所有题目的分数加总
}

✍️说明：
	•	questions 是数组，存每道题，保留每题分数。
	•	total_score 是所有题的 score 加和，可在导入时计算好。

⸻

📚 旧数据库（MySQL）结构整理

1. 表：ym_article （试卷基本信息）

字段	说明	备注
id	试卷ID	主键
title	试卷标题	必须字段
username	所属用户	可选
desc	试卷描述	可选
pic	缩略图	可选
content	试卷正文内容	可选（通常不使用）
click	点击量	可忽略
cateid	分类ID	可忽略
time	发布时间（时间戳）	需要保存
close	是否关闭	可忽略



⸻

2. 表：ym_article_contens （试卷题目）

字段	说明	备注
id	题目ID	主键
cate	所属试卷ID (ym_article.id)	外键关系
title	题目标题	必须字段
type	题目类型	0:未知，1:单选，2:判断，3:多选，…（根据实际定义）
score	该题分数	必须字段
correct	正确答案	格式统一（如A/B/C/正确/错误）
number	排序用编号	可选
contens	题目内容正文	额外说明（例如题干长内容）
answer1-6	选项内容	A-F，分别存不同选项
close	是否关闭	可忽略



⸻

📖 英文名词对照表

中文	英文（建议）
试卷	Paper
试卷题目	Question
学生作答	Answer
总分	Total Score
题目分数	Score
题目选项	Options
正确答案	Correct Answer
发布时间	Publish Time
描述	Description



⸻

⚙️下一步计划
	•	从MySQL中读取 ym_article 和对应的 ym_article_contens
	•	转成Redis指定的数据结构，保存到对应的目录下
	•	在导入时自动计算 total_score
	•	对题目的选项(answer1-6)进行数组化处理
	•	正确处理时间戳、空字段、默认值

⸻

