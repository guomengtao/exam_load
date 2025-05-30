.
├── Ai-todo.md
├── app
│   ├── biz
│   │   ├── role_biz.go
│   │   ├── role_biz_test.go
│   │   ├── teacher_biz.go
│   │   ├── user_biz.go
│   │   └── user_biz.go.back
│   ├── controllers
│   │   ├── admin_controller.go
│   │   ├── admin_controller_test.go
│   │   ├── boy_controller.go
│   │   ├── csv_import_controller.go
│   │   ├── csv_import_controller_test.go
│   │   ├── export_controller.go
│   │   ├── export_controller_test.go
│   │   ├── helloa_controller.go
│   │   ├── hello_controller.go
│   │   ├── hello.go
│   │   ├── king_controller.go
│   │   ├── main_test.go
│   │   ├── role_controller.go
│   │   ├── role_controller_test.go
│   │   ├── status_controller.go
│   │   ├── task_controller.go
│   │   ├── teacher_controller.go
│   │   └── user_controller.go
│   ├── models
│   │   ├── admin.go
│   │   ├── exam_answer.go
│   │   ├── member.go
│   │   ├── role.go
│   │   ├── teacher.go
│   │   └── user.go
│   ├── services
│   │   ├── admin_service.go
│   │   ├── admin_service_test.go
│   │   ├── csv_importer.go
│   │   ├── csv_importer_test.go
│   │   ├── export_answer_service.go
│   │   ├── export_answer_service_test.go
│   │   ├── main_test.go
│   │   ├── member_service.go
│   │   ├── member_service_test.go
│   │   ├── mock_user_pool.go
│   │   ├── mock_user_pool_test.go
│   │   ├── redis_cleaner.go
│   │   ├── redis_cleaner_test.go
│   │   ├── redis_importer.go
│   │   ├── redis_importer_test.go
│   │   ├── redis_writer.go
│   │   ├── redis_writer_test.go
│   │   ├── role_service.go
│   │   ├── role_service_test.go
│   │   ├── task_runner.go
│   │   ├── teacher_service.go
│   │   ├── user_service.go
│   │   ├── user_service.go.back
│   │   └── user_service_test.go
│   └── validators
│       └── role_validator.go
├── ARCHITECTURE.md
├── auth
│   ├── auth.go
│   ├── middleware.go
│   └── permission.go
├── config
│   └── config.go
├── dev_notes.md
├── docs
│   ├── data
│   │   ├── full_schema_and_data.sql
│   │   ├── task.data.sql
│   │   ├── task_log.data.sql
│   │   ├── task_log.mk
│   │   ├── task.mk
│   │   ├── tm_admin.data.sql
│   │   ├── tm_admin.mk
│   │   ├── tm_exam_answers.data.sql
│   │   ├── tm_exam_answers.mk
│   │   ├── tm_exam_papers.data.sql
│   │   ├── tm_exam_papers.mk
│   │   ├── tm_exam_template.data.sql
│   │   ├── tm_exam_template.mk
│   │   ├── tm_permission.data.sql
│   │   ├── tm_permission.mk
│   │   ├── tm_role.data.sql
│   │   ├── tm_role.mk
│   │   ├── tm_role_permission.data.sql
│   │   ├── tm_role_permission.mk
│   │   ├── tm_user.data.sql
│   │   ├── tm_user.mk
│   │   ├── ym_admin.data.sql
│   │   ├── ym_admin.mk
│   │   ├── ym_answer.data.sql
│   │   ├── ym_answer.mk
│   │   ├── ym_answer_old.data.sql
│   │   ├── ym_answer_old.mk
│   │   ├── ym_article_contens.data.sql
│   │   ├── ym_article_contens.mk
│   │   ├── ym_article.data.sql
│   │   ├── ym_article_limit.data.sql
│   │   ├── ym_article_limit.mk
│   │   ├── ym_article.mk
│   │   ├── ym_cate_contens.data.sql
│   │   ├── ym_cate_contens.mk
│   │   ├── ym_cate.data.sql
│   │   ├── ym_cate.mk
│   │   ├── ym_guest.data.sql
│   │   ├── ym_guest.mk
│   │   ├── ym_link.data.sql
│   │   ├── ym_link.mk
│   │   ├── ym_member.data.sql
│   │   ├── ym_member.mk
│   │   ├── ym_score_stat.data.sql
│   │   └── ym_score_stat.mk
│   ├── data_import_export.md
│   ├── docs.go
│   ├── GEN_CURD.md
│   ├── goapp_structure.md
│   ├── swagger.json
│   └── swagger.yaml
├── gin.log
├── go.mod
├── go.sum
├── handlers
│   ├── dbinfo.go
│   ├── exam_answer.go
│   ├── exam.go
│   ├── exam_paper.go
│   ├── exam_paper_redis.go
│   ├── exam_template.go
│   ├── hello.go
│   ├── source_check.go
│   ├── status.go
│   ├── upload.go
│   └── version.go
├── hello
├── hello.log
├── LICENSE.md
├── main.go
├── output.log
├── package.json
├── package-lock.json
├── README.md
├── routes
│   ├── gen_routes.go
│   └── routes.go
├── run.sh
├── sql.sh
├── TREE.md
├── tree.sh
└── utils
    ├── db.go
    ├── db_gorm.go
    ├── db_sqlx.go
    ├── gen
    │   ├── gen_back.go.back
    │   ├── gen.go
    │   ├── Gen_logic.md
    │   ├── meta
    │   │   └── field.go
    │   └── templates
    │       ├── biz
    │       │   ├── biz_skeleton.tpl
    │       │   └── biz.tpl
    │       ├── biz_skeleton.tpl
    │       ├── biz.tpl
    │       ├── controller
    │       │   └── controller.tpl
    │       ├── controller_skeleton.tpl
    │       ├── controller_test.tpl
    │       ├── controller.tpl
    │       ├── model
    │       │   └── model.tpl
    │       ├── model.tpl
    │       ├── service
    │       │   ├── service.tpl
    │       │   └── skeleton.tpl
    │       ├── service_skeleton.tpl
    │       ├── service.tpl
    │       └── validator
    │           └── validator.tpl
    ├── generated
    │   ├── biz
    │   │   ├── role_biz_skeleton.go
    │   │   ├── teacher_biz_skeleton.go
    │   │   └── user_biz_skeleton.go
    │   ├── controller
    │   │   ├── boy_skeleton.go
    │   │   ├── king_skeleton.go
    │   │   ├── role_skeleton.go
    │   │   ├── teacher_skeleton.go
    │   │   └── user_skeleton.go
    │   ├── models
    │   │   └── member.go
    │   ├── service
    │   │   ├── role_service_skeleton.go
    │   │   ├── teacher_service_skeleton.go
    │   │   └── user_service_skeleton.go
    │   └── services
    │       ├── member_biz.go
    │       ├── member_service.go
    │       └── member_service_test.go
    ├── genlib
    │   ├── gen_biz.go
    │   ├── gen_controller.go
    │   ├── gen_model.go
    │   ├── gen_router.go
    │   ├── gen_service.go
    │   ├── gen_skeleton.go
    │   └── gen_utills.go
    ├── jwt.go
    ├── password.go
    ├── path.go
    ├── queue.go
    ├── redis.go
    ├── status.go
    └── time.go

28 directories, 191 files
