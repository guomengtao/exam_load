.
├── Ai-todo.md
├── app
│   ├── biz
│   │   ├── admin_biz.go
│   │   ├── badminton_game_biz.go
│   │   ├── badminton_game_biz_test.go
│   │   ├── file_info_biz.go
│   │   ├── role_biz.go
│   │   ├── role_biz_test.go
│   │   ├── teacher_biz.go
│   │   └── user_biz.go
│   ├── controllers
│   │   ├── admin_controller_test.go
│   │   ├── badminton_game_controller.go
│   │   ├── badminton_game_controller_mysql_test.go
│   │   ├── badminton_game_controller_test.go
│   │   ├── csv_import_controller.go
│   │   ├── csv_import_controller_test.go
│   │   ├── export_controller.go
│   │   ├── export_controller_test.go
│   │   ├── file_info_controller.go
│   │   ├── main_test.go
│   │   ├── role_controller.go
│   │   ├── role_controller_test.go
│   │   ├── status_controller.go
│   │   ├── task_controller.go
│   │   ├── teacher_controller.go
│   │   └── user_controller.go
│   ├── models
│   │   ├── admin.go
│   │   ├── badminton_game.go
│   │   ├── exam_answer.go
│   │   ├── file_info.go
│   │   ├── member.go
│   │   ├── role.go
│   │   ├── teacher.go
│   │   └── user.go
│   ├── services
│   │   ├── admin_service.go
│   │   ├── admin_service_test.go
│   │   ├── badminton_game_service.go
│   │   ├── badminton_game_service_test.go
│   │   ├── csv_importer.go
│   │   ├── csv_importer_test.go
│   │   ├── export_answer_service.go
│   │   ├── file_info_service.go
│   │   ├── file_info_service_test.go
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
│   │   ├── teacher_service_test.go
│   │   ├── user_service.go
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
│   │   ├── badminton_games.data.sql
│   │   ├── full_schema_and_data.sql
│   │   ├── task.data.sql
│   │   ├── task_log.data.sql
│   │   ├── task_log.mk
│   │   ├── task.mk
│   │   ├── tm_admin.data.sql
│   │   ├── tm_admin.mk
│   │   ├── tm_badminton_game.data.sql
│   │   ├── tm_file_info.data.sql
│   │   ├── tm_permission.data.sql
│   │   ├── tm_permission.mk
│   │   ├── tm_role.data.sql
│   │   ├── tm_role.mk
│   │   ├── tm_role_permission.data.sql
│   │   ├── tm_role_permission.mk
│   │   ├── tm_teacher.data.sql
│   │   ├── tm_user.data.sql
│   │   └── tm_user.mk
│   ├── data_import_export.md
│   ├── DEV_GUIDE.md
│   ├── docs.go
│   ├── FIELD_UPDATE_POLICY.md
│   ├── GEN_CURD.md
│   ├── GEN_UPGRADE.md
│   ├── goapp_structure.md
│   ├── swagger.json
│   └── swagger.yaml
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
├── LICENSE.md
├── main.go
├── mysql.sh
├── package.json
├── package-lock.json
├── README.md
├── routes
│   ├── gen_routes.go
│   └── routes.go
├── run.sh
├── sql.sh
├── test_gen_api.sh
├── TREE.md
├── tree.sh
└── utils
    ├── db.go
    ├── db_gorm.go
    ├── db_sqlx.go
    ├── gen
    │   ├── gen.go
    │   ├── Gen_logic.md
    │   ├── meta
    │   │   ├── db.go
    │   │   ├── field.go
    │   │   └── naming.go
    │   └── templates
    │       ├── biz_skeleton.tpl
    │       ├── biz.tpl
    │       ├── controller_skeleton.tpl
    │       ├── controller_test.tpl
    │       ├── controller.tpl
    │       ├── model.tpl
    │       ├── service_skeleton.tpl
    │       ├── service.tpl
    │       └── validator.tpl
    ├── generated
    │   ├── biz
    │   │   ├── badminton_game_biz_skeleton.go
    │   │   ├── file_info_biz_skeleton.go
    │   │   ├── role_biz_skeleton.go
    │   │   ├── teacher_biz_skeleton.go
    │   │   └── user_biz_skeleton.go
    │   ├── controller
    │   │   ├── badminton_game_skeleton.go
    │   │   ├── file_info_skeleton.go
    │   │   ├── role_skeleton.go
    │   │   ├── teacher_skeleton.go
    │   │   └── user_skeleton.go
    │   ├── models
    │   │   └── member.go
    │   ├── service
    │   │   ├── badminton_game_service_skeleton.go
    │   │   ├── common.go
    │   │   ├── file_info_service_skeleton.go
    │   │   ├── role_service_skeleton.go
    │   │   ├── teacher_service_skeleton.go
    │   │   └── user_service_skeleton.go
    │   └── services
    ├── genlib
    │   ├── gen_biz.go
    │   ├── gen_controller.go
    │   ├── gen_model.go
    │   ├── gen_router.go
    │   ├── gen_service.go
    │   └── gen_utills.go
    ├── jwt.go
    ├── password.go
    ├── path.go
    ├── queue.go
    ├── redis.go
    ├── response.go
    ├── response_test.go
    ├── status.go
    └── time.go

23 directories, 170 files
