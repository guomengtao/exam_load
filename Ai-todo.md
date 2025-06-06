 

前置条件说明（开发固定规范）

🛠️ 文件重要性
	•	本文件为开发沟通的 固定前置条件。
	•	每次开发前，必须 完整读取并遵循 本文件内容。
	•	如果未读取或未遵循本文件内容导致问题，责任在执行方。

📜 使用规则
	•	开发与沟通，默认以本前置条件为准，除非另行特别说明。
	•	如果实际需求与本文件存在冲突或不合理，必须及时反馈，不得自行更改。
	•	如发现前置条件需要优化、补充，双方协商后更新本文件。

⚙️ 开发基础设定

项目	内容
开发语言	Golang
框架	Gin (gin-gonic/gin)
数据库	MySQL
表前缀	来自 .env 文件中的 TABLE_PREFIX，如默认 ym_
接口规范	标准 RESTful 风格，返回结构固定：code、msg、data
JSON返回格式	使用 gin.H{}
分页参数	使用 page 和 limit，默认 page=1，limit=10
查询默认行为	不传 id 时，默认返回10条记录，支持分页
时间字段	统一使用 created_at、updated_at，插入数据时用 NOW()

🚫 禁止事项
	•	不使用 Echo、Fiber、原生 net/http，统一使用 Gin。
	•	不随意修改框架、数据库、返回结构。
	•	不擅自变更接口风格或字段命名。

❗ 冲突与优化处理
	•	发现前置条件内部有逻辑冲突、不合理要求时，必须即时提示。
	•	必须在沟通时指出冲突，并提出修正建议。
	•	不允许无提示地忽略或擅自修改前置条件。

⸻

⚡ 特别提醒

本文件作为【开发固定规范】，每次开发开始前，执行方必须首先确认已完整阅读，并在开发过程中严格遵守。
如遇变化，需双方沟通一致后，正式更新本文件内容。

⸻

✅ 这就是完整版本！非常正式、专业，并且符合你的要求。
你可以保存成一个 .md 文件，比如叫：

project-development-conditions.md

⸻

⚡下一步建议

要不要我顺便帮你也写一版「初版.env 文件模板」？这样前后都完整了～
要的话告诉我：“顺便给我.env模板！” 🔥

要继续？要我把这个 .md 文件直接打包发给你？📄

非常好！我明白你的需求了，下面我会根据你提供的要求整理一个简明的 命名规范文档，并且结合你提出的结构进行标准化，确保以后能方便 AR 按照这个标准开发接口和命名文件。

⸻

📝 命名规范文档

3. 接口命名规范
	•	接口的路径：接口路径应该以 API 开头，使用 RESTful API 规范，并且路径使用 表名，采用下划线命名法（snake_case）。
	•	示例：
	•	POST /api/exam_paper: 用于创建试卷的接口。
	•	GET /api/exam_paper/{id}: 用于获取单个试卷的接口。
	•	GET /api/exam_paper: 用于获取试卷列表的接口。
	•	PUT /api/exam_paper/{id}: 用于更新试卷的接口。
	•	DELETE /api/exam_paper/{id}: 用于删除试卷的接口。
	•	GET /api/exam_template: 用于获取试卷模板列表的接口。

⸻

4. 接口请求方式
	•	采用 HTTP 方法 来区分操作：
	•	GET: 获取资源（如获取单个记录、获取列表）。
	•	POST: 创建资源（如创建新记录）。
	•	PUT: 更新资源（如更新现有记录）。
	•	DELETE: 删除资源（如删除记录）。

⸻

示例：按照规范开发的接口

操作类型	文件名	函数名(重要)	接口路径
创建记录	exam_paper.go	CreateExamPaper	POST /api/exam_paper
查询单个记录	exam_paper.go	GetExamPaper	GET /api/exam_paper/{id}
查询列表	exam_paper.go	GetExamPaperList	GET /api/exam_paper
更新记录	exam_paper.go	UpdateExamPaper	PUT /api/exam_paper/{id}
删除记录	exam_paper.go	DeleteExamPaper	DELETE /api/exam_paper/{id}



⸻

5. 命名优先级和清晰度
	•	文件名优先级：文件名直接和表名挂钩，可以让我们清楚知道哪个文件处理哪个表的数据。
	•	函数命名优先级：函数名应该能清楚表达操作类型和数据表，这样后续开发和维护时能一眼看出功能。
	•	接口路径优先级：接口路径应该清晰明了，符合 RESTful 风格，便于开发和团队协作。



## 统一前缀
所有接口统一前缀 `/api/`

## 资源层级
路径格式为： `/api/{模块名}/{动作名}`  
模块名即表名或业务模块名；动作名表示具体操作或来源类型（如 redis, mysql）


 