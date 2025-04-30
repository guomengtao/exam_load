生成的试卷表（exam_papers）设计

以下是关于生成的试卷表的 markdown 格式设计，包括每个字段的详细说明及 questions 字段的 JSON 格式存储结构。此表结构完全保留了你所要求的 JSON 格式存储题目列表。

# 生成的试卷表 (`exam_papers`)

## 表结构说明：

| 字段名            | 类型            | 是否为空 | 默认值 | 说明                             |
|------------------|-----------------|---------|--------|----------------------------------|
| id               | BIGINT(20)      | 否      | 无     | 主键ID，自增                      |
| title            | VARCHAR(255)    | 否      | 无     | 试卷标题                          |
| description      | TEXT            | 是      | NULL   | 试卷描述                          |
| cover_image      | VARCHAR(255)    | 是      | NULL   | 封面图片URL                       |
| total_score      | INT(11)         | 否      | 0      | 试卷总分                          |
| questions        | JSON            | 否      | 无     | 题目列表（JSON数组），详见下方示例 |
| view_count       | INT(11)         | 是      | 0      | 浏览量                            |
| category_id      | BIGINT(20)      | 是      | NULL   | 分类ID                            |
| publish_time     | INT(10)         | 是      | 0      | 最近一次发布时间                    |
| status           | TINYINT(1)      | 是      | 1      | 状态：1=正常，0=关闭               |
| creator          | VARCHAR(150)    | 是      | NULL   | 创建人用户名                      |
| created_at       | TIMESTAMP       | 否      | CURRENT_TIMESTAMP | 创建时间                    |
| updated_at       | TIMESTAMP       | 否      | CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP | 更新时间 |
| deleted_at       | TIMESTAMP       | 是      | NULL   | 删除时间（NULL=未删除）            |

### `questions` 字段内部 JSON 格式（标准版）

```json
[
  {
    "id": 1,
    "type": "single",             // 题目类型：single=单选, multi=多选, judge=判断, fill=填空, img=图片题
    "title": "下面哪个是红色的？",
    "options": ["红色", "蓝色", "绿色", "黄色"],
    "correct_answer": "红色",
    "score": 5,
    "image_url": null             // 如果是图片题，这里放图片文件名或完整URL
  },
  {
    "id": 2,
    "type": "img",
    "title": null,
    "options": [],
    "correct_answer": "A",         // 比如图片题配合选项，正确选A
    "score": 10,
    "image_url": "image_12345.png"
  }
]

备注：
	•	questions 字段：此字段是 JSON 格式，包含了试卷的题目列表。每一题可以有不同的类型（单选、多选、判断、填空、图片题等），具体内容和正确答案也包含在其中。
	•	total_score 字段：记录试卷的总分数。
	•	created_at 和 updated_at 字段：自动记录生成时间及更新时间。
	•	status 字段：表示试卷的状态（例如：正常、关闭）。

### 解释：
1. **`questions` 字段的 `JSON` 格式**：根据你给出的要求，我们保留了 `JSON` 格式用于存储题目列表。每道题目都有 `id`、`type`、`title`、`options`（单选题、多选题等选项）以及 `correct_answer`（正确答案）等信息。如果是图片题，还会包含 `image_url` 字段，用于存储题目相关的图片链接。
2. **表字段的设计**：该设计确保试卷生成表能存储关于试卷的基本信息以及题目列表，并且通过 `JSON` 格式存储题目数据，使得该字段具有较高的灵活性和扩展性。
3. **数据表用途**：这个表主要用于存储生成的试卷信息，每个试卷的题目通过 `questions` 字段以 `JSON` 格式保存，支持不同题型和分数设置。

支持题目数据的灵活管理。

 CREATE TABLE `ym_exam_papers` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID，自增',         -- 主键ID，自增
  `title` VARCHAR(255) NOT NULL COMMENT '试卷标题',                   -- 试卷标题
  `description` TEXT DEFAULT NULL COMMENT '试卷描述',                 -- 试卷描述
  `cover_image` VARCHAR(255) DEFAULT NULL COMMENT '封面图片URL',       -- 封面图片URL
  `total_score` INT(11) NOT NULL DEFAULT 0 COMMENT '试卷总分',        -- 试卷总分
  `questions` JSON NOT NULL COMMENT '题目列表（JSON数组）',           -- 题目列表（JSON数组）
  `view_count` INT(11) DEFAULT 0 COMMENT '浏览量',                   -- 浏览量
  `category_id` BIGINT(20) DEFAULT NULL COMMENT '分类ID',             -- 分类ID
  `publish_time` INT(10) DEFAULT 0 COMMENT '最近一次发布时间',        -- 最近一次发布时间
  `status` TINYINT(1) DEFAULT 1 COMMENT '状态：1=正常，0=关闭',       -- 状态：1=正常，0=关闭
  `creator` VARCHAR(150) DEFAULT NULL COMMENT '创建人用户名',         -- 创建人用户名
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间', -- 创建时间
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间', -- 更新时间
  `deleted_at` TIMESTAMP DEFAULT NULL COMMENT '删除时间（NULL=未删除）', -- 删除时间（NULL=未删除）
  PRIMARY KEY (`id`) COMMENT '主键：试卷ID'                           -- 设置主键
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='试卷表：存储试卷的基本信息与题目内容';


 