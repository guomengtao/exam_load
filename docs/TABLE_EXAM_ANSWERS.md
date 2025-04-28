# 学生答卷表（exam_answers）

| 字段名         | 类型        | 说明                                  | 默认值         |
| --------------- | ----------- | ------------------------------------- | -------------- |
| submission_id   | VARCHAR(64) | 提交ID，唯一标识一次提交，使用UUID  | NULL           |
| exam_id         | VARCHAR(64) | 试卷ID                               | 无默认值      |
| student_id      | VARCHAR(64) | 学生ID                               | NULL           |
| student_name    | VARCHAR(255) | 学生姓名                              | NULL           |
| answers         | JSON        | 学生提交的所有题目答案               | NULL           |
| score_details   | JSON        | 每题得分情况（题号->得分）            | NULL           |
| total_score     | INT         | 总分                                  | 0              |
| created_at      | TIMESTAMP   | 创建时间（提交时间）                 | 当前时间（`CURRENT_TIMESTAMP`） |
| updated_at      | TIMESTAMP   | 更新时间                             | 当前时间（`CURRENT_TIMESTAMP`） |

**备注：**
- `submission_id` 字段使用 UUID 来唯一标识每一次提交。
- `exam_id` 和 `student_id` 需要保证唯一性组合，可以保证每个学生在同一试卷上只能提交一次。
- `student_name` 字段用于存储学生的姓名，避免出现填写错误的情况。
- 默认 `total_score` 为 0，如果没有得分数据。
- `created_at` 和 `updated_at` 都使用当前时间戳。

# 备注
- `answers` 字段示例：`{"1":"A", "2":"B,C", "3":"D"}`
- `score_details` 示例：`{"1":5, "2":3, "3":0}`
- `total_score`：计算得出的总分，例如 8分
- 保证 exam_id + student_id 组合唯一，避免重复提交。
- 如果高并发，可以考虑加上 Redis 缓存、批量入库。

# 核心设计 - 学生答卷存储方案

## 选择理由
- 支持高并发写入。
- 支持不同题型、不同题数的试卷。
- 存储和读取都很高效。
- JSON 字段灵活，可扩展性强。
- 可以直接缓存至 Redis，进一步加速查询。
- 便于后期统计总分、单题得分、答题分析。

## 表设计
见 exam_answers 表定义。

## 特别说明
- 答案、得分都存 JSON，方便解析。
- 需要后端统一封装答案格式校验，避免前端传错。
- 后期可根据数据量分表、归档历史数据。


CREATE TABLE exam_answers (
    submission_id VARCHAR(64) PRIMARY KEY,  -- 提交ID，唯一标识一次提交，使用UUID
    exam_id VARCHAR(64) NOT NULL,           -- 试卷ID
    student_id VARCHAR(64) NOT NULL,        -- 学生ID
    student_name VARCHAR(255) DEFAULT NULL, -- 学生姓名，默认为NULL
    answers JSON DEFAULT NULL,              -- 学生提交的所有题目答案，默认为NULL
    score_details JSON DEFAULT NULL,        -- 每题得分情况（题号->得分），默认为NULL
    total_score INT DEFAULT 0,              -- 总分，默认为0
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- 创建时间（提交时间），默认当前时间戳
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  -- 更新时间，默认当前时间戳，更新时自动修改
);

-- 注意：可以根据业务需求在适当的字段上添加索引，例如：exam_id, student_id 和 submission_id 的组合索引。
-- 示例：
CREATE INDEX idx_exam_student_submission ON exam_answers(exam_id, student_id, submission_id);